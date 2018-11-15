insert into person (name, email, password) values
	('Sterling Archer', 'sterling@gmail.com', 'pass'),
	('Malory Archer', 'malory@gmail.com', 'pass'),
	('Pamela Poovey', 'pamela@gmail.com', 'pass');

insert into project (person_id, title) values
	((select id from person limit 1), 'Stuff'),
	((select id from person limit 1), 'Different stuff'),
	((select id from person limit 1 offset 1), 'Whatever'),
	((select id from person limit 1 offset 1), 'Different whatever'),
	((select id from person limit 1 offset 2), 'Things'),
	((select id from person limit 1 offset 2), 'Different things');
