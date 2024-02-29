CREATE TABLE IF NOT EXISTS `user` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `fullname` varchar(100) NOT NULL DEFAULT '',
    `phone_number` varchar(11) NOT NULL DEFAULT '',
    `email` varchar(100) NOT NULL DEFAULT '',
    `username` varchar(100) NOT NULL DEFAULT '',
    `campaign_id` bigint(20) unsigned NOT NULL DEFAULT 0,
    `status` tinyint NOT NULL DEFAULT 1,
    `login_attempt` tinyint NOT NULL DEFAULT 3,
    `checksum` bigint(20) unsigned NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    PRIMARY KEY (`id`),
    KEY `idx_phone_number` (`phone_number`),
    KEY `idx_email` (`email`),
    KEY `idx_username` (`username`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE IF NOT EXISTS `user_security` (
    `user_id` bigint(20) unsigned NOT NULL,
    `password` varbinary(60) NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    PRIMARY KEY (`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
