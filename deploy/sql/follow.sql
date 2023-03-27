-- giligili.follow definition

CREATE TABLE `follow`
(
    `id`          int      NOT NULL AUTO_INCREMENT COMMENT '关注ID',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `delete_time` datetime          DEFAULT NULL COMMENT '删除时间',
    `follower_id` int      NOT NULL COMMENT '关注者ID',
    `followee_id` int      NOT NULL COMMENT '被关注者ID',
    PRIMARY KEY (`id`),
    KEY           `idx_follower` (`follower_id`),
    KEY           `idx_followee` (`followee_id`),
    CONSTRAINT `fk_followee` FOREIGN KEY (`followee_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_follower` FOREIGN KEY (`follower_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='关注表';