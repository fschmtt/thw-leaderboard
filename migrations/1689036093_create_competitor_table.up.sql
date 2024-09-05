CREATE TABLE IF NOT EXISTS `competitor`
(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) DEFAULT NULL,
    `identifier` int(11) NOT NULL,
    `target` decimal(10, 2) NOT NULL,
    `measurement` decimal(10, 3) NOT NULL,
    `offset` decimal(10, 3) NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `identifier` (`identifier`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
