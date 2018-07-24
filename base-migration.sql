CREATE extension IF NOT EXISTS pg_hashids;

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
