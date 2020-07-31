#!/bin/bash

go run ./cmd/resizer.go \
  --source-path="$1" \
  --output-path="${1}_resized" \
  --width=$2 \
  --quality=75 \
  --large-only \
  --verbose \
