#!/usr/bin/env bash
#
# Shared benchmark tooling for the two benchmark workflows:
#   - benchmark-main.yml: records the dashboard baseline on gh-pages after a merge to main.
#   - benchmark-pr.yml:   compares a pull request against its base and writes the result to the job summary.
#
# Both call this script so that they cannot drift apart: the benchmark list, the Redis startup and the
# report formatting live only here.
#
# Usage:
#   bench.sh start-redis [compose-dir]
#   bench.sh run main <out.txt>
#   bench.sh run pr <base-dir> <head-dir> <outdir>
#   bench.sh report main <in.txt> <averaged.txt>
#   bench.sh report pr <outdir> <base-sha> <head-sha>
#
# pipefail matters here: the benchmark pipelines end in `tee`, so without it a panic, a timeout or a
# b.Errorf in the test binary would be hidden behind tee's exit status.

set -euo pipefail

# The single list of benchmarks tracked by both the dashboard and the PR comparison.
BENCH='^(BenchmarkClient_Do|BenchmarkClient_Do_Parallelism|BenchmarkClient_DoMulti)$'
# The file that defines them. `run pr` builds the base with the head's copy, so it must be self-contained.
BENCH_FILE='bench_e2e_test.go'

# Samples per side. `run main` uses -count, `run pr` alternates base and head for this many rounds.
SAMPLES="${BENCH_SAMPLES:-10}"
# Time limit for one benchmark run. Use a single unit (for example 30m) so that both go test and timeout(1) accept it.
TIMEOUT="${BENCH_TIMEOUT:-30m}"

# Significant slowdowns above this percentage are flagged as severe in the PR report.
REGRESSION_THRESHOLD="${REGRESSION_THRESHOLD:-10}"

usage() {
  cat >&2 <<'EOF'
Usage:
  bench.sh start-redis [compose-dir]
  bench.sh run main <out.txt>
  bench.sh run pr <base-dir> <head-dir> <outdir>
  bench.sh report main <in.txt> <averaged.txt>
  bench.sh report pr <outdir> <base-sha> <head-sha>
EOF
  exit 2
}

start_redis() {
  local dir="${1:-.}" count=0
  cd "$dir"
  docker compose up -d redis
  until docker compose exec -T redis redis-cli ping | grep -q PONG; do
    if [ "$count" -gt 30 ]; then
      echo "Timeout waiting for Redis" >&2
      exit 1
    fi
    echo "Waiting for Redis..."
    sleep 1
    count=$((count + 1))
  done
}

# Print, once each, the benchmarks in file $2 that printed a "--- $1:" line, where $1 is FAIL or SKIP. A failing
# sub-benchmark also fails its parent, so only the innermost names are kept.
innermost() {
  { grep -E "^[[:space:]]*--- $1: " "$2" || true; } | awk '
    !seen[$3]++ { name[++n] = $3 }
    END {
      for (i = 1; i <= n; i++) {
        leaf = 1
        for (j = 1; j <= n; j++) if (index(name[j], name[i] "/") == 1) leaf = 0
        if (leaf) print name[i]
      }
    }'
}

# Dashboard run: any failure fails the step, so partial results are never published. A skipped benchmark counts as a
# failure too, because in CI it means the environment broke, for example Redis was unreachable. -v is what makes
# the testing package print the --- SKIP line of a skipped benchmark.
run_main() {
  local out="${1:?usage: bench.sh run main <out.txt>}" skipped
  go test -v -bench="$BENCH" -run='^$' -benchmem -count="$SAMPLES" -timeout "$TIMEOUT" . | tee "$out"
  skipped=$(innermost SKIP "$out")
  if [ -n "$skipped" ]; then
    echo "::error::Some benchmarks were skipped, so nothing is published: ${skipped//$'\n'/, }"
    exit 1
  fi
}

