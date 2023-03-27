-- giligili.chat definition

CREATE TABLE `chat`
(
    `id`           int          NOT NULL AUTO_INCREMENT COMMENT '聊天ID',
    `create_time`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `delete_time`  datetime              DEFAULT NULL COMMENT '删除时间',
    `from_user_id` int          NOT NULL COMMENT '发送者ID',
    `to_user_id`   int          NOT NULL COMMENT '接收者ID',
    `content`      varchar(255) NOT NULL COMMENT '聊天内容',
    PRIMARY KEY (`id`),
    KEY            `idx_from_user` (`from_user_id`),
    KEY            `idx_to_user` (`to_user_id`),
    CONSTRAINT `fk_from_user` FOREIGN KEY (`from_user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_to_user` FOREIGN KEY (`to_user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='聊天表';