alter default privileges revoke execute on functions from public;

-- TODO not secure
create role postgraphile_user login password 'postgraphile-password';

-- TODO not secure
create role golang_user login password 'golang_user';

create role anonymous_user;
grant anonymous_user to postgraphile_user;

create role logged_in_user;
grant logged_in_user to postgraphile_user;



grant usage on schema public to anonymous_user, logged_in_user;


-- create extension pgcrypto;

create extension pg_hashids;

create or replace function hashid(bigint) returns text
as $$
	-- select id_encode($1, '$hashid_salt', $hashid_min_length, '$hashid_alphabet');
	select id_encode($1, '$hashid_salt', 8, 'abcdefghijklmnopqrstuvwxyz');
$$ language sql;

create or replace function unhashid(text) returns bigint
as $$
	select id_decode_once($1, '$hashid_salt', 8, 'abcdefghijklmnopqrstuvwxyz');
$$ language sql;


create extension ap_pgutils;

create or replace function hash_password(text) returns text
as $$
	-- https://www.postgresql.org/docs/current/static/pgcrypto.html#id-1.11.7.35.9
	select argon2($1, encode(gen_random_bytes(64), 'base64'));
$$ language sql;

create or replace function verify_password(text, text) returns boolean
as $$
	select argon2_verify($1, $2);
$$ language sql;



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
	new.slug := hashid(new.id);
	return new;
end;
$$ language plpgsql;

create or replace function default_url_slug() returns trigger
as $$
begin
	new.url_slug := new.slug;
	return new;
end;
$$ language plpgsql;
