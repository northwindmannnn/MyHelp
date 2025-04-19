create table admins (
	id SERIAL PRIMARY KEY,
	username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
	password TEXT UNIQUE NOT NULL,
	is_active BOOLEAN DEFAULT TRUE
);