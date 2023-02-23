/*
 Navicat Premium Data Transfer

 Source Server         : sql
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : localhost:9910
 Source Schema         : tiktok

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 23/02/2023 14:01:54
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `video_id` bigint NULL DEFAULT NULL,
  `user_id` bigint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_comment_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7034396283762900993 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of comment
-- ----------------------------
INSERT INTO `comment` VALUES (7034386098705137664, '2023-02-23 12:58:58.505', '2023-02-23 12:58:58.505', NULL, '测试评论', 13, 9);
INSERT INTO `comment` VALUES (7034389109783396352, '2023-02-23 13:10:56.402', '2023-02-23 13:10:56.402', NULL, '测试评论', 14, 9);
INSERT INTO `comment` VALUES (7034391244768018432, '2023-02-23 13:19:25.423', '2023-02-23 13:19:25.423', NULL, '测试评论', 15, 9);
INSERT INTO `comment` VALUES (7034393488053764096, '2023-02-23 13:28:20.264', '2023-02-23 13:28:20.264', NULL, '测试评论', 15, 9);
INSERT INTO `comment` VALUES (7034393646321631232, '2023-02-23 13:28:57.998', '2023-02-23 13:28:57.998', NULL, '测试评论', 15, 9);
INSERT INTO `comment` VALUES (7034394014153703424, '2023-02-23 13:30:25.696', '2023-02-23 13:30:25.696', NULL, '测试评论', 15, 9);
INSERT INTO `comment` VALUES (7034394410305716224, '2023-02-23 13:32:00.146', '2023-02-23 13:32:00.146', NULL, '测试评论', 15, 9);
INSERT INTO `comment` VALUES (7034394441054158848, '2023-02-23 13:32:07.477', '2023-02-23 13:32:07.477', NULL, '测试评论', 15, 9);
INSERT INTO `comment` VALUES (7034394872442519552, '2023-02-23 13:33:50.328', '2023-02-23 13:33:50.328', NULL, '测试评论', 15, 9);
INSERT INTO `comment` VALUES (7034395095176839168, '2023-02-23 13:34:43.431', '2023-02-23 13:34:43.431', '2023-02-23 13:34:43.467', '测试评论', 15, 9);
INSERT INTO `comment` VALUES (7034395359204081664, '2023-02-23 13:35:46.381', '2023-02-23 13:35:46.381', '2023-02-23 13:35:46.405', '测试评论', 15, 9);
INSERT INTO `comment` VALUES (7034395389495345152, '2023-02-23 13:35:53.602', '2023-02-23 13:35:53.602', '2023-02-23 13:35:53.621', '测试评论', 16, 9);
INSERT INTO `comment` VALUES (7034395843478421504, '2023-02-23 13:37:41.841', '2023-02-23 13:37:41.841', '2023-02-23 13:37:41.857', '测试评论', 17, 9);
INSERT INTO `comment` VALUES (7034395851074306048, '2023-02-23 13:37:43.652', '2023-02-23 13:37:43.652', '2023-02-23 13:37:43.669', '测试评论', 18, 9);
INSERT INTO `comment` VALUES (7034395857843912704, '2023-02-23 13:37:45.266', '2023-02-23 13:37:45.266', '2023-02-23 13:37:45.281', '测试评论', 19, 9);
INSERT INTO `comment` VALUES (7034395865116835840, '2023-02-23 13:37:47.000', '2023-02-23 13:37:47.000', '2023-02-23 13:37:47.014', '测试评论', 20, 9);
INSERT INTO `comment` VALUES (7034395872389758976, '2023-02-23 13:37:48.733', '2023-02-23 13:37:48.733', '2023-02-23 13:37:48.747', '测试评论', 21, 9);
INSERT INTO `comment` VALUES (7034395880451211264, '2023-02-23 13:37:50.656', '2023-02-23 13:37:50.656', '2023-02-23 13:37:50.671', '测试评论', 22, 9);
INSERT INTO `comment` VALUES (7034395887912878080, '2023-02-23 13:37:52.434', '2023-02-23 13:37:52.434', '2023-02-23 13:37:52.448', '测试评论', 23, 9);
INSERT INTO `comment` VALUES (7034395895148052480, '2023-02-23 13:37:54.160', '2023-02-23 13:37:54.160', '2023-02-23 13:37:54.175', '测试评论', 24, 9);
INSERT INTO `comment` VALUES (7034395902609719296, '2023-02-23 13:37:55.938', '2023-02-23 13:37:55.938', '2023-02-23 13:37:55.953', '测试评论', 25, 9);
INSERT INTO `comment` VALUES (7034395909761007616, '2023-02-23 13:37:57.643', '2023-02-23 13:37:57.643', '2023-02-23 13:37:57.657', '测试评论', 26, 9);
INSERT INTO `comment` VALUES (7034395917004570624, '2023-02-23 13:37:59.370', '2023-02-23 13:37:59.370', '2023-02-23 13:37:59.385', '测试评论', 27, 9);
INSERT INTO `comment` VALUES (7034396205035814912, '2023-02-23 13:39:08.043', '2023-02-23 13:39:08.043', '2023-02-23 13:39:08.059', '测试评论', 29, 9);
INSERT INTO `comment` VALUES (7034396213147598848, '2023-02-23 13:39:09.976', '2023-02-23 13:39:09.976', '2023-02-23 13:39:09.992', '测试评论', 29, 9);
INSERT INTO `comment` VALUES (7034396221381017600, '2023-02-23 13:39:11.940', '2023-02-23 13:39:11.940', '2023-02-23 13:39:11.956', '测试评论', 30, 9);
INSERT INTO `comment` VALUES (7034396229136285696, '2023-02-23 13:39:13.788', '2023-02-23 13:39:13.788', '2023-02-23 13:39:13.807', '测试评论', 31, 9);
INSERT INTO `comment` VALUES (7034396237273235456, '2023-02-23 13:39:15.729', '2023-02-23 13:39:15.729', '2023-02-23 13:39:15.743', '测试评论', 32, 9);
INSERT INTO `comment` VALUES (7034396244843954176, '2023-02-23 13:39:17.534', '2023-02-23 13:39:17.534', '2023-02-23 13:39:17.550', '测试评论', 33, 9);
INSERT INTO `comment` VALUES (7034396252599222272, '2023-02-23 13:39:19.383', '2023-02-23 13:39:19.383', '2023-02-23 13:39:19.397', '测试评论', 34, 9);
INSERT INTO `comment` VALUES (7034396260757143552, '2023-02-23 13:39:21.328', '2023-02-23 13:39:21.328', '2023-02-23 13:39:21.343', '测试评论', 35, 9);
INSERT INTO `comment` VALUES (7034396268336250880, '2023-02-23 13:39:23.135', '2023-02-23 13:39:23.135', '2023-02-23 13:39:23.151', '测试评论', 36, 9);
INSERT INTO `comment` VALUES (7034396275948912640, '2023-02-23 13:39:24.949', '2023-02-23 13:39:24.949', '2023-02-23 13:39:24.965', '测试评论', 37, 9);
INSERT INTO `comment` VALUES (7034396283762900992, '2023-02-23 13:39:26.813', '2023-02-23 13:39:26.813', '2023-02-23 13:39:26.827', '测试评论', 38, 9);

-- ----------------------------
-- Table structure for favorite
-- ----------------------------
DROP TABLE IF EXISTS `favorite`;
CREATE TABLE `favorite`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `user_id` bigint NULL DEFAULT NULL,
  `video_id` bigint NULL DEFAULT NULL,
  `status` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_favorite_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 42 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of favorite
-- ----------------------------
INSERT INTO `favorite` VALUES (6, '2023-02-22 13:56:13.677', '2023-02-22 13:56:13.677', NULL, 3, 11, 0);
INSERT INTO `favorite` VALUES (7, '2023-02-22 13:58:20.068', '2023-02-22 13:58:20.068', NULL, 4, 11, 0);
INSERT INTO `favorite` VALUES (8, '2023-02-22 13:58:23.352', '2023-02-22 13:58:23.352', NULL, 4, 10, 0);
INSERT INTO `favorite` VALUES (9, '2023-02-22 14:01:43.903', '2023-02-22 14:01:43.903', NULL, 5, 10, 0);
INSERT INTO `favorite` VALUES (10, '2023-02-22 14:01:46.368', '2023-02-22 14:01:46.368', NULL, 5, 11, 0);
INSERT INTO `favorite` VALUES (11, '2023-02-22 14:02:21.782', '2023-02-22 14:02:21.782', NULL, 5, 12, 0);
INSERT INTO `favorite` VALUES (12, '2023-02-22 14:07:17.753', '2023-02-22 14:07:17.753', NULL, 6, 10, 0);
INSERT INTO `favorite` VALUES (13, '2023-02-22 14:07:19.918', '2023-02-22 14:07:19.918', NULL, 6, 11, 0);
INSERT INTO `favorite` VALUES (14, '2023-02-22 14:07:21.685', '2023-02-22 14:07:21.685', NULL, 6, 12, 0);
INSERT INTO `favorite` VALUES (15, '2023-02-22 14:07:58.402', '2023-02-22 14:07:58.402', NULL, 7, 12, 0);
INSERT INTO `favorite` VALUES (16, '2023-02-23 12:58:58.457', '2023-02-23 13:13:50.842', NULL, 9, 13, 1);
INSERT INTO `favorite` VALUES (17, '2023-02-23 13:10:56.362', '2023-02-23 13:10:56.362', NULL, 9, 14, 0);
INSERT INTO `favorite` VALUES (18, '2023-02-23 13:16:28.792', '2023-02-23 13:16:41.907', NULL, 9, 15, 1);
INSERT INTO `favorite` VALUES (19, '2023-02-23 13:35:53.551', '2023-02-23 13:35:53.551', NULL, 9, 16, 0);
INSERT INTO `favorite` VALUES (20, '2023-02-23 13:37:41.808', '2023-02-23 13:37:41.808', NULL, 9, 17, 0);
INSERT INTO `favorite` VALUES (21, '2023-02-23 13:37:43.613', '2023-02-23 13:37:43.613', NULL, 9, 18, 0);
INSERT INTO `favorite` VALUES (22, '2023-02-23 13:37:45.233', '2023-02-23 13:37:45.233', NULL, 9, 19, 0);
INSERT INTO `favorite` VALUES (23, '2023-02-23 13:37:46.966', '2023-02-23 13:37:46.966', NULL, 9, 20, 0);
INSERT INTO `favorite` VALUES (24, '2023-02-23 13:37:48.699', '2023-02-23 13:37:48.699', NULL, 9, 21, 0);
INSERT INTO `favorite` VALUES (25, '2023-02-23 13:37:50.620', '2023-02-23 13:37:50.620', NULL, 9, 22, 0);
INSERT INTO `favorite` VALUES (26, '2023-02-23 13:37:52.397', '2023-02-23 13:37:52.397', NULL, 9, 23, 0);
INSERT INTO `favorite` VALUES (27, '2023-02-23 13:37:54.123', '2023-02-23 13:37:54.123', NULL, 9, 24, 0);
INSERT INTO `favorite` VALUES (28, '2023-02-23 13:37:55.896', '2023-02-23 13:37:55.896', NULL, 9, 25, 0);
INSERT INTO `favorite` VALUES (29, '2023-02-23 13:37:57.604', '2023-02-23 13:37:57.604', NULL, 9, 26, 0);
INSERT INTO `favorite` VALUES (30, '2023-02-23 13:37:59.328', '2023-02-23 13:37:59.328', NULL, 9, 27, 0);
INSERT INTO `favorite` VALUES (31, '2023-02-23 13:39:07.988', '2023-02-23 13:39:07.988', NULL, 9, 28, 0);
INSERT INTO `favorite` VALUES (32, '2023-02-23 13:39:09.930', '2023-02-23 13:39:09.930', NULL, 9, 29, 0);
INSERT INTO `favorite` VALUES (33, '2023-02-23 13:39:11.881', '2023-02-23 13:39:11.881', NULL, 9, 30, 0);
INSERT INTO `favorite` VALUES (34, '2023-02-23 13:39:13.736', '2023-02-23 13:39:13.736', NULL, 9, 31, 0);
INSERT INTO `favorite` VALUES (35, '2023-02-23 13:39:15.682', '2023-02-23 13:39:15.682', NULL, 9, 32, 0);
INSERT INTO `favorite` VALUES (36, '2023-02-23 13:39:17.484', '2023-02-23 13:39:17.484', NULL, 9, 33, 0);
INSERT INTO `favorite` VALUES (37, '2023-02-23 13:39:19.332', '2023-02-23 13:39:19.332', NULL, 9, 34, 0);
INSERT INTO `favorite` VALUES (38, '2023-02-23 13:39:21.272', '2023-02-23 13:39:21.272', NULL, 9, 35, 0);
INSERT INTO `favorite` VALUES (39, '2023-02-23 13:39:23.078', '2023-02-23 13:39:23.078', NULL, 9, 36, 0);
INSERT INTO `favorite` VALUES (40, '2023-02-23 13:39:24.895', '2023-02-23 13:39:24.895', NULL, 9, 37, 0);
INSERT INTO `favorite` VALUES (41, '2023-02-23 13:39:26.762', '2023-02-23 13:39:26.762', NULL, 9, 38, 0);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `username` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `follow_count` bigint UNSIGNED NOT NULL,
  `follower_count` bigint UNSIGNED NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username` ASC) USING BTREE,
  INDEX `idx_user_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 35 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (3, '2023-02-22 13:53:00.572', '2023-02-22 13:53:00.572', NULL, '777', 'f63f4fbc9f8c85d409f2f59f2b9e12d5', 0, 0);
INSERT INTO `user` VALUES (4, '2023-02-22 13:58:09.061', '2023-02-22 13:58:09.061', NULL, '123', '4297f44b13955235245b2497399d7a93', 0, 0);
INSERT INTO `user` VALUES (5, '2023-02-22 14:01:38.397', '2023-02-22 14:01:38.397', NULL, '123456', 'e10adc3949ba59abbe56e057f20f883e', 0, 0);
INSERT INTO `user` VALUES (6, '2023-02-22 14:06:24.352', '2023-02-22 14:06:24.352', NULL, 'llc', '37cf3a6a04dd419a568236dc9dc93930', 0, 0);
INSERT INTO `user` VALUES (7, '2023-02-22 14:07:47.995', '2023-02-22 14:07:47.995', NULL, 'rrr', 'ff2f24f8b6d253bb5a8bc55728ca7372', 0, 0);
INSERT INTO `user` VALUES (8, '2023-02-23 12:58:57.421', '2023-02-23 12:58:57.421', NULL, 'douyin45324', '5998bff2eebc82a40008eac744a79d56', 0, 0);
INSERT INTO `user` VALUES (9, '2023-02-23 12:58:57.443', '2023-02-23 12:58:57.443', NULL, 'douyinTestUserA', 'b6fb5b170ed89a5a366e83a635abb270', 0, 0);
INSERT INTO `user` VALUES (10, '2023-02-23 13:10:55.914', '2023-02-23 13:10:55.914', NULL, 'douyin50924', '22deb310d6375a9591511e52c58eb271', 0, 0);
INSERT INTO `user` VALUES (11, '2023-02-23 13:35:45.793', '2023-02-23 13:35:45.793', NULL, 'douyin38434', '91ee66acd77114e88d47effb4933eedd', 0, 0);
INSERT INTO `user` VALUES (12, '2023-02-23 13:35:52.728', '2023-02-23 13:35:52.728', NULL, 'douyin55635', '36c1ff554903c474421e53ec21396821', 0, 0);
INSERT INTO `user` VALUES (13, '2023-02-23 13:37:41.322', '2023-02-23 13:37:41.322', NULL, 'douyin18844', '09c7b5d32d597cbff043f40446172d5b', 0, 0);
INSERT INTO `user` VALUES (14, '2023-02-23 13:37:43.207', '2023-02-23 13:37:43.207', NULL, 'douyin26206', '3542cfd0e760db9c8bc9b5685eff811d', 0, 0);
INSERT INTO `user` VALUES (15, '2023-02-23 13:37:44.866', '2023-02-23 13:37:44.866', NULL, 'douyin59135', 'c583c16c2fff240649d0a016a6e20945', 0, 0);
INSERT INTO `user` VALUES (16, '2023-02-23 13:37:46.633', '2023-02-23 13:37:46.633', NULL, 'douyin8601', 'ebb00929aef4c2b884d63c30ec12e2c5', 0, 0);
INSERT INTO `user` VALUES (17, '2023-02-23 13:37:48.335', '2023-02-23 13:37:48.335', NULL, 'douyin26084', '7aea62fe0ada8d8d4cdbd2abf3686db1', 0, 0);
INSERT INTO `user` VALUES (18, '2023-02-23 13:37:50.221', '2023-02-23 13:37:50.221', NULL, 'douyin58493', '8432f70737c28360ff7debb422f1f93e', 0, 0);
INSERT INTO `user` VALUES (19, '2023-02-23 13:37:51.946', '2023-02-23 13:37:51.946', NULL, 'douyin14524', 'd1039705da53eeac4c4d0fd1640b7f3b', 0, 0);
INSERT INTO `user` VALUES (20, '2023-02-23 13:37:53.669', '2023-02-23 13:37:53.669', NULL, 'douyin15917', 'ad8ba642b8a5ad6c5b32ab4d5f744f1f', 0, 0);
INSERT INTO `user` VALUES (21, '2023-02-23 13:37:55.533', '2023-02-23 13:37:55.533', NULL, 'douyin20812', 'fd2f2e0bd0a21f7b93329790c5e2fd92', 0, 0);
INSERT INTO `user` VALUES (22, '2023-02-23 13:37:57.260', '2023-02-23 13:37:57.260', NULL, 'douyin62399', '903d49b59c8235b86f25fe0e7a4bb9c7', 0, 0);
INSERT INTO `user` VALUES (23, '2023-02-23 13:37:58.970', '2023-02-23 13:37:58.970', NULL, 'douyin46204', 'dbe82ba48fa95edf3e1194ce07c9cfd8', 0, 0);
INSERT INTO `user` VALUES (24, '2023-02-23 13:39:07.476', '2023-02-23 13:39:07.476', NULL, 'douyin42952', '59f5bee5e254ed996e0be4dd444cf1da', 0, 0);
INSERT INTO `user` VALUES (25, '2023-02-23 13:39:09.562', '2023-02-23 13:39:09.562', NULL, 'douyin34524', '48157bd30592416c1b6da34923d5e2b6', 0, 0);
INSERT INTO `user` VALUES (26, '2023-02-23 13:39:11.434', '2023-02-23 13:39:11.434', NULL, 'douyin24711', '43ab4c09a53c327cac29bfcb02d5cfca', 0, 0);
INSERT INTO `user` VALUES (27, '2023-02-23 13:39:13.342', '2023-02-23 13:39:13.342', NULL, 'douyin37453', 'fb07fe8649d22efa445503a6d16225f4', 0, 0);
INSERT INTO `user` VALUES (28, '2023-02-23 13:39:15.214', '2023-02-23 13:39:15.214', NULL, 'douyin42279', '3c9961a497084525a0992406beff1180', 0, 0);
INSERT INTO `user` VALUES (29, '2023-02-23 13:39:17.083', '2023-02-23 13:39:17.083', NULL, 'douyin27401', 'd91db8f52b92f4d2a6d494754962c1c7', 0, 0);
INSERT INTO `user` VALUES (30, '2023-02-23 13:39:18.955', '2023-02-23 13:39:18.955', NULL, 'douyin48495', 'cbb1cfe2aa94a86c24506060716ca91c', 0, 0);
INSERT INTO `user` VALUES (31, '2023-02-23 13:39:20.823', '2023-02-23 13:39:20.823', NULL, 'douyin23985', 'a4e41df51dd61f61631d61c5a41bef97', 0, 0);
INSERT INTO `user` VALUES (32, '2023-02-23 13:39:22.675', '2023-02-23 13:39:22.675', NULL, 'douyin45508', '748beff87d57afbe832cef985defdfa9', 0, 0);
INSERT INTO `user` VALUES (33, '2023-02-23 13:39:24.538', '2023-02-23 13:39:24.538', NULL, 'douyin40719', '446e6b3ee9f81d5e01e29d4c433ae3bc', 0, 0);
INSERT INTO `user` VALUES (34, '2023-02-23 13:39:26.374', '2023-02-23 13:39:26.374', NULL, 'douyin30863', 'e29e7b8af44d25c64ba1eabbab6084cc', 0, 0);

-- ----------------------------
-- Table structure for video
-- ----------------------------
DROP TABLE IF EXISTS `video`;
CREATE TABLE `video`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `play_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `cover_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `favorite_count` bigint NOT NULL,
  `comment_count` bigint NOT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `author_id` bigint NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_video_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 40 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of video
-- ----------------------------
INSERT INTO `video` VALUES (10, '2023-02-22 13:55:07.671', '2023-02-22 14:07:17.759', NULL, 'http://rp6phmpic.bkt.clouddn.com/7034037838899314688.mp4', 'http://rp6phmpic.bkt.clouddn.com/7034037838899314688.png', 3, 2, '美女', 3);
INSERT INTO `video` VALUES (11, '2023-02-22 13:55:39.157', '2023-02-22 14:07:19.923', NULL, 'http://rp6phmpic.bkt.clouddn.com/7034037971808419840.mp4', 'http://rp6phmpic.bkt.clouddn.com/7034037971808419840.png', 4, 0, '修狗', 3);
INSERT INTO `video` VALUES (12, '2023-02-22 14:02:09.183', '2023-02-22 14:07:58.408', NULL, 'http://rp6phmpic.bkt.clouddn.com/7034039608123850752.mp4', 'http://rp6phmpic.bkt.clouddn.com/7034039608123850752.png', 3, 0, 'bear', 5);
INSERT INTO `video` VALUES (13, '2023-02-22 14:08:30.574', '2023-02-23 13:13:50.844', NULL, 'http://rp6phmpic.bkt.clouddn.com/7034041205876850688.mp4', 'http://rp6phmpic.bkt.clouddn.com/7034041205876850688.png', 2, 0, '下雪了', 7);
INSERT INTO `video` VALUES (14, '2023-02-23 12:58:58.408', '2023-02-23 13:10:56.422', NULL, 'http://rp6phmpic.bkt.clouddn.com/7034386096267984896.mp4', 'http://rp6phmpic.bkt.clouddn.com/7034386096267984896.png', 1, 0, 'Bear', 9);
INSERT INTO `video` VALUES (15, '2023-02-23 13:10:56.330', '2023-02-23 13:35:46.410', NULL, 'http://rp6phmpic.bkt.clouddn.com/7034389107866337280.mp4', 'http://rp6phmpic.bkt.clouddn.com/7034389107866337280.png', 2, 4, 'Bear', 9);
INSERT INTO `video` VALUES (16, '2023-02-23 13:35:46.319', '2023-02-23 13:35:53.626', NULL, 'http://rp6phmpic.bkt.clouddn.com/7034395356892758016.mp4', 'http://rp6phmpic.bkt.clouddn.com/7034395356892758016.png', 1, 0, 'Bear', 9);
INSERT INTO `video` VALUES (17, '2023-02-23 13:35:53.505', '2023-02-23 13:37:41.862', NULL, 'http://rp6phmpic.bkt.clouddn.com/7034395387871887360.mp4', 'http://rp6phmpic.bkt.clouddn.com/7034395387871887360.png', 1, 0, 'Bear', 9);
INSERT INTO `video` VALUES (18, '2023-02-23 13:37:41.774', '2023-02-23 13:37:43.672', NULL, 'http://rp6phmpic.bkt.clouddn.com/7034395841427144704.mp4', 'http://rp6phmpic.bkt.clouddn.com/7034395841427144704.png', 1, 0, 'Bear', 9);
INSERT INTO `video` VALUES (19, '2023-02-23 13:37:43.576', '2023-02-23 13:37:45.286', NULL, 'http://rp6phmpic.bkt.clouddn.com/7034395849325019136.mp4', 'http://rp6phmpic.bkt.clouddn.com/7034395849325019136.png', 1, 0, 'Bear', 9);
INSERT INTO `video` VALUES (20, '2023-02-23 13:37:45.197', '2023-02-23 13:37:47.019', NULL, 'http://rp6phmpic.bkt.clouddn.com/7034395856316923904.mp4', 'http://rp6phmpic.bkt.clouddn.com/7034395856316923904.png', 1, 0, 'Bear', 9);
INSERT INTO `video` VALUES (21, '2023-02-23 13:37:46.933', '2023-02-23 13:37:48.751', NULL, 'http://rp6phmpic.bkt.clouddn.com/7034395863690510336.mp4', 'http://rp6phmpic.bkt.clouddn.com/7034395863690510336.png', 1, 0, 'Bear', 9);
INSERT INTO `video` VALUES (22, '2023-02-23 13:37:48.664', '2023-02-23 13:37:50.676', NULL, 'http://rp6phmpic.bkt.clouddn.com/7034395870820827136.mp4', 'http://rp6phmpic.bkt.clouddn.com/7034395870820827136.png', 1, 0, 'Bear', 9);
INSERT INTO `video` VALUES (23, '2023-02-23 13:37:50.585', '2023-02-23 13:37:52.452', NULL, 'http://rp6phmpic.bkt.clouddn.com/7034395878727090176.mp4', 'http://rp6phmpic.bkt.clouddn.com/7034395878727090176.png', 1, 0, 'Bear', 9);
INSERT INTO `video` VALUES (24, '2023-02-23 13:37:52.358', '2023-02-23 13:37:54.179', NULL, 'http://rp6phmpic.bkt.clouddn.com/7034395885962264576.mp4', 'http://rp6phmpic.bkt.clouddn.com/7034395885962264576.png', 1, 0, 'Bear', 9);
INSERT INTO `video` VALUES (25, '2023-02-23 13:37:54.085', '2023-02-23 13:37:55.957', NULL, 'http://rp6phmpic.bkt.clouddn.com/7034395893197438976.mp4', 'http://rp6phmpic.bkt.clouddn.com/7034395893197438976.png', 1, 0, 'Bear', 9);
INSERT INTO `video` VALUES (26, '2023-02-23 13:37:55.854', '2023-02-23 13:37:57.660', NULL, 'http://rp6phmpic.bkt.clouddn.com/7034395901011427328.mp4', 'http://rp6phmpic.bkt.clouddn.com/7034395901011427328.png', 1, 0, 'Bear', 9);
INSERT INTO `video` VALUES (27, '2023-02-23 13:37:57.562', '2023-02-23 13:37:59.390', NULL, 'http://rp6phmpic.bkt.clouddn.com/7034395908254990336.mp4', 'http://rp6phmpic.bkt.clouddn.com/7034395908254990336.png', 1, 0, 'Bear', 9);
INSERT INTO `video` VALUES (28, '2023-02-23 13:37:59.279', '2023-02-23 13:39:07.991', NULL, 'http://rp6phmpic.bkt.clouddn.com/7034395915427250176.mp4', 'http://rp6phmpic.bkt.clouddn.com/7034395915427250176.png', 1, 0, 'Bear', 9);
INSERT INTO `video` VALUES (29, '2023-02-23 13:39:07.935', '2023-02-23 13:39:09.997', NULL, 'http://rp6phmpic.bkt.clouddn.com/7034396202779017216.mp4', 'http://rp6phmpic.bkt.clouddn.com/7034396202779017216.png', 1, 0, 'Bear', 9);
INSERT INTO `video` VALUES (30, '2023-02-23 13:39:09.879', '2023-02-23 13:39:11.961', NULL, 'http://rp6phmpic.bkt.clouddn.com/7034396211540918272.mp4', 'http://rp6phmpic.bkt.clouddn.com/7034396211540918272.png', 1, 0, 'Bear', 9);
INSERT INTO `video` VALUES (31, '2023-02-23 13:39:11.826', '2023-02-23 13:39:13.812', NULL, 'http://rp6phmpic.bkt.clouddn.com/7034396219359100928.mp4', 'http://rp6phmpic.bkt.clouddn.com/7034396219359100928.png', 1, 0, 'Bear', 9);
INSERT INTO `video` VALUES (32, '2023-02-23 13:39:13.683', '2023-02-23 13:39:15.749', NULL, 'http://rp6phmpic.bkt.clouddn.com/7034396227370221568.mp4', 'http://rp6phmpic.bkt.clouddn.com/7034396227370221568.png', 1, 0, 'Bear', 9);
INSERT INTO `video` VALUES (33, '2023-02-23 13:39:15.630', '2023-02-23 13:39:17.555', NULL, 'http://rp6phmpic.bkt.clouddn.com/7034396235213570048.mp4', 'http://rp6phmpic.bkt.clouddn.com/7034396235213570048.png', 1, 0, 'Bear', 9);
INSERT INTO `video` VALUES (34, '2023-02-23 13:39:17.431', '2023-02-23 13:39:19.402', NULL, 'http://rp6phmpic.bkt.clouddn.com/7034396243048529920.mp4', 'http://rp6phmpic.bkt.clouddn.com/7034396243048529920.png', 1, 0, 'Bear', 9);
INSERT INTO `video` VALUES (35, '2023-02-23 13:39:19.278', '2023-02-23 13:39:21.349', NULL, 'http://rp6phmpic.bkt.clouddn.com/7034396250925432832.mp4', 'http://rp6phmpic.bkt.clouddn.com/7034396250925432832.png', 1, 0, 'Bear', 9);
INSERT INTO `video` VALUES (36, '2023-02-23 13:39:21.215', '2023-02-23 13:39:23.155', NULL, 'http://rp6phmpic.bkt.clouddn.com/7034396258760392704.mp4', 'http://rp6phmpic.bkt.clouddn.com/7034396258760392704.png', 1, 0, 'Bear', 9);
INSERT INTO `video` VALUES (37, '2023-02-23 13:39:23.017', '2023-02-23 13:39:24.972', NULL, 'http://rp6phmpic.bkt.clouddn.com/7034396266524049408.mp4', 'http://rp6phmpic.bkt.clouddn.com/7034396266524049408.png', 1, 0, 'Bear', 9);
INSERT INTO `video` VALUES (38, '2023-02-23 13:39:24.837', '2023-02-23 13:39:26.830', NULL, 'http://rp6phmpic.bkt.clouddn.com/7034396274325454848.mp4', 'http://rp6phmpic.bkt.clouddn.com/7034396274325454848.png', 1, 0, 'Bear', 9);
INSERT INTO `video` VALUES (39, '2023-02-23 13:39:26.704', '2023-02-23 13:39:26.704', NULL, 'http://rp6phmpic.bkt.clouddn.com/7034396282026196992.mp4', 'http://rp6phmpic.bkt.clouddn.com/7034396282026196992.png', 0, 0, 'Bear', 9);

SET FOREIGN_KEY_CHECKS = 1;
