#!/bin/bash
set -e

echo "=== Starting keycloak... ==="
exec /opt/keycloak/bin/kc.sh $@
