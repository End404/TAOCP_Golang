CREATE TABLE 'tbl_file'(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `file_sha1` char(40) NOT NULL DEFAULT '' COMMENT'文件hash',
    `file_name` varchar(256) NOT NULL DEFAULT '' COMMENT '文件名',
    `file_size` bigint(20) DEFAULT '0' COMMENT '文件大小',
    `file_addr` varchar(1024) NOT NULL DEFAULT '' COMMENT '文件存储位置',
   
    `create_at` datetiem default NOW() COMMENT '创建日期',
    `update_at` datetiem default NOW() on update current_timestamp() COMMENT '更新',
    `status` int(11) NOT NULL DEFAULT '0' COMMENT '状态（可用、禁用、已删除等状态）',
    
    `ext1` int(11) DEFAULT '0' COMMENT '备用字段1',
    `ext2` text COMMENT '备用字段2',
    PRIMARY KEY(`id`),
    UNIQUE KEY `idx_file_hash`(`file_sha1`),
    KEY `idx_status` (`status`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE 'tbl_user'(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `user_name` varchar(64) NOT NULL DEFAULT '' COMMENT '用户名',
    `user_pwd` varchar(256) NOT NULL DEFAULT '' COMMENT '用户密码',
    
    `email` varchar(64) DEFAULT '' COMMENT '邮箱',
    `phone` varchar(128) DEFAULT '' COMMENT '手机号',
    `email_validated` tinyint(1) DEFAULT 0 COMMENT '邮箱号是否已验证',
    `phone_validated` tinyint(1) DEFAULT 0 COMMENT '手机号是否已验证',

    `signup_at` datetiem DEFAULT CURRENT_TIMESTAMP COMMENT '注册日期',
    `last_active` datetiem DEFAULT CURRENT_TIMESTAMP ON UPDATE_TIMESTAMP '最后一次登录或使用的时间',
    
    `profile` text COMMENT '用户属性',
    `status` int(11) NOT NULL DEFAULT '0' COMMENT '账户状态（启用、禁用、锁定、标记已删除等状态）',

    PRIMARY KEY(`id`),
    UNIQUE KEY `idx_phone`(`phone`),
    KEY `idx_status` (`status`)
)ENGINE=InnoDB  AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;


CREATE TABLE 'tbl_user_file'(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `user_name` varchar(64) NOT NULL ,

    `file_sha1` varchar(64) NOT NULL DEFAULT '' COMMENT '文件hash',
    `file_size` bigint(20) DEFAULT '0' COMMENT '文件大小',
    `file_name` varchar(256) NOT NULL DEFAULT '' COMMENT '文件名',

    `upload_at` datetiem DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间',
    `last_active` datetiem DEFAULT CURRENT_TIMESTAMP ON UPDATE_TIMESTAMP '最后一次修改时间',
    
    `status` int(11) NOT NULL DEFAULT '0' COMMENT '文件状态（正常、已删除、禁用）',

    UNIQUE KEY `idx_user_file`(`user_name`,`file_sha1`),
    KEY `idx_status` (`status`)
    KEY `idx_user_id` (`user_name`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
