CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	internal_slug TEXT NOT NULL,
	slug TEXT NOT NULL,

	name TEXT NOT NULL,
	-- https://stackoverflow.com/questions/386294/what-is-the-maximum-length-of-a-valid-email-address
	email VARCHAR(254) NOT NULL,
	password bytea NOT NULL,

	profile_photo_slug TEXT,

	UNIQUE(email),
	UNIQUE(slug)
);

CREATE TABLE projects (
	id SERIAL PRIMARY KEY,
	internal_slug TEXT NOT NULL,
	slug TEXT NOT NULL,

	title TEXT,

	UNIQUE(slug)
);

CREATE extension pg_hashids;

CREATE OR REPLACE FUNCTION hashid(BIGINT) RETURNS TEXT
AS $$
	SELECT id_encode($1, 'id& obfuscation sys$tem here', 8, 'abcdefghijklmnopqrstuvwxyz-ABCDEFGHIJKLMNOPQRSTUVWXYZ');
$$
LANGUAGE sql;

CREATE OR REPLACE FUNCTION unhashid(TEXT) RETURNS BIGINT
AS $$
	SELECT id_decode_once($1, 'id& obfuscation sys$tem here', 8, 'abcdefghijklmnopqrstuvwxyz-ABCDEFGHIJKLMNOPQRSTUVWXYZ');
$$
LANGUAGE sql;


CREATE OR REPLACE FUNCTION default_slug() RETURNS trigger AS
$$
BEGIN
	NEW.slug := hashid(NEW.id);
	NEW.internal_slug := NEW.slug;
	RETURN NEW;
END;
$$
LANGUAGE plpgsql;

CREATE TRIGGER default_slug_for_users
	BEFORE INSERT
	ON users
	FOR EACH ROW
	EXECUTE PROCEDURE default_slug();

CREATE TRIGGER default_slug_for_projects
	BEFORE INSERT
	ON projects
	FOR EACH ROW
	EXECUTE PROCEDURE default_slug();
