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

CREATE TABLE `video`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `playurl`    varchar(255) NOT NULL DEFAULT '' COMMENT 'PlayUrl',
    `coverurl`   varchar(255) NOT NULL DEFAULT '' COMMENT 'CoverUrl',
    `favorite_count`    bigint unsigned DEFAULT 0 COMMENT 'FavoriteCount',
    `comment_count` bigint unsigned DEFAULT 0 COMMENT 'CommentCount',
    `title`      varchar(255) NOT NULL DEFAULT '' COMMENT 'Title',
    `author_id`  bigint unsigned NOT NULL COMMENT 'AuthorId',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Video publish time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Video update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'Video delete time',
    PRIMARY KEY (`id`),
    KEY          `author_id` (`author_id`) COMMENT 'AuthorId index',
    KEY          `created_at`(`created_at`) COMMENT `Video publish time index` 
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Video table';

CREATE TABLE `favorite`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `user_id`    bigint unsigned NOT NULL COMMENT 'User_Id',
    `video_id`   bigint unsigned NOT NULL COMMENT `Video_Id`,
    `status`     tinyint(1) DEFAULT '0' COMMENT 'favorite status(0favorite 1cancel)',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'favorite action time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'favorite status change time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'favorite cancel time',
    PRIMARY KEY (`id`),
    UNIQUE KEY  `user_video_index`(`user_id`, `video_id`) COMMENT 'UserId_VideoId_Index',
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='favorite table';