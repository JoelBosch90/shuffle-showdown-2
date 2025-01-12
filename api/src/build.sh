#!/bin/bash

# Find all handler directories
for handler_dir in $(find . -type d -name "handlers"); do
    # Enter each handler subdirectory
    for dir in $handler_dir/*/; do
        if [ -d "$dir" ]; then
            echo "Building handler in $dir"
            cd "$dir"
            GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o bootstrap
            cd - > /dev/null
        fi
    done
done