CREATE TABLE IF NOT EXISTS `competitor`
(
    `id`           int(11)      NOT NULL AUTO_INCREMENT,
    `identifier`   varchar(255) NOT NULL,
    `name`         varchar(255) NOT NULL,
    `offset_x`     int(11)      NOT NULL,
    `offset_y`     int(11)      NOT NULL,
    `fault_points` int(11)      NOT NULL,
    `created_at`   datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;