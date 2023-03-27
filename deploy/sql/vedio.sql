-- giligili.video definition

CREATE TABLE `video`
(
    `id`          int          NOT NULL AUTO_INCREMENT COMMENT '视频ID',
    `create_time` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time` datetime              DEFAULT NULL COMMENT '删除时间',
    `title`       varchar(255) NOT NULL COMMENT '视频标题',
    `url`         varchar(255) NOT NULL COMMENT '视频URL',
    `user_id`     int          NOT NULL COMMENT '上传用户ID',
    `description` varchar(255)          DEFAULT NULL COMMENT '视频描述',
    `like`        int          NOT NULL DEFAULT '0' COMMENT '点赞数',
    `dislike`     int          NOT NULL DEFAULT '0' COMMENT '踩数',
    `status`      tinyint(1) NOT NULL DEFAULT '0' COMMENT '审核状态，0表示未审核，1表示已审核',
    PRIMARY KEY (`id`),
    KEY           `idx_user` (`user_id`),
    CONSTRAINT `fk_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='视频表';