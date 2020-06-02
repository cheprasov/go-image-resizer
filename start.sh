#!/bin/bash

./start \
  --source-path="$1" \
  --output-path="${1}_resized" \
  --skip-small=true \
  --width=620 \
  --quality=75 \

