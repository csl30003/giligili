-- giligili.`user` definition

CREATE TABLE `user`
(
    `id`          int          NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `create_time` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time` datetime              DEFAULT NULL COMMENT '删除时间',
    `username`    varchar(255) NOT NULL COMMENT '用户名',
    `password`    varchar(255) NOT NULL COMMENT '密码',
    `nickname`    varchar(255) NOT NULL COMMENT '昵称',
    `email`       varchar(255)          DEFAULT NULL COMMENT '邮箱',
    `phone`       varchar(255)          DEFAULT NULL COMMENT '电话',
    `avatar`      varchar(255)          DEFAULT NULL COMMENT '头像',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表';