# One round for one side. A failed round is recorded in failures.txt instead of aborting the run, so the
# remaining rounds still produce data and the report can say the results are incomplete.
#
# The test binary's own -test.timeout does not cover benchmarks (the testing package stops its alarm before running
# them), so timeout(1) bounds the round instead. SIGQUIT makes the binary print every goroutine's stack before it
# exits, which shows where it hung. -test.v is what makes the testing package print the --- SKIP line of a skipped
# benchmark.
run_round() {
  local round="$1" side="$2" dir="$3" outdir="$4" rc=0 failed skipped name
  local raw="$outdir/$side.round"

  (cd "$dir" && timeout --signal=QUIT --kill-after=1m "$TIMEOUT" "$outdir/$side.test" -test.run='^$' -test.bench="$BENCH" -test.benchmem -test.count=1 -test.v) | tee "$raw" || rc=$?
  cat "$raw" >> "$outdir/$side.txt"

  # Name each benchmark that failed or was skipped. A skip counts as a failure: in CI it means the environment broke,
  # for example Redis crashed partway through.
  failed=$(innermost FAIL "$raw")
  skipped=$(innermost SKIP "$raw")
  while IFS= read -r name; do
    [ -z "$name" ] || echo "round $round ($side): \`$name\` failed" >> "$outdir/failures.txt"
  done <<< "$failed"
  while IFS= read -r name; do
    [ -z "$name" ] || echo "round $round ($side): \`$name\` skipped" >> "$outdir/failures.txt"
  done <<< "$skipped"
  # A benchmark failure exits with status 1 and is named above. timeout(1) exits with 124 when the round ran out of
  # time. A panic exits with another status and prints no --- FAIL line; its output is in the step log.
  if [ "$rc" -eq 124 ]; then
    echo "round $round ($side): timed out after $TIMEOUT" >> "$outdir/failures.txt"
    TIMED_OUT=1
  elif [ "$rc" -ne 0 ] && { [ -z "$failed" ] || [ "$rc" -ne 1 ]; }; then
    echo "round $round ($side): test binary exited with status $rc" >> "$outdir/failures.txt"
  fi
}

# Build <side>.test from <dir> with <src> as its $BENCH_FILE, so that both sides run the same benchmark code and only
# the library differs. If it does not build that way, typically because a benchmark uses an API the PR adds, build
# <dir> as it is instead: with its own copy, or without one if it has none. USED records which copy was built (other,
# own or none), and <dir> is left as it was.
build_with() {
  local side="$1" dir="$2" src="$3" outdir="$4" rc=0
  local file="$dir/$BENCH_FILE" own="$outdir/$BENCH_FILE.$side"

  USED=other
  if ! cmp -s "$src" "$file"; then
    rm -f "$own"
    if [ -f "$file" ]; then cp "$file" "$own"; fi
    cp "$src" "$file"
    (cd "$dir" && go test -c -o "$outdir/$side.test" .) > "$outdir/$side-build.log" 2>&1 || rc=$?
    if [ -f "$own" ]; then cp "$own" "$file"; else rm -f "$file"; fi
    if [ "$rc" -eq 0 ]; then
      return 0
    fi
    # These compile errors are expected, so keep setup-go's problem matcher from turning them into error annotations.
    echo "::remove-matcher owner=go::"
    echo "The $side does not build with the other side's $BENCH_FILE:"
    cat "$outdir/$side-build.log"
    if [ -f "$own" ]; then
      USED=own
      echo "Building the $side with its own copy instead."
    else
      USED=none
      echo "Building the $side without it instead."
    fi
  fi
  (cd "$dir" && go test -c -o "$outdir/$side.test" .)
}

# Build both test binaries with the same copy of $BENCH_FILE: the head's, or the base's if the head has none, as on a
# branch that predates it. Any fallback is noted in the summary.
build_pr() {
  local base_dir="$1" head_dir="$2" outdir="$3" notes="$3/notes.txt"

  if [ -f "$head_dir/$BENCH_FILE" ]; then
    # Build the head first, so that a head that does not compile fails here instead of sending the base build into
    # its fallback.
    (cd "$head_dir" && go test -c -o "$outdir/head.test" .)
    build_with base "$base_dir" "$head_dir/$BENCH_FILE" "$outdir"
    case "$USED" in
      own) echo "The base doesn't build with the head's \`$BENCH_FILE\`, so it used its own copy. Benchmarks that differ between the two copies are not compared like for like." >> "$notes" ;;
      none) echo "The base doesn't build with the head's \`$BENCH_FILE\` and has no copy of its own, so it ran without it." >> "$notes" ;;
    esac
  elif [ -f "$base_dir/$BENCH_FILE" ]; then
    (cd "$base_dir" && go test -c -o "$outdir/base.test" .)
    build_with head "$head_dir" "$base_dir/$BENCH_FILE" "$outdir"
    case "$USED" in
      other) echo "The head has no \`$BENCH_FILE\`, so both sides ran the base's copy." >> "$notes" ;;
      none) echo "The head has no \`$BENCH_FILE\` and doesn't build with the base's copy, so it ran without it." >> "$notes" ;;
    esac
  else
    (cd "$head_dir" && go test -c -o "$outdir/head.test" .)
    (cd "$base_dir" && go test -c -o "$outdir/base.test" .)
  fi
}

