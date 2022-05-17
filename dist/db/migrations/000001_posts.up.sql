CREATE TABLE IF NOT EXISTS `t_posts` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(256) NOT NULL DEFAULT '' COLLATE 'latin1_swedish_ci',
	`content` MEDIUMTEXT NOT NULL COLLATE 'latin1_swedish_ci',
	`ts` DATETIME NOT NULL DEFAULT current_timestamp(),
	`ts_updated` DATETIME NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
	PRIMARY KEY (`id`) USING BTREE,
	UNIQUE INDEX `name` (`name`) USING BTREE
)
COLLATE='latin1_swedish_ci'
ENGINE=InnoDB
;
