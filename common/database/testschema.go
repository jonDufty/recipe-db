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
}
