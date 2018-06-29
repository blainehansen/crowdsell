-- Your SQL goes here
CREATE TABLE projects (
	id SERIAL PRIMARY KEY,
	title VARCHAR,
);

CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	name VARCHAR NOT NULL,
	-- https://stackoverflow.com/questions/386294/what-is-the-maximum-length-of-a-valid-email-address
	email VARCHAR(254) NOT NULL,
	slug VARCHAR NOT NULL,
	password VARCHAR NOT NULL,

	hash_id VARCHAR,
	profile_photo_hash VARCHAR,

	UNIQUE(email),
	UNIQUE(slug),
);
