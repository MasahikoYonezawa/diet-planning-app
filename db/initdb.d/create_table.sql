DROP TABLE IF EXISTS `test_database`.`users`;
CREATE TABLE IF NOT EXISTS `test_database`.`users` (
    `id` BIGINT AUTO_INCREMENT,
    `name` VARCHAR(100) NOT NULL,
    `gender` TINYINT NOT NULL,
    `age` INT NOT NULL,
    `height` FLOAT NOT NULL,
    `weight` FLOAT NOT NULL,
    `activity_level` TINYINT NOT NULL DEFAULT 0,
    `bmr` INT NOT NULL DEFAULT 0,
    `tdee` INT NOT NULL DEFAULT 0,
    `target_weight` FLOAT NOT NULL DEFAULT 0.00,
    `term` INT NOT NULL DEFAULT 0,
    `term_type` TINYINT NOT NULL DEFAULT 0,
    `protein` FLOAT NOT NULL DEFAULT 0.00,
    `fat` FLOAT NOT NULL DEFAULT 0.00,
    `carbohydrate` FLOAT DEFAULT 0.00,
    `created_at` TIMESTAMP NOT NULL DEFAULT NOW(),
    `updated_at` TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
INSERT INTO `test_database`.`users` (`name`, `gendar`, `age`, `height`, `weight`, `activity_level`) VALUES ("test user", 1, 30, 180, 80, 1);