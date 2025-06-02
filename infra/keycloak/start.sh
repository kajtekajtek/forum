#!/bin/bash
set -e

REALM_NAME=forum
IMPORT_DIR=/opt/keycloak/data/import

# check if realm is already imported
if [ ! -f /opt/keycloak/data/h2/keycloakdb.mv.db ]; then
    echo "=== Realm $REALM_NAME does not exist. Importing from JSON... ==="
    /opt/keycloak/bin/kc.sh import --file "$IMPORT_DIR/realm-export.json" --override true
else
    echo "=== Realm $REALM_NAME already exists. ==="
fi

echo "=== Starting keycloak... ==="
exec /opt/keycloak/bin/kc.sh start-dev
