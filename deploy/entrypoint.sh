#!/bin/sh
# Runs DB migrations, then execs the application binary.
#
# DATABASE_URL для migrate задаётся так:
#   - POSTGRES_DSN — если задан (например в K8s), используется как есть;
#   - иначе собирается из DB_* / DB_USE_TLS как в internal/config.Load().
#
# Пропуск миграций (инциденты / отладка): SKIP_DB_MIGRATE=1
set -eu

_resolve_database_url() {
	if [ -n "${POSTGRES_DSN:-}" ]; then
		printf '%s' "$POSTGRES_DSN"
		return
	fi
	_sslmode="disable"
	case "${DB_USE_TLS:-false}" in
	true | True | TRUE | 1) _sslmode="require" ;;
	esac
	_host="${DB_HOST:-localhost}"
	_port="${DB_PORT:-5432}"
	_name="${DB_NAME:-selection_service}"
	_user="${DB_USER:-selection}"
	_pass="${DB_PASSWORD:-selection}"
	printf 'postgres://%s:%s@%s:%s/%s?sslmode=%s' \
		"$_user" "$_pass" "$_host" "$_port" "$_name" "$_sslmode"
}

APP_BIN="${APP_BIN:-/selection-service}"

if [ "${SKIP_DB_MIGRATE:-0}" != "1" ]; then
	_db_url="$(_resolve_database_url)"
	# golang-migrate для PostgreSQL сериализует параллельные up через advisory lock — безопасно при нескольких репликах.
	migrate \
		-path=/migrations \
		-database="$_db_url" \
		up
fi

exec "$APP_BIN"
