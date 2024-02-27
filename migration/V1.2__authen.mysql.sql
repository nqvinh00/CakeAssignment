CREATE TABLE `user` IF NOT EXISTS (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `fullname` varchar(100) NOT NULL DEFAULT '',
    `phone_number` varchar(11) NOT NULL DEFAULT '',
    `email` varchar(100) NOT NULL DEFAULT '',
    `username` varchar(100) NOT NULL DEFAULT '',
    `register_campaign` varchar(12) NOT NULL DEFAULT '',
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

CREATE TABLE `user_security` IF NOT EXISTS (
    `user_id` bigint(20) unsigned NOT NULL,
    `password` varbinary(60) NOT NULL,
    `checksum` bigint(20) unsigned NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    PRIMARY KEY (`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
