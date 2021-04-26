CREATE DATABASE `scoring_system`;

USE `scoring_system`;

CREATE TABLE `user` (
   `id` int(11) NOT NULL AUTO_INCREMENT,
   `nick` varchar(32) NOT NULL COMMENT '用户昵称',
   `username` varchar(32) NOT NULL COMMENT '登录账号',
   `password` varchar(32) NOT NULL COMMENT '登录密码, 存md5',
   `salt` varchar(32) NOT NULL COMMENT '盐',
   `type` tinyint(4) NOT NULL COMMENT '用户类型: 1-管理员, 2-评委, 3-选手',
   `status` tinyint(4) NOT NULL DEFAULT 0 COMMENT  '账号状态: 0-存续, 1-废除',
   `createtime` datetime NOT NULL COMMENT '创建时间',
   `lastupdate` datetime NOT NULL COMMENT '更新时间',
   PRIMARY KEY (`id`),
   UNIQUE KEY (`username`),
   UNIQUE KEY (`nick`, `type`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户登录信息表';
INSERT INTO `user` VALUES (1, '管理员', 'admin', 'e160fa11757ed2883e890e483f1c6208', 'h6du1cxo', 1, 0, '2021-04-18 22:58:58', '2021-04-18 22:59:01');

CREATE TABLE `project` (
   `id` int(11) NOT NULL AUTO_INCREMENT,
   `name` varchar(32) NOT NULL COMMENT '项目名称',
   `content` varchar(32) DEFAULT NULL COMMENT '项目说明',
   `createtime` datetime NOT NULL COMMENT '创建时间',
   `lastupdate` datetime NOT NULL COMMENT '更新时间',
   PRIMARY KEY (`id`),
   UNIQUE KEY (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='项目信息表';

CREATE TABLE `project_user_map` (
   `id` int(11) NOT NULL AUTO_INCREMENT,
   `project_id` varchar(32) NOT NULL COMMENT '项目id',
   `user_id` varchar(32) NOT NULL COMMENT '项目绑定的用户id',
   `type` tinyint(4) NOT NULL COMMENT '用户类型: 1-管理员, 2-评委, 3-选手',
   `createtime` datetime NOT NULL COMMENT '创建时间',
   `lastupdate` datetime NOT NULL COMMENT '更新时间',
   PRIMARY KEY (`id`),
   UNIQUE KEY (`project_id`, `user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='项目和用户关系表';

CREATE TABLE `score` (
   `id` int(11) NOT NULL AUTO_INCREMENT,
   `project_id` varchar(32) NOT NULL COMMENT '项目id',
   `player_id` varchar(32) NOT NULL COMMENT '选手id',
   `score` varchar(32) NOT NULL COMMENT '分数',
   `judges_id` tinyint(4) NOT NULL COMMENT '评委id',
   `createtime` datetime NOT NULL COMMENT '创建时间',
   `lastupdate` datetime NOT NULL COMMENT '更新时间',
   PRIMARY KEY (`id`),
   UNIQUE KEY (`project_id`, `player_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='评分表';
alter table `score` drop index `project_id`;
alter table `score` add unique key (`project_id`, `player_id`, `judges_id`);