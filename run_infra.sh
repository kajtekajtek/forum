#!/bin/sh

set -e

echo "=== starting Docker containers...  ==="
docker compose -f infra/docker-compose.yml \
    --env-file infra/.env \
    up -d
