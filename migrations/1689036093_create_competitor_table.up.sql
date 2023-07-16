CREATE TABLE IF NOT EXISTS `competitor`
(
    `id`         int(11)     NOT NULL AUTO_INCREMENT,
    `name`       varchar(255)         DEFAULT NULL,
    `identifier` int(11)     NOT NULL,
    `offset_x`   int(11)     NOT NULL,
    `offset_y`   int(11)     NOT NULL,
    `score`      float(8, 2) NOT NULL,
    `created_at` datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;