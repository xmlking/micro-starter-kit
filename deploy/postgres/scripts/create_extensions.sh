#!/bin/bash

set -e
set -u

function create_extension() {
    local database=$1
    echo "  Creating uuid-ossp extension"
    psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
        CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
EOSQL
}

if [ -n "$POSTGRES_DB" ]; then
	create_extension $POSTGRES_DB
fi
