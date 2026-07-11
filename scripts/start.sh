#!/bin/bash

set -e

echo "Starting auth assignment..."

echo "Generating keys..."

./scripts/generate-keys.sh


echo "Starting Docker containers..."

docker compose up --build