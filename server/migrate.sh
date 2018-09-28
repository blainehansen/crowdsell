source ../.env.dev.sh

PGPASSWORD=$DATABASE_PASSWORD psql -U $DATABASE_USER -h $SYSTEM_DATABASE_HOST $DATABASE_DB_NAME << EOF

CREATE extension IF NOT EXISTS pg_hashids;

CREATE OR REPLACE FUNCTION hashid(BIGINT) RETURNS TEXT
AS \$$
  SELECT id_encode(\$1, '$HASHID_SALT', $HASHID_MIN_LENGTH, '$HASHID_ALPHABET');
\$$ LANGUAGE sql;

CREATE OR REPLACE FUNCTION unhashid(TEXT) RETURNS BIGINT
AS \$$
  SELECT id_decode_once(\$1, '$HASHID_SALT', $HASHID_MIN_LENGTH, '$HASHID_ALPHABET');
\$$ LANGUAGE sql;

CREATE OR REPLACE FUNCTION trigger_set_created() RETURNS TRIGGER
AS \$$
BEGIN
  NEW.date_created = NOW();
  NEW.date_updated = NOW();
  RETURN NEW;
END;
\$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION trigger_set_updated() RETURNS TRIGGER
AS \$$
BEGIN
  NEW.date_updated = NOW();
  RETURN NEW;
END;
\$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION default_slug() RETURNS trigger
AS \$$
BEGIN
  NEW.slug := hashid(NEW.id);
  RETURN NEW;
END;
\$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION default_url_slug() RETURNS trigger
AS \$$
BEGIN
  NEW.url_slug := NEW.slug;
  RETURN NEW;
END;
\$$ LANGUAGE plpgsql;

CREATE TYPE project_pledges_state_enum AS ENUM (
	'UNPAID',
	'PAID',
	'RELEASED'
);

BEGIN;

CREATE TABLE users (
	id serial NOT NULL PRIMARY KEY,
	date_created timestamptz NOT NULL,
	date_updated timestamptz NOT NULL,
	slug text NOT NULL,
	url_slug text NOT NULL,
	name text,
	bio text,
	location text,
	links text,
	email text NOT NULL UNIQUE,
	has_payment_user boolean DEFAULT false NOT NULL,
	password bytea NOT NULL,
	profile_photo_version text,
	forgot_password_token bytea,
	general_search_vector tsvector
);

CREATE TRIGGER set_created_for_users
BEFORE INSERT ON users
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_created();

CREATE TRIGGER set_updated_for_users
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_updated();

CREATE TRIGGER _1_default_slug_for_users
BEFORE INSERT ON users
FOR EACH ROW
EXECUTE PROCEDURE default_slug();

CREATE TRIGGER _2_default_url_slug_for_users
BEFORE INSERT ON users
FOR EACH ROW
EXECUTE PROCEDURE default_url_slug();

CREATE TRIGGER search_update_users_general_search_vector
BEFORE INSERT OR UPDATE OF name, bio ON users
FOR EACH ROW
EXECUTE PROCEDURE tsvector_update_trigger(general_search_vector, 'pg_catalog.english', name, bio);

CREATE INDEX users_general_search_vector_idx ON users USING gin (general_search_vector);

CREATE TABLE projects (
	id serial NOT NULL PRIMARY KEY,
	date_created timestamptz NOT NULL,
	date_updated timestamptz NOT NULL,
	slug text NOT NULL,
	url_slug text NOT NULL,
	name text,
	description text,
	story text,
	promises text[] DEFAULT ARRAY[]::text[] NOT NULL,
	upload_images text[] DEFAULT ARRAY[]::text[] NOT NULL,
	user_id bigint NOT NULL REFERENCES users(id),
	general_search_vector tsvector
);

CREATE TRIGGER set_created_for_projects
BEFORE INSERT ON projects
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_created();

CREATE TRIGGER set_updated_for_projects
BEFORE UPDATE ON projects
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_updated();

CREATE TRIGGER _1_default_slug_for_projects
BEFORE INSERT ON projects
FOR EACH ROW
EXECUTE PROCEDURE default_slug();

CREATE TRIGGER _2_default_url_slug_for_projects
BEFORE INSERT ON projects
FOR EACH ROW
EXECUTE PROCEDURE default_url_slug();

CREATE TRIGGER search_update_projects_general_search_vector
BEFORE INSERT OR UPDATE OF name, description ON projects
FOR EACH ROW
EXECUTE PROCEDURE tsvector_update_trigger(general_search_vector, 'pg_catalog.english', name, description);

CREATE INDEX projects_general_search_vector_idx ON projects USING gin (general_search_vector);

CREATE TABLE project_pledges (
	id serial NOT NULL PRIMARY KEY,
	date_created timestamptz NOT NULL,
	date_updated timestamptz NOT NULL,
	slug text NOT NULL,
	project_id bigint NOT NULL REFERENCES projects(id),
	user_id bigint NOT NULL REFERENCES users(id),
	amount bigint NOT NULL,
	state project_pledges_state_enum DEFAULT 'UNPAID' NOT NULL
);

CREATE TRIGGER set_created_for_project_pledges
BEFORE INSERT ON project_pledges
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_created();

CREATE TRIGGER set_updated_for_project_pledges
BEFORE UPDATE ON project_pledges
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_updated();

CREATE TRIGGER _1_default_slug_for_project_pledges
BEFORE INSERT ON project_pledges
FOR EACH ROW
EXECUTE PROCEDURE default_slug();


CREATE TABLE project_confirmations (
	id serial NOT NULL PRIMARY KEY,
	date_created timestamptz NOT NULL,
	date_updated timestamptz NOT NULL,
	slug text NOT NULL,
	project_id bigint NOT NULL REFERENCES projects(id),
	user_id bigint NOT NULL REFERENCES users(id),
	proceed boolean NOT NULL,
	almost_promises text[] DEFAULT ARRAY[]::text[] NOT NULL,
	fraudulent_flag boolean NOT NULL,
	broken_promises text[] DEFAULT ARRAY[]::text[] NOT NULL,
	commentary text,
	CONSTRAINT project_confirmations_unique_project_user unique (project_id, user_id),
check (
  (proceed OR cardinality(almost_promises) != 0) != (fraudulent_flag OR cardinality(broken_promises) != 0)
),
check (proceed != (char_length(commentary) != 0))

);

CREATE TRIGGER set_created_for_project_confirmations
BEFORE INSERT ON project_confirmations
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_created();

CREATE TRIGGER set_updated_for_project_confirmations
BEFORE UPDATE ON project_confirmations
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_updated();

CREATE TRIGGER _1_default_slug_for_project_confirmations
BEFORE INSERT ON project_confirmations
FOR EACH ROW
EXECUTE PROCEDURE default_slug();


COMMIT;


EOF