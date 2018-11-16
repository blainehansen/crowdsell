alter default privileges revoke execute on functions from public;

-- TODO not secure
create role postgraphile_server_user login password 'postgraphile' noinherit;

create role postgraphile_known_user;
grant postgraphile_known_user to postgraphile_server_user;

create role postgraphile_inspect_user login password 'postgraphile';
grant postgraphile_server_user to postgraphile_inspect_user;
grant postgraphile_known_user to postgraphile_inspect_user;


-- TODO not secure
create role golang_server_user login password 'golang' noinherit;

create role golang_known_user;
grant golang_known_user to golang_server_user;

grant usage on schema public to postgraphile_server_user, postgraphile_known_user, golang_server_user, golang_known_user;


create or replace function current_person_id() returns uuid
as $$
begin
	return current_setting('jwt.claims.id')::uuid
end;
$$ language plpgsql;


create or replace function trigger_set_created() returns trigger
as $$
begin
	new.date_created = now();
	new.date_updated = now();
	return new;
end;
$$ language plpgsql;

create or replace function trigger_set_updated() returns trigger
as $$
begin
	new.date_updated = now();
	return new;
end;
$$ language plpgsql;

create or replace function default_slug() returns trigger
as $$
begin
	new.slug := new.id::text;
	return new;
end;
$$ language plpgsql;
