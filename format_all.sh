#!/bin/bash

set -e

find . -type f -name '*.go' | while read -r file; do
    echo "Formatting: $file"
    gofmt -w "$file" 2>&1 | head
done
