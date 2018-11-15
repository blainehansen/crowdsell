source .env.dev.sh

postgraphile \
	-p 4444 \
	-c "postgres://postgraphile_inspect_user:$POSTGRAPHILE_INSPECT_DATABASE_PASSWORD@$SYSTEM_DATABASE_HOST:$DATABASE_PORT/$DATABASE_DB_NAME" \
	--no-ignore-rbac \
	--append-plugins '@graphile-contrib/pg-simplify-inflector'
