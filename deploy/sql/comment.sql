CREATE TABLE `comment`
(
    `id`              bigint        NOT NULL AUTO_INCREMENT COMMENT 'id',
    `create_time`     datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`     datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time`     datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `content`         varchar(1000) NOT NULL COMMENT '评论内容',
    `user_id`         bigint        NOT NULL COMMENT '用户id',
    `vedio_id`        bigint        NOT NULL COMMENT '视频id',
    `like`            int           NOT NULL DEFAULT '0' COMMENT '点赞',
    `root_comment_id` bigint        NOT NULL DEFAULT '0' COMMENT '顶层评论id，是0表示为顶层评论，否则为顶层评论id',
    `to_comment_id`   bigint        NOT NULL DEFAULT '0' COMMENT '被回复的评论id，是0为顶层评论，否则为目标评论id',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='评论';