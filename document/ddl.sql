---- user struct ----

CREATE TABLE `account`
(
  `id`          BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `account_id`  VARCHAR(128)    NOT NULL COMMENT 'uniq id'
  `username` VARCHAR(64)     NOT NULL COMMENT '用户名',
  `password`   VARCHAR(256)    NOT NULL COMMENT '密码md5',
  `salt`       VARCHAR(256)    NOT NULL COMMENT '盐',
  `status`     VARCHAR(32)     NOT NULL COMMENT '帐号状态',
  `created_at`  DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at`  DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` BIGINT UNSIGNED DEFAULT 0 COMMENT '软删除标记',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `uniq_user_id` (`account_id`),
  UNIQUE INDEX `uniq_username` (`username`, `deleted_at`)
)ENGINE = INNODB
  DEFAULT CHARSET = utf8mb4,
    COMMENT '用户账号表';

CREATE TABLE `user`
(
    `id`          BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `account_id`  VARCHAR(128)    NOT NULL COMMENT '',
    `user_id`     VARCHAR(128)    NOT NULL COMMENT '用户唯一ID',
    `name`    VARCHAR(64)     NOT NULL COMMENT '用户名称',
    `gender`      VARCHAR(16)     NOT NULL COMMENT '性别',
    `phone`       VARCHAR(32)     NOT NULL COMMENT '用户手机号码',
    `email`       VARCHAR(64)     NOT NULL COMMENT '用户邮箱',
    `description` VARCHAR(256)    NULL COMMENT '用户描述',
    `created_at`  DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` BIGINT UNSIGNED DEFAULT 0 COMMENT '软删除标记',
    PRIMARY KEY (`id`),
    UNIQUE INDEX `uniq_user_id` (`user_id`),
    UNIQUE INDEX `uniq_account_id` (`account_id`, `deleted_at`)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8mb4,
    COMMENT '用户信息表';

-- CREATE TABLE `login_record`
-- (
--   `id`        BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
--   `user_id`     VARCHAR(128)    NOT NULL COMMENT '用户唯一ID',
--   `login_at`  DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '',
--   `ip`      VARCHAR(128)     NOT NULL COMMENT '登陆IP',
--   `device`    VARCHAR(128)     NOT NULL COMMENT '登陆设备',
--   `reason`    VARCHAR(64)     NOT NULL COMMENT '登陆失败的原因',
--   `status`    VARCHAR(32)     NOT NULL COMMENT '登陆状态success/fail',
--   `description` VARCHAR(256)    NULL COMMENT '用户详情',
--   `created_at`  DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
--   `updated_at`  DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
--   `deleted_at` DATETIME NULL COMMENT '软删除标记',
--   PRIMARY KEY (`id`),
--   INDEX `idx_username_login_at` (`user_id`, `login_at`)
-- ) ENGINE = INNODB
--   DEFAULT CHARSET = utf8mb4
--     COMMENT '登陆记录表';

---- domain struct ----

CREATE TABLE `domain` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `domain_id` VARCHAR(128) NOT NULL COMMENT '作用域唯一ID',
  `name` VARCHAR(64)     NOT NULL COMMENT '名称',
  `description` VARCHAR(256)    NULL COMMENT '描述',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` BIGINT UNSIGNED DEFAULT 0 COMMENT '软删除标记',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `uniq_domain_id` (`domain_id`),
  UNIQUE INDEX `uniq_domain_name` (`name`, `deleted_at`)
)ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT '作用域表';

---- rbac struct ----

