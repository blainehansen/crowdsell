create table person (
	id uuid primary key default gen_random_uuid(),
	date_created timestamptz not null,
	date_updated timestamptz not null,

	slug text not null unique,
	email text not null unique check (email ~* '^.+@.+\..+$'),
	password text not null,

	name text,
	bio text,
	location text,
	links text,
	profile_photo_version text,

	has_payment_user boolean default false not null,
	forgot_password_token text,

	general_search_vector tsvector
);

alter table person enable row level security;

grant select (id, slug, name, bio, location, links, profile_photo_version)
	on table person to postgraphile_server_user, postgraphile_known_user;
create policy pg_select_person on person for select
	using (true);

grant update (name, bio, location, links)
	on table person to postgraphile_known_user;
create policy pg_update_person on person for update to postgraphile_known_user
	using (id = current_person_id());

-- grant delete on table person to postgraphile_known_user;
-- create policy delete_person on person for delete to postgraphile_known_user
--   using (id = current_person_id());


grant select (id, slug, email, password) on table person to golang_server_user;
create policy go_select_person on person for select using (true);

grant insert (name, email, password) on table person to golang_server_user;
create policy go_insert_person on person for insert to golang_server_user with check (true);


create trigger set_created_for_person
before insert on person
for each row
execute procedure trigger_set_created();

create trigger set_updated_for_person
before update on person
for each row
execute procedure trigger_set_updated();

create trigger default_slug_for_person
before insert on person
for each row
execute procedure default_slug();

create trigger search_update_person_general_search_vector
before insert or update of name, bio on person
for each row
execute procedure tsvector_update_trigger(general_search_vector, 'pg_catalog.english', name, bio);

create index person_general_search_vector_idx on person using gin (general_search_vector);


-- create function current_person() returns person
-- as $$
--   select *
--   from person
--   where id = current_person_id()
-- $$ language sql stable;

-- grant execute on function current_person() to postgraphile_server_user, postgraphile_known_user;

-- comment on function current_person() is E'@omit';
