-- DROP TABLE IF EXISTS user;


-- CREATE TABLE user (
--   id BIGINT UNSIGNED NOT NULL AUTOINCREMENT, -- 識別子
--   name VARCHAR( 80 ) NOT NULL, -- ユーザー名
--   password VARCHAR( 80 ) NOT NULL, -- パスワードハッシュ
--   role VARCHAR( 80 ) NOT NULL, -- ロール
--   created DATETIME( 6 ) NOT NULL, -- 作成日付時刻
--   modified DATETIME( 6 ) NOT NULL, -- 修正日付時刻
--   -- PRIMARY KEY
--   CONSTRAINT pk_user PRIMARY KEY (
--     serial
--   )
-- );

CREATE TABLE `user`
(
    `id`       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ユーザーの識別子',
    `name`     varchar(20) NOT NULL COMMENT 'ユーザー名',
    `password` VARCHAR(80) NOT NULL COMMENT 'パスワードハッシュ',
    `role`     VARCHAR(80) NOT NULL COMMENT 'ロール',
    `created`  DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `modified` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uix_name` (`name`) USING BTREE
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザー';

CREATE TABLE `task`
(
    `id`       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'タスクの識別子',
    `title`    VARCHAR(128) NOT NULL COMMENT 'タスクのタイトル',
    `status`   VARCHAR(20)  NOT NULL COMMENT 'タスクの状態',
    `created`  DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `modified` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='タスク';