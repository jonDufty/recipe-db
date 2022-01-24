CREATE TABLE `user` (
    `id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `email` varchar(255) NOT NULL,
    `time_created` time NOT NULL,
    `password_hash` varchar(255) NOT NULL, 
    PRIMARY KEY (id)
    );

CREATE TABLE `session` (
    `id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `email` varchar(255) NOT NULL,
    `time_created` time NOT NULL,
    `password_hash` varchar(255) NOT NULL, 
    PRIMARY KEY (id)
);