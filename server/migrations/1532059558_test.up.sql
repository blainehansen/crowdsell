BEGIN;

CREATE TABLE users (
	id serial NOT NULL PRIMARY KEY,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	slug text NOT NULL UNIQUE,
	internal_slug text NOT NULL UNIQUE,
	name text,
	email text NOT NULL UNIQUE,
	password bytea NOT NULL,
	profile_photo_slug text
);


CREATE TABLE projects (
	id serial NOT NULL PRIMARY KEY,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	slug text NOT NULL UNIQUE,
	internal_slug text NOT NULL UNIQUE,
	name text,
	description text,
	user_id bigint NOT NULL REFERENCES users(id)
);


COMMIT;
