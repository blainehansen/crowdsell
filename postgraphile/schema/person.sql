create table person (
	id serial not null primary key,
	date_created timestamptz not null,
	date_updated timestamptz not null,

	slug text not null unique,
	url_slug text not null unique,
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

comment on table person is E'@omit create';

alter table person enable row level security;

grant select (slug, url_slug, name, bio, location, links, profile_photo_version)
	on table person to anonymous_user, logged_in_user;
create policy select_person on person for select
  using (true);

grant update (url_slug, name, bio, location, links)
	on table person to logged_in_user;
create policy update_person on person for update to logged_in_user
  using (id = current_setting('jwt.claims.id')::integer);

grant delete on table person to logged_in_user;
create policy delete_person on person for delete to logged_in_user
  using (id = current_setting('jwt.claims.id')::integer);


create function create_person(name text, email text, password text) returns person
as $$
declare
	new_person person;
begin
	insert into person (name, email, password) values
		(name, email, hash_password(password))
		returning * into new_person;

	return new_person;
end;
$$ language plpgsql strict security definer;

grant execute on function create_person(text, text, text) to anonymous_user;


create trigger set_created_for_person
before insert on person
for each row
execute procedure trigger_set_created();

create trigger set_updated_for_person
before update on person
for each row
execute procedure trigger_set_updated();

create trigger _1_default_slug_for_person
before insert on person
for each row
execute procedure default_slug();

create trigger _2_default_url_slug_for_person
before insert on person
for each row
execute procedure default_url_slug();

create trigger search_update_person_general_search_vector
before insert or update of name, bio on person
for each row
execute procedure tsvector_update_trigger(general_search_vector, 'pg_catalog.english', name, bio);

create index person_general_search_vector_idx on person using gin (general_search_vector);


create type jwt_token as (
  role text,
  id integer,
  exp integer
);

create function login(email text, password text) returns jwt_token
as $$
declare
  attempting_person person;
begin
  select a.* into attempting_person
  from person as a
  where a.email = $1;

  if verify_password(attempting_person.password, password) then
    return (
    	'logged_in_user',
    	attempting_person.person_id,
    	extract(epoch from now() + interval '2 hours')
  	)::jwt_token;
  else
    return null;
  end if;
end;
$$ language plpgsql strict security definer;

grant execute on function login(text, text) to anonymous_user, logged_in_user;


create function current_person() returns person
as $$
  select *
  from person
  where id = current_setting('jwt.claims.id')::integer
$$ language sql stable;

grant execute on function current_person() to anonymous_user, logged_in_user;

comment on function current_person() is E'@omit';
