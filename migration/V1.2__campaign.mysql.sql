CREATE TABLE `user_voucher` IF NOT EXISTS (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) unsigned NOT NULL,
    `campaign_id` bigint(20) unsigned NOT NULL,
    `voucher` varchar(8) NOT NULL DEFAULT '',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE `campaign` IF NOT EXISTS (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(100) NOT NULL DEFAULT '',
    `status` tinyint NOT NULL DEFAULT 1,
    `voucher_capacity` int unsigned NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
