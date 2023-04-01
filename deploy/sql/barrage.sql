CREATE TABLE `barrage`
(
    `id`            int          NOT NULL AUTO_INCREMENT COMMENT '弹幕ID',
    `create_time`   datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `delete_time`   datetime              DEFAULT NULL COMMENT '删除时间',
    `user_id`       int          NOT NULL COMMENT '用户ID',
    `user_nickname` varchar(255) NOT NULL COMMENT '用户昵称',
    `video_id`      int          NOT NULL COMMENT '视频ID',
    `text`          varchar(255) NOT NULL COMMENT '弹幕内容',
    `color`         int          NOT NULL COMMENT '弹幕颜色',
    `type`          tinyint(1)   NOT NULL DEFAULT '0' COMMENT '弹幕类型，0表示滚动弹幕，1表示顶部弹幕，2表示底部弹幕',
    `timestamp`     int          NOT NULL DEFAULT '0' COMMENT '弹幕出现时间（秒）',
    PRIMARY KEY (`id`),
    KEY `idx_user` (`user_id`),
    KEY `idx_video` (`video_id`),
    CONSTRAINT `fk_user_barrage` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_video_barrage` FOREIGN KEY (`video_id`) REFERENCES `video` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='弹幕表';