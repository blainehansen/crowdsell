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
	UNIQUE(email),
	UNIQUE(slug),
);
