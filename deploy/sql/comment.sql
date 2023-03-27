-- giligili.comment definition

CREATE TABLE `comment`
(
    `id`          int          NOT NULL AUTO_INCREMENT COMMENT '评论ID',
    `create_time` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time` datetime              DEFAULT NULL COMMENT '删除时间',
    `user_id`     int          NOT NULL COMMENT '评论用户ID',
    `video_id`    int          NOT NULL COMMENT '评论视频ID',
    `reply_id`    int                   DEFAULT NULL COMMENT '回复评论ID',
    `content`     varchar(255) NOT NULL COMMENT '评论内容',
    `like`        int          NOT NULL DEFAULT '0' COMMENT '点赞数',
    `dislike`     int          NOT NULL DEFAULT '0' COMMENT '踩数',
    PRIMARY KEY (`id`),
    KEY           `idx_user` (`user_id`),
    KEY           `idx_video` (`video_id`),
    KEY           `idx_reply` (`reply_id`),
    CONSTRAINT `fk_reply_comment` FOREIGN KEY (`reply_id`) REFERENCES `comment` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_user_comment` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_video_comment` FOREIGN KEY (`video_id`) REFERENCES `video` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='评论表';