#!/bin/bash
set -e

echo "Waiting for DB to be up..."
/app/scripts/wait-for.sh -t 30 pgdb:5432
echo "DB is up, ready for migrations"

echo "Starting DB migrations"
cd /app/migrations
/app/scripts/migrate.sh
echo "Migrations completed!"

echo "Starting TermIT"
/app/termit
