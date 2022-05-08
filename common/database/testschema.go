package database

type Table struct {
	Table string
	Sql   string
}

type Database []Table

var Schema = Database{
	{
		"user",
		`CREATE TABLE IF NOT EXISTS user (
			id INT NOT NULL AUTO_INCREMENT,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			time_created DATETIME NOT NULL,
			password_hash VARCHAR(255) NOT NULL,
			PRIMARY KEY (id)
		);`,
	},
	{
		"session",
		`CREATE TABLE IF NOT EXISTS session (
			id VARCHAR(255) NOT NULL,
			user_id INT NOT NULL,
			created_at DATETIME NOT NULL,
			expires_at DATETIME NOT NULL,
			ip VARCHAR(16),
			PRIMARY KEY (id)
		);`,
	},
	{
		"recipes",
		`CREATE TABLE IF NOT EXISTS recipes (
		recipe_id INT NOT NULL AUTO_INCREMENT,
		title VARCHAR(255) NOT NULL,
		description VARCHAR(255),
		PRIMARY KEY (recipe_id)
		);`,
	},

	{
		"ingredient",
		`CREATE TABLE IF NOT EXISTS ingredient (
		ingredient_id INT NOT NULL AUTO_INCREMENT,
		label         VARCHAR(64) NOT NULL,
		PRIMARY KEY (ingredient_id)
		);`,
	},

	{
		"recipe_ingredient",
		`CREATE TABLE IF NOT EXISTS recipe_ingredient (
		recipe_ingredient_id INT NOT NULL AUTO_INCREMENT,
		recipe_id            INT NOT NULL,
		ingredient_id        INT NOT NULL,
		unit_id              INT NOT NULL,
		amount               DECIMAL(4, 2) DEFAULT NULL,
		PRIMARY KEY (recipe_ingredient_id)
		);`,
	},

	{
		"measurement_unit",
		`CREATE TABLE IF NOT EXISTS measurement_unit (
		unit_id INT NOT NULL AUTO_INCREMENT,
		label VARCHAR(64) NOT NULL,
		PRIMARY KEY (unit_id)
		);`,
	},

	{
		"instructions",
		`CREATE TABLE IF NOT EXISTS instructions (
		instruction_id INT NOT NULL AUTO_INCREMENT,
		step INT NOT NULL,
		text TEXT,
		recipe_id INT NOT NULL,
		PRIMARY KEY (instruction_id)
	);`,
	},
}
