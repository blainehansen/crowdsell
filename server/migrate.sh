source ./.env
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
  NEW.url_slug := NEW.slug;
  RETURN NEW;
END;
\$$ LANGUAGE plpgsql;


BEGIN;

CREATE TABLE users (
	id serial NOT NULL PRIMARY KEY,
	date_created timestamptz,
	date_updated timestamptz,
	slug text,
	url_slug text,
	name text,
	email text NOT NULL UNIQUE,
	password bytea NOT NULL,
	profile_photo_slug text,
	forgot_password_token bytea
);

CREATE TRIGGER set_created_for_users
BEFORE INSERT ON users
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_created();

CREATE TRIGGER set_updated_for_users
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_updated();

CREATE TRIGGER default_slug_for_users
BEFORE INSERT ON users
FOR EACH ROW
EXECUTE PROCEDURE default_slug();


CREATE TABLE projects (
	id serial NOT NULL PRIMARY KEY,
	date_created timestamptz,
	date_updated timestamptz,
	slug text,
	url_slug text,
	name text,
	description text,
	user_id bigint NOT NULL REFERENCES users(id)
);

CREATE TRIGGER set_created_for_projects
BEFORE INSERT ON projects
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_created();

CREATE TRIGGER set_updated_for_projects
BEFORE UPDATE ON projects
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_updated();

CREATE TRIGGER default_slug_for_projects
BEFORE INSERT ON projects
FOR EACH ROW
EXECUTE PROCEDURE default_slug();


COMMIT;


EOF