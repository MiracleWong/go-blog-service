

CREATE TABLE blog_tag
(
id int(10) unsigned NOT NULL AUTO_INCREMENT,
name varchar(100) DEFAULT '' COMMENT '标签名称',
created_on int(10) unsigned DEFAULT '0' COMMENT '创建时间',
created_by varchar(100) DEFAULT '' COMMENT '创建人',
modified_on int(10) unsigned DEFAULT '0' COMMENT '修改时间',
modified_by varchar(100) DEFAULT '' COMMENT '修改人',
deleted_on int(10) unsigned DEFAULT '0' COMMENT '删除时间',
is_del tinyint(3) unsigned DEFAULT '1' COMMENT '是否删除，0为未删除，1为删除',
state tinyint(3) unsigned DEFAULT '1' COMMENT '状态0为禁用，状态1为启用',
PRIMARY KEY (id)
) ENGINE = InnoDB
DEFAULT CHARSET = utf8mb4 COMMENT ='标签管理';


CREATE TABLE blog_article
(
id int(10) unsigned NOT NULL AUTO_INCREMENT,
title varchar(100) DEFAULT '' COMMENT '文章标题',
`desc` varchar(255) DEFAULT '' COMMENT '文章简述',
cover_image_url varchar(255) DEFAULT '' COMMENT '封面图片地址',
content longtext COMMENT '文章内容',
created_on int(10) unsigned DEFAULT '0' COMMENT '创建时间',
created_by varchar(100) DEFAULT '' COMMENT '创建人',
modified_on int(10) unsigned DEFAULT '0' COMMENT '修改时间',
modified_by varchar(100) DEFAULT '' COMMENT '修改人',
deleted_on int(10) unsigned DEFAULT '0' COMMENT '删除时间',
is_del tinyint(3) unsigned DEFAULT '1' COMMENT '是否删除，0为未删除，1为删除',
state tinyint(3) unsigned DEFAULT '1' COMMENT '状态0为禁用，状态1为启用',
PRIMARY KEY (id)
) ENGINE = InnoDB
DEFAULT CHARSET = utf8mb4 COMMENT ='文章管理';


CREATE TABLE blog_auth
(
id int(10) unsigned NOT NULL AUTO_INCREMENT,
app_key varchar(20) DEFAULT '' COMMENT 'Key',
app_secret varchar(50) DEFAULT '' COMMENT 'Secret',
created_on int(10) unsigned DEFAULT '0' COMMENT '创建时间',
created_by varchar(100) DEFAULT '' COMMENT '创建人',
modified_on int(10) unsigned DEFAULT '0' COMMENT '修改时间',
modified_by varchar(100) DEFAULT '' COMMENT '修改人',
deleted_on int(10) unsigned DEFAULT '0' COMMENT '删除时间',
is_del tinyint(3) unsigned DEFAULT '1' COMMENT '是否删除，0为未删除，1为删除',
PRIMARY KEY (id) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COMMENT = '认证管理';


INSERT INTO blog_service.blog_auth (id, app_key, app_secret, created_on, created_by, modified_on,
modified_by, deleted_on,
is_del) VALUE (1, 'eddycjy', 'go-programming-tour-book', 0, 'eddycjy', 0, '', 0, 0);

INSERT INTO blog_service.blog_auth (id, app_key, app_secret, created_on, created_by, modified_on,
modified_by, deleted_on,
is_del) VALUE (1, 'MiracleWong', 'go-programming-tour-book', 0, 'MiracleWong', 0, '', 0, 0);