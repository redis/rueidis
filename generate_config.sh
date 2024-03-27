#!/usr/bin/env bash

# Find all directories containing go.mod files
modules=$(find . -maxdepth 2 -type f -name "go.mod" | xargs -n 1 dirname)

# Start the generated config with the version
cat << EOF
version: 2.1

orbs:
  go: circleci/go@1.7.3

jobs:
  build:
    machine:
      image: ubuntu-2204:current
    resource_class: large
    parallelism: 3
    steps:
EOF

# Loop through each module and generate job configurations
for module in $modules; do
  module_name=$(basename "$module")
  cat << EOF
      - run: # test ./$module_name/go.mod
          name: Test $module_name
          command: |
            cd "\$CIRCLE_WORKING_DIRECTORY/$module_name"
            list=\$(go list ./... | circleci tests split --split-by=timings)
            echo "Test Packages: \$list"
            for n in {1..5}; do
              ../dockertest.sh \$list && break
            done
          no_output_timeout: 15m
EOF
done
