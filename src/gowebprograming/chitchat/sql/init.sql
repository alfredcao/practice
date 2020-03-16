CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `uuid` varchar(64) NOT NULL COMMENT 'uuid',
  `name` varchar(20) NOT NULL COMMENT '名称',
  `email` varchar(200) NOT NULL COMMENT '电子邮箱',
  `password` varchar(200)  NULL COMMENT '登录密码',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_uuid` (`uuid`),
  UNIQUE KEY `uk_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';

CREATE TABLE `session` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `uuid` varchar(64) NOT NULL COMMENT 'uuid',
  `email` varchar(200) NOT NULL COMMENT '电子邮箱',
  `user_id` bigint(20) NOT NULL COMMENT '用户ID，关联user表id字段',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_uuid` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='会话表';

CREATE TABLE `thread` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `uuid` varchar(64) NOT NULL COMMENT 'uuid',
  `title` varchar(400) NOT NULL COMMENT '标题',
  `content` text NOT NULL COMMENT '内容',
  `user_id` bigint(20) NOT NULL COMMENT '发帖用户ID，关联user表id字段',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_uuid` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='帖子表';

CREATE TABLE `post` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `uuid` varchar(64) NOT NULL COMMENT 'uuid',
  `content` text NOT NULL COMMENT '内容',
  `user_id` bigint(20) NOT NULL COMMENT '回帖用户ID，关联user表id字段',
  `thread_id` bigint(20) NOT NULL COMMENT '帖子ID，关联thread表id字段',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_uuid` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='回帖表';