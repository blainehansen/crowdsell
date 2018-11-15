-- create type jwt_token as (
--   role text,
--   id integer,
--   exp integer
-- );

-- create function login(email text, password text) returns jwt_token
-- as $$
-- declare
--   attempting_person person;
-- begin
--   select a.* into attempting_person
--   from person as a
--   where a.email = $1;

--   if verify_password(attempting_person.password, password) then
--     return (
--     	'postgraphile_known_user',
--     	attempting_person.person_id,
--     	extract(epoch from now() + interval '2 hours')
--   	)::jwt_token;
--   else
--     return null;
--   end if;
-- end;
-- $$ language plpgsql strict security definer;

-- grant execute on function login(text, text) to postgraphile_server_user, postgraphile_known_user;


-- create extension pgcrypto;

-- create extension ap_pgutils;

-- create or replace function hash_password(text) returns text
-- as $$
-- 	-- https://www.postgresql.org/docs/current/static/pgcrypto.html#id-1.11.7.35.9
-- 	select argon2($1, encode(gen_random_bytes(64), 'base64'));
-- $$ language sql;

-- create or replace function verify_password(text, text) returns boolean
-- as $$
-- 	select argon2_verify($1, $2);
-- $$ language sql;
