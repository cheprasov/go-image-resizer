#!/bin/bash

go run ./cmd/resizer.go \
  --source-path="$1" \
  --prefix='resized_' \
  --convert-to="jpg" \
  --width=$2 \
  --quality=75 \
  --large-only \
  --info
