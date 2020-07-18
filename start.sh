#!/bin/bash

go run ./cmd/resizer.go \
  --source-path="/Users/cheprasov/Photos/" \
  --output-path="/Users/cheprasov/Photos2/" \
  --width=620 \
  --quality=75 \
  --prefix=small_ \
  --verbose \
