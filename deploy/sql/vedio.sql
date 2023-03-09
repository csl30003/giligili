CREATE TABLE `vedio`
(
    `id`          BIGINT       NOT NULL AUTO_INCREMENT COMMENT 'id',
    `create_time` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `title`       VARCHAR(32)  NOT NULL COMMENT '标题',
    `info`        VARCHAR(100) NOT NULL COMMENT '简介',
    `url`         MEDIUMTEXT   NOT NULL COMMENT '路由',
    `avatar`      MEDIUMTEXT   NOT NULL COMMENT '封面',
    `status`      tinyint      NOT NULL DEFAULT '0' COMMENT '状态，0表示未审核，1表示已审核',
    `user_id`     BIGINT       NOT NULL COMMENT '用户id',
    PRIMARY KEY (`id`)
) COMMENT '视频';