# PR run: build a test binary per side, both with the same benchmark code (see build_pr), then alternate base and head
# so runner noise and drift affect both sides equally. A failed round does not stop the remaining rounds, but the step
# fails at the end so the check shows the failure. The report step still runs and explains it in the job summary.
run_pr() {
  local base_dir="${1:?usage: bench.sh run pr <base-dir> <head-dir> <outdir>}"
  local head_dir="${2:?usage: bench.sh run pr <base-dir> <head-dir> <outdir>}"
  local outdir="${3:?usage: bench.sh run pr <base-dir> <head-dir> <outdir>}"

  mkdir -p "$outdir"
  base_dir=$(cd "$base_dir" && pwd)
  head_dir=$(cd "$head_dir" && pwd)
  outdir=$(cd "$outdir" && pwd)
  rm -f "$outdir/base.txt" "$outdir/head.txt"
  : > "$outdir/failures.txt"
  : > "$outdir/notes.txt"

  build_pr "$base_dir" "$head_dir" "$outdir"

  # Create the result files only once both binaries are built, so that a build failure is reported as a failed
  # run rather than as a run without results.
  : > "$outdir/base.txt"
  : > "$outdir/head.txt"

  local i runs=0
  TIMED_OUT=0
  for i in $(seq "$SAMPLES"); do
    echo "::group::Round $i/$SAMPLES"
    run_round "$i" base "$base_dir" "$outdir"
    runs=$((runs + 1))
    # Check after each side: if the base hangs, the head would most likely hang too.
    if [ "$TIMED_OUT" -eq 0 ]; then
      run_round "$i" head "$head_dir" "$outdir"
      runs=$((runs + 1))
    fi
    echo "::endgroup::"
    if [ "$TIMED_OUT" -eq 1 ]; then
      # The remaining runs would most likely hang the same way, so stop and report now rather than hours later.
      if [ "$runs" -lt $((2 * SAMPLES)) ]; then
        echo "round $i: stopped after the timeout, the remaining runs were skipped" >> "$outdir/failures.txt"
      fi
      break
    fi
  done

  if [ -s "$outdir/failures.txt" ]; then
    echo "::error::Some benchmarks failed, results may be incomplete"
    exit 1
  fi
}

# Dashboard report: average the samples per benchmark into the format github-action-benchmark reads.
#
# Only result lines count: a benchmark name followed by an iteration count and ns/op. go test always prints PASS and
# ok, and -v also prints each benchmark's name on its own line before running it, so neither a non-empty file nor a
# line that merely starts with Benchmark means there are results.
report_main() {
  local in="${1:?usage: bench.sh report main <in.txt> <averaged.txt>}"
  local out="${2:?usage: bench.sh report main <in.txt> <averaged.txt>}"
  if ! grep -Eq '^Benchmark[^[:space:]]*[[:space:]]+[0-9]+[[:space:]]+[0-9.]+ ns/op' "$in"; then
    echo "No benchmark results in $in" >&2
    exit 1
  fi
  awk '/^Benchmark/ && $2 ~ /^[0-9]+$/ {
    iters[$1] = $2
    count[$1]++
    for (i=4; i<=NF; i+=2) {
      if ($i == "ns/op") time[$1] += $(i-1)
      if ($i == "B/op") bytes[$1] += $(i-1)
      if ($i == "allocs/op") allocs[$1] += $(i-1)
    }
  } END {
    for (b in count) printf "%s %d %.2f ns/op %.2f B/op %.2f allocs/op\n", b, iters[b], time[b]/count[b], bytes[b]/count[b], allocs[b]/count[b]
  }' "$in" > "$out"
}

