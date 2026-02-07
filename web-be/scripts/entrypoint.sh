#!/bin/sh
set -e

# Run migrations if DATABASE_URL is set
if [ -n "$DATABASE_URL" ]; then
  echo "Running database migrations..."
  migrate -path /app/database/migrations -database "$DATABASE_URL" up
  echo "Migrations complete."
fi

exec "$@"