CREATE TABLE `user_role` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `relation_id` VARCHAR(128) NOT NULL COMMENT '用户-角色关系id',
  `user_id` VARCHAR(128) NOT NULL COMMENT '用户唯一ID',
  `role_id` VARCHAR(128) NOT NULL COMMENT '角色唯一ID',
  `status` VARCHAR(32) NOT NULL COMMENT '映射状态',
  `expire_at` DATETIME NOT NULL COMMENT '过期时间',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
 `deleted_at` BIGINT UNSIGNED DEFAULT 0 COMMENT '软删除标记',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `uniq_relation_id` (`relation_id`),
  UNIQUE INDEX `uniq_user_role` (`user_id`, `role_id`, `deleted_at`),
  INDEX `idx_role_id` (`role_id`, `deleted_at`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT '用户-角色映射表';

CREATE TABLE `role` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `role_id` VARCHAR(128) NOT NULL COMMENT '角色唯一ID',
  `domain_id` VARCHAR(64) NOT NULL COMMENT '作用域ID',
  `parent_role_id` VARCHAR(128) NOT NULL COMMENT '父角色ID',
  `name` VARCHAR(64) NOT NULL COMMENT '角色名称',
  `status` VARCHAR(32) NOT NULL COMMENT '角色状态',
  `description` VARCHAR(256) NULL COMMENT '角色描述',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` BIGINT UNSIGNED DEFAULT 0 COMMENT '软删除标记',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `uniq_role_id` (`role_id`),
  UNIQUE INDEX `uniq_role_name` (`name`, `deleted_at`),
  INDEX `idx_parent_role_id` (`parent_role_id`),
  INDEX `idx_domain` (`domain_id`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT '角色表';

CREATE TABLE `resource` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `resource_id` VARCHAR(128) NOT NULL COMMENT '资源唯一ID',
    `domain_id` VARCHAR(64) NOT NULL COMMENT '作用域ID',
    `name` VARCHAR(64)     NOT NULL COMMENT '资源名称',
    `status` VARCHAR(32) NOT NULL COMMENT '资源状态',
    `description` VARCHAR(256) NULL COMMENT '资源描述',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` BIGINT UNSIGNED DEFAULT 0 COMMENT '软删除标记',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `uniq_resource_id` (`resource_id`),
  UNIQUE INDEX `uniq_resource_name` (`name`, `deleted_at`),
  INDEX `idx_domain` (`domain_id`)
)ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT '';

CREATE TABLE `permission` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `permission_id` VARCHAR(128) NOT NULL COMMENT '权限ID',
  `role_id` VARCHAR(128) NOT NULL COMMENT '角色唯一ID',
  `resource_id` VARCHAR(128) NOT NULL COMMENT '资源ID',
  `action` VARCHAR(32) NOT NULL COMMENT '权限操作',
  `effect` VARCHAR(32) NOT NULL COMMENT 'allow/deny',
  `status` VARCHAR(32) NOT NULL COMMENT '状态',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` BIGINT UNSIGNED DEFAULT 0 COMMENT '软删除标记',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `uniq_permission_id` (`permission_id`),
  UNIQUE INDEX `uniq_role_resource_action` (`role_id`, `resource_id`,`action`, `deleted_at`),
  INDEX `idx_resource` (`resource_id`, `deleted_at`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT '权限表';

-- CREATE TABLE `operation_record` (
--   `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
--   `record_id` VARCHAR(128) NOT NULL COMMENT '',
--   `data_type` VARCHAR(32) NOT NULL COMMENT '',
--   `data_id` VARCHAR(128) NOT NULL COMMENT '',
--   `operation` VARCHAR(32) NOT NULL COMMENT '',
--   `operator_id` VARCHAR(128) NOT NULL COMMENT '',
--   `trace_id` VARCHAR(128) NOT NULL COMMENT '',
--   `previous_value` TEXT NULL COMMENT '',
--   `current_value` TEXT NULL COMMENT '',
--   `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
--   PRIMARY KEY (`id`),
--   UNIQUE INDEX `uniq_record_id` (`record_id`),
--   INDEX `idx_data_id_created_at` (`data_id`, `created_at`),
--   INDEX `idx_operator_id_created_at` (`operator_id`, `created_at`)
-- ) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT '权限表';