# PR report: benchstat table written to the job summary, which needs no token and so works for fork PRs.
report_pr() {
  local outdir="${1:?usage: bench.sh report pr <outdir> <base-sha> <head-sha>}"
  local base_sha="${2:-}" head_sha="${3:-}"
  local summary="${GITHUB_STEP_SUMMARY:-/dev/stdout}"
  local server="${GITHUB_SERVER_URL:-https://github.com}" repo="${GITHUB_REPOSITORY:-}"
  local run_url="$server/$repo/actions/runs/${GITHUB_RUN_ID:-}"
  local csv="$outdir/out.csv" md="$outdir/md_output.txt" regressions="$outdir/regressions.txt"
  local failures="$outdir/failures.txt" notes="$outdir/notes.txt" new="$outdir/new.txt"
  local benchstat line
  benchstat=$(command -v benchstat || echo "$(go env GOPATH)/bin/benchstat")

  # Label the results with the PR's own commits. GITHUB_SHA is the test merge commit on pull_request events.
  local label=""
  if [ -n "$head_sha" ] && [ -n "$base_sha" ]; then
    label="Head [\`${head_sha:0:7}\`]($server/$repo/commit/$head_sha) compared with base [\`${base_sha:0:7}\`]($server/$repo/commit/$base_sha)."
  fi

  if [ ! -f "$outdir/base.txt" ] || [ ! -f "$outdir/head.txt" ]; then
    {
      echo "## 🚀 Rueidis Benchmark Results"
      echo
      [ -z "$label" ] || { echo "$label"; echo; }
      echo "❌ **The benchmark run failed before producing results.** See the [workflow run]($run_url) for details. Push a new commit or re-run the workflow to try again."
    } >> "$summary"
    return 0
  fi

  : > "$md"
  : > "$new"
  rm -f "$regressions"
  : > "$csv"
  if ! "$benchstat" -format csv "$outdir/base.txt" "$outdir/head.txt" > "$csv" 2> "$outdir/benchstat.err"; then
    echo "::warning::benchstat failed: $(head -c 300 "$outdir/benchstat.err")"
    : > "$csv"
  fi

  awk -v threshold="$REGRESSION_THRESHOLD" -v md="$md" -v reg="$regressions" -v newf="$new" -v headfile="$outdir/head.txt" '
  function format_time(v) {
      ns = v * 1e9;
      if (ns >= 1e6) return sprintf("%.2f ms/op", ns / 1e6);
      if (ns >= 1e3) return sprintf("%.2f µs/op", ns / 1e3);
      return sprintf("%.2f ns/op", ns);
  }
  BEGIN { FS="," }
  # Column header row: a blank first field, then the input files. A single file column means only one side has
  # results. If that side is the head, all its rows are new benchmarks.
  $1 == "" && $2 != "" && $2 !~ "/op$" { onecol = (NF == 3); headonly = (onecol && $2 == headfile); next }
  $2 ~ "/op$" { current_unit = $2; next }
  current_unit !~ "sec/op|time/op" { next }
  # Head results without base results: listed on their own, with only the head numbers.
  $1 != "" && $1 != "geomean" && ((headonly && $2 ~ /^[0-9]/) || (!onecol && $2 == "" && $4 ~ /^[0-9]/)) {
      if (headonly) { value = $2; ci = $3 } else { value = $4; ci = $5 }
      printf "| `%s` | %s ± %s |\n", $1, format_time(value), ci >> newf;
      next;
  }
  NF >= 7 && $1 != "" && $1 != "geomean" && $1 !~ /^,/ {
      if ($2 !~ /^[0-9]/ || $4 !~ /^[0-9]/) next;

      fullname = $1;
      split(fullname, parts, "/");
      cat = parts[1];
      shortname = substr(fullname, length(cat) + 2);
      if (shortname == "") { cat = "Ungrouped"; shortname = fullname; }

      old = $2; new = $4; delta = $6; note = $7;

      color = delta;
      if (delta ~ /\+/) {
          val = delta; gsub(/[%+]/, "", val); val = val + 0.0;
          if (val > threshold + 0) color = "🔴 **" delta "**"; else color = "🟠 " delta;
      } else if (delta ~ /-/) color = "🟢 " delta;
      else if (delta ~ /~/) color = "🟡 ~";

      fmt_old = format_time(old);
      fmt_new = format_time(new);

      row = sprintf("| <span title=\"%s\">`%s`</span> | %s | %s | %s | _%s_ |\n", fullname, shortname, fmt_old, fmt_new, color, note);

      if (!(cat in groups)) {
          groups[cat] = "";
          order[++num_cats] = cat;
      }
      groups[cat] = groups[cat] row;

      if (color ~ /🔴/) {
          regressions = regressions sprintf("| `%s` | %s | %s | %s | _%s_ |\n", fullname, fmt_old, fmt_new, color, note);
      }
  }
  END {
      for (i = 1; i <= num_cats; i++) {
          c = order[i];
          printf "#### 📦 Category: `%s`\n", c >> md;
          printf "| Benchmark | Base | Head | Delta | Note |\n|---|---|---|---|---|\n" >> md;
          printf "%s\n", groups[c] >> md;
      }
      if (regressions != "") {
          printf "%s", regressions > reg;
      }
  }' "$csv"

  local failed=0 nfail=0
  if [ -s "$failures" ]; then
    failed=1
    nfail=$(wc -l < "$failures")
  fi

  {
    echo "## 🚀 Rueidis Benchmark Results"
    echo
    if [ -n "$label" ]; then
      echo "$label"
      echo
    fi

    # How the comparison was built, for example when the base could not use the head's benchmark code.
    if [ -s "$notes" ]; then
      while IFS= read -r line; do
        echo "ℹ️ $line"
        echo
      done < "$notes"
    fi

    if [ "$failed" -eq 1 ]; then
      echo "⚠️ **Some benchmarks failed, results may be incomplete.** See the [workflow run]($run_url) for the failing output."
      echo
      head -n 10 "$failures" | sed 's/^/- /'
      if [ "$nfail" -gt 10 ]; then
        echo "- … and $((nfail - 10)) more"
      fi
      echo
    fi

    if [ ! -s "$md" ] && [ ! -s "$new" ]; then
      if [ "$failed" -eq 1 ]; then
        echo "No comparable results were produced."
      elif [ -s "$notes" ]; then
        echo "⚠️ **No comparable results were produced.** See the note above and the [workflow run]($run_url) for details."
      else
        echo "⚠️ **No comparable results were produced.** Check that the \`BENCH\` pattern in \`.github/scripts/bench.sh\` matches the benchmarks, and see the [workflow run]($run_url) for details."
      fi
    fi

    if [ -s "$md" ]; then
      if [ -s "$regressions" ]; then
        echo "⚠️ **Severe Regressions Detected:**"
        echo
        echo "| Benchmark | Base | Head | Delta | Note |"
        echo "|---|---|---|---|---|"
        cat "$regressions"
      elif [ "$failed" -eq 0 ]; then
        echo "✅ **All benchmarks passed successfully!** No severe regressions (>${REGRESSION_THRESHOLD}%) detected."
      fi
      echo
    fi

    if [ -s "$new" ]; then
      echo "#### 🆕 New benchmarks"
      echo
      echo "These benchmarks have no base results, so only the head's numbers are shown."
      echo
      echo "| Benchmark | Head |"
      echo "|---|---|"
      cat "$new"
      echo
    fi

    if [ -s "$md" ]; then
      echo "<br>"
      echo "<details>"
      echo "<summary><b>📊 Click here to view all metrics</b></summary>"
      echo
      cat "$md"
      echo
      echo "</details>"
    fi
  } >> "$summary"
}

cmd="${1:-}"
[ $# -eq 0 ] || shift
case "$cmd" in
  start-redis) start_redis "$@" ;;
  run)
    mode="${1:-}"
    [ $# -eq 0 ] || shift
    case "$mode" in
      main) run_main "$@" ;;
      pr) run_pr "$@" ;;
      *) usage ;;
    esac
    ;;
  report)
    mode="${1:-}"
    [ $# -eq 0 ] || shift
    case "$mode" in
      main) report_main "$@" ;;
      pr) report_pr "$@" ;;
      *) usage ;;
    esac
    ;;
  *) usage ;;
esac
