package om

import (
	"context"
	"fmt"
	"regexp"
	"slices"
	"strconv"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/internal/cmds"
)

// createAndAliasIndex creates a new versioned index, aliases it to idx, and then drops all
// existing versioned indexes (idx_vN) that were previously associated with that alias.
func createAndAliasIndex(ctx context.Context, idx string, client rueidis.Client, createCmd func(idx string) cmds.FtCreatePrefixPrefix, cmdFn func(schema FtCreateSchema) rueidis.Completed) error {
	idxRE, err := regexp.Compile("^" + regexp.QuoteMeta(idx) + "_v(\\d+)$")
	if err != nil {
		return fmt.Errorf("failed to compile regular expression: %w", err)
	}

	listCmd := client.B().FtList().Build()
	listResp, err := client.Do(ctx, listCmd).ToArray()
	if err != nil {
		return fmt.Errorf("failed to list indexes: %w", err)
	}

	currVers := make([]int, 0)
	for _, message := range listResp {
		n, err := message.ToString()
		if err != nil {
			return fmt.Errorf("FT._LIST retured non-string response: %w", err)
		}
		match := idxRE.FindStringSubmatch(n)
		if len(match) < 2 {
			continue
		}
		ver, err := strconv.Atoi(match[1])
		if err != nil {
			return fmt.Errorf("failed converting version number for index %q: %w", n, err)
		}
		currVers = append(currVers, ver)
	}

	newIndex := idx + "_v1"
	if len(currVers) > 0 {
		newIndex = fmt.Sprintf("%s_v%d", idx, slices.Max(currVers)+1)
	}

	// Create the new index
	if err := client.Do(ctx, cmdFn(createCmd(newIndex).Schema())).Error(); err != nil {
		return err
	}

	if err := client.Do(ctx, client.B().FtAliasupdate().Alias(idx).Index(newIndex).Build()).Error(); err != nil {
		return fmt.Errorf("failed to update alias: %w", err)
	}

	for _, ver := range currVers {
		currIdx := fmt.Sprintf("%s_v%d", idx, ver)
		if err := client.Do(ctx, client.B().FtDropindex().Index(currIdx).Build()).Error(); err != nil {
			return fmt.Errorf("failed to drop old index %q: %w", currIdx, err)
		}
	}

	return nil
}
