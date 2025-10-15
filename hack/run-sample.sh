#!/usr/bin/env bash
set -euo pipefail

OUT_FILE="sample-recommendation.yaml"

echo "Running analyzer sample and writing recommendation to ${OUT_FILE}"

go run ../cmd/controller --sample > "${OUT_FILE}"

echo "Wrote ${OUT_FILE}"
