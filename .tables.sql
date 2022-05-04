CREATE TABLE users_db.users (
	id BIGINT(20) auto_increment NOT NULL,
	first_name varchar(45) NULL,
	last_name varchar(45) NULL,
	email varchar(45) NOT NULL,
	date_created VARCHAR(45) NULL,
	PRIMARY KEY (`id`),
	UNIQUE INDEX `email_UNIQUE` (`email` ASC)
);