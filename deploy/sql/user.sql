CREATE TABLE `user`
(
    `id`          BIGINT      NOT NULL AUTO_INCREMENT COMMENT 'id',
    `create_time` DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time` DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `name`        VARCHAR(32) NOT NULL COMMENT '用户名',
    `password`    VARCHAR(32) NOT NULL COMMENT '密码',
    `avatar`      MEDIUMTEXT  NOT NULL COMMENT '头像',
    PRIMARY KEY (`id`)
) COMMENT '用户';