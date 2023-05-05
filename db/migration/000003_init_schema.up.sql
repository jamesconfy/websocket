CREATE TABLE IF NOT EXISTS `token`(
    `_id` INT NOT NULL AUTO_INCREMENT UNIQUE,
    `user_id` VARCHAR(250) NOT NULL UNIQUE,
    `token` VARCHAR(6) NOT NULL,
    `token_id` VARCHAR(250) NOT NULL,
    `expiry` VARCHAR(250) NOT NULL,

    PRIMARY KEY(`token_id`),
    FOREIGN KEY(`user_id`) REFERENCES users(`user_id`) ON DELETE CASCADE
) 