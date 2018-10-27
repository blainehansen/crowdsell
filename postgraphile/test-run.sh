# postgraphile -p 5555 -c 'postgres://postgraphile_user:postgraphile-password@localhost:5432/dev_database' \
postgraphile -p 5555 -c 'postgres://user:asdf@localhost:5432/dev_database' \
	--no-ignore-rbac \
	--append-plugins '@graphile-contrib/pg-simplify-inflector'
