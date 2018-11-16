create type project_category_type as enum (
	'COMPUTER_HARDWARE',
	'COMPUTER_SOFTWARE'
);

create table project (
	id uuid primary key default gen_random_uuid(),
	date_created timestamptz not null,
	date_updated timestamptz not null,

	slug text not null unique,
	person_id bigint not null references person(id),

	title text not null unique,
	description text,
	story text,
	promises text[] default '{}' not null,
	category project_category_type,
	goal bigint,
	upload_images text[] default '{}' not null,

	general_search_vector tsvector
);

alter table project enable row level security;

grant select (id, slug, title, description, story, promises, category, goal, upload_images)
	on table project to postgraphile_server_user, postgraphile_known_user;
create policy select_project on project for select
	using (true);

grant insert, update (slug, title, description, story, promises, category, goal, upload_images)
	on table project to postgraphile_known_user;
create policy insert_project on project for insert to postgraphile_known_user
	with check (person_id = current_person_id());
create policy update_project on project for update to postgraphile_known_user
	using (person_id = current_person_id());

grant delete
	on table project to postgraphile_known_user;
create policy delete_project on project for delete to postgraphile_known_user
	using (person_id = current_person_id());



create trigger set_created_for_project
before insert on project
for each row
execute procedure trigger_set_created();

create trigger set_updated_for_project
before update on project
for each row
execute procedure trigger_set_updated();

create trigger _1_default_slug_for_project
before insert on project
for each row
execute procedure default_slug();

create trigger search_update_project_general_search_vector
before insert or update of title, description on project
for each row
execute procedure tsvector_update_trigger(general_search_vector, 'pg_catalog.english', title, description);

create index project_general_search_vector_idx on project using gin (general_search_vector);



-- create type project_tag_type as enum (
-- 	'yep'
-- );

-- create table project_tag (
-- 	id serial not null primary key,
-- 	date_created timestamptz not null,
-- 	date_updated timestamptz not null,
-- 	slug text not null,
-- 	tag_type project_tag_type not null,
-- 	project_id bigint not null references project(id),

-- 	constraint project_tag_unique_project_tag unique (tag_type, project_id)
-- );

-- create trigger set_created_for_project_tag
-- before insert on project_tag
-- for each row
-- execute procedure trigger_set_created();

-- create trigger set_updated_for_project_tag
-- before update on project_tag
-- for each row
-- execute procedure trigger_set_updated();

-- create trigger _1_default_slug_for_project_tag
-- before insert on project_tag
-- for each row
-- execute procedure default_slug();



-- create type project_pledge_state_type as enum (
-- 	'UNPAID',
-- 	'PAID',
-- 	'RELEASED'
-- );

-- create table project_pledge (
-- 	id serial not null primary key,
-- 	date_created timestamptz not null,
-- 	date_updated timestamptz not null,
-- 	-- slug text not null,

-- 	project_id bigint not null references project(id),
-- 	person_id bigint not null references person(id),

-- 	amount bigint not null,
-- 	state project_pledge_state_type default 'UNPAID' not null
-- );


-- grant select (slug, title, description, story, promises, category, goal, upload_images)
-- 	on table project_pledge to postgraphile_server_user, postgraphile_known_user;
-- grant insert (project_id, person_id, amount)
-- 	on table project_pledge to postgraphile_known_user;
-- grant delete
-- 	on table project_pledge to postgraphile_known_user;
-- grant usage
-- 	on sequence project_pledge_id_seq to postgraphile_known_user;

-- alter table project_pledge enable row level security;

-- create policy select_project_pledge on project_pledge for select
--   using (true);

-- create policy insert_project_pledge on project_pledge for insert to postgraphile_known_user
--   with check (person_id = current_person_id());

-- create policy update_project_pledge on project_pledge for update to postgraphile_known_user
--   using (person_id = current_person_id());

-- create policy delete_project_pledge on project_pledge for delete to postgraphile_known_user
--   using (person_id = current_person_id());



-- create trigger set_created_for_project_pledge
-- before insert on project_pledge
-- for each row
-- execute procedure trigger_set_created();

-- create trigger set_updated_for_project_pledge
-- before update on project_pledge
-- for each row
-- execute procedure trigger_set_updated();

-- create trigger _1_default_slug_for_project_pledge
-- before insert on project_pledge
-- for each row
-- execute procedure default_slug();



-- create table project_confirmation (
-- 	project_id bigint not null references project(id),
-- 	person_id bigint not null references person(id),
-- 	primary key (project_id, person_id),

-- 	date_created timestamptz not null,
-- 	date_updated timestamptz not null,

-- 	proceed boolean not null,
-- 	almost_promises text[] default '{}' not null,
-- 	-- fraudulent_flag boolean not null,
-- 	broken_promises text[] default '{}' not null,
-- 	commentary text,

-- 	check (
-- 		(proceed or cardinality(almost_promises) != 0) != (cardinality(broken_promises) != 0)
-- 	),
-- 	check (proceed != (char_length(commentary) != 0))
-- );

-- CREATE FUNCTION project_confirmation_fraudulent_flag(project_confirmation project_confirmation) RETURNS bool
-- AS $$
-- 	SELECT cardinality(project_confirmation.broken_promises) > 0
-- $$ LANGUAGE sql STABLE;


-- create trigger set_created_for_project_confirmation
-- before insert on project_confirmation
-- for each row
-- execute procedure trigger_set_created();

-- create trigger set_updated_for_project_confirmation
-- before update on project_confirmation
-- for each row
-- execute procedure trigger_set_updated();
