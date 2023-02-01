CREATE TABLE `user`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `username`  varchar(32) NOT NULL DEFAULT '' COMMENT 'UserName',
    `password`   varchar(32) NOT NULL DEFAULT '' COMMENT 'Password',
    `follow_count`  BIGINT         unsigned DEFAULT '0' COMMENT '关注数',
    `follower_count`    BIGINT         unsigned DEFAULT '0' COMMENT '粉丝数',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User account create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User account update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User account delete time',
    PRIMARY KEY (`id`),
    KEY          `idx_user_name` (`username`) COMMENT 'UserName index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User account table';
