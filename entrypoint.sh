#!/bin/sh
set -e

if [ "$MODE" = "migrate" ]; then
  echo "Running migration..."
  go run main.go --mode=migrate
  go run main.go --mode=seed
  ./main
else
  echo "Starting server..."
  ./main
fi
