CREATE TABLE IF NOT EXISTS `recipe` (
    `recipe_id` INT NOT NULL AUTO_INCREMENT,
    `title` VARCHAR(255) NOT NULL,
    `description` VARCHAR(255),
    PRIMARY KEY (recipe_id)
);

CREATE TABLE IF NOT EXISTS `unit` (
  `unit_id` INT NOT NULL AUTO_INCREMENT,
  `label`   VARCHAR(64) NOT NULL,
  PRIMARY KEY (unit_id),
);

CREATE TABLE IF NOT EXISTS `ingredient` (
  `ingredient_id` INT NOT NULL AUTO_INCREMENT,
  `label`         VARCHAR(64) NOT NULL,
  PRIMARY KEY (ingredient_id)
);

CREATE TABLE IF NOT EXISTS `recipe_ingredient` (
  `recipe_ingredient_id` INT NOT NULL AUTO_INCREMENT,
  `recipe_id`            INT NOT NULL,
  `ingredient_id`        INT NOT NULL,
  `unit_id`              INT NOT NULL,
  `amount`               DECIMAL(4, 2) DEFAULT NULL,
  PRIMARY KEY (recipe_ingredient_id)
);

