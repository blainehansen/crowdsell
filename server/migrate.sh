source ./.env

PGPASSWORD=$DATABASE_PASSWORD psql -U $DATABASE_USER -h $DATABASE_HOST $DATABASE_DB_NAME << EOF
BEGIN;

CREATE TABLE users (
	id serial NOT NULL PRIMARY KEY,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	slug text NOT NULL UNIQUE,
	url_slug text NOT NULL UNIQUE,
	name text,
	email text NOT NULL UNIQUE,
	password bytea NOT NULL,
	profile_photo_slug text,
	forgot_password_token text
);


CREATE TABLE projects (
	id serial NOT NULL PRIMARY KEY,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	slug text NOT NULL UNIQUE,
	url_slug text NOT NULL UNIQUE,
	name text,
	description text,
	user_id bigint NOT NULL REFERENCES users(id)
);

COMMIT;


CREATE extension IF NOT EXISTS pg_hashids;

CREATE OR REPLACE FUNCTION hashid(BIGINT) RETURNS TEXT
AS \$$
	SELECT id_encode(\$1, '$HASHID_SALT', $HASHID_MIN_LENGTH, '$HASHID_ALPHABET');
\$$
LANGUAGE sql;

CREATE OR REPLACE FUNCTION unhashid(TEXT) RETURNS BIGINT
AS \$$
	SELECT id_decode_once(\$1, '$HASHID_SALT', $HASHID_MIN_LENGTH, '$HASHID_ALPHABET');
\$$
LANGUAGE sql;


CREATE OR REPLACE FUNCTION default_slug() RETURNS trigger AS
\$$
BEGIN
	NEW.slug := hashid(NEW.id);
	NEW.url_slug := NEW.slug;
	RETURN NEW;
END;
\$$
LANGUAGE plpgsql;

CREATE TRIGGER default_slug_for_user
	BEFORE INSERT
	ON users
	FOR EACH ROW
	EXECUTE PROCEDURE default_slug();

CREATE TRIGGER default_slug_for_project
	BEFORE INSERT
	ON projects
	FOR EACH ROW
	EXECUTE PROCEDURE default_slug();
EOF
