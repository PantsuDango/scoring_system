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