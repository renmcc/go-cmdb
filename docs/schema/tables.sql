CREATE TABLE IF NOT EXISTS `host` (
   `xid` char(64) not null primary key  comment '唯一id',
   `status` smallint not null default 0 comment "0删除 1新增 2 更新",
   `create_at` bigint not null comment '创建时间',
   `create_by` char(10) not null default '' comment '创建操作人',
   `update_at` bigint not null default 0 comment '更新时间',
   `update_by` char(10) not null default '' comment '更新操作人',
   `delete_at` bigint not null default 0 comment '删除时间',
   `delete_by` char(10) not null default '' comment '删除操作人',
   `resource_id` char(32) not null default '' comment '资源id',
   `vendor` char(16) not null default '' comment '厂商',
   `name` char(10) not null default '' comment '资源名',
   `region` char(10) not null default '' comment '地域',
   `expire_at` bigint not null default 0 comment '过期时间',
   `public_ip` char(32) not null default '' comment '公网IP',
   `private_ip` char(32) not null default '' comment '内网ip',
   `cpu` char(10) not null default '' comment 'cpu',
   `memory` char(10) not null default '' comment '内存',
   `os_type` char(10) not null default '' comment '操作系统类型',
   `os_name` char(10) not null default '' comment '系统名称',
   `serial_number` char(50) not null default '' comment '序列号',
  KEY `idx_xid` (`xid`) USING BTREE COMMENT '唯一id搜索'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

