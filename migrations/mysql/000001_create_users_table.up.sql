CREATE TABLE IF NOT EXISTS `users`
(
    `id` INT UNSIGNED AUTO_INCREMENT,
    `username` VARCHAR(45) NOT NULL,
    `email` VARCHAR(45) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `idx_unique_username` (`username` ASC) VISIBLE);