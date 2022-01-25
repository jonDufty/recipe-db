CREATE TABLE IF NOT EXISTS `user` (
    `id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `email` varchar(255) NOT NULL,
    `time_created` time NOT NULL,
    `password_hash` varchar(255) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS`session` (
    `id` VARCHAR(255) NOT NULL,
    `user_id` int NOT NULL,
    `created_at` time NOT NULL,
    `expires_at` time NOT NULL,
    `ip` varchar(16),
    PRIMARY KEY (id)
);