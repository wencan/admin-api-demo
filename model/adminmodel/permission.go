package adminmodel

import "time"

/*
	Permission admin权限。

CREATE TABLE `admin_permission` (

	`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
	`title` varchar(60) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '名称',
	`permission_code` varchar(30) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '权限',
	`deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '逻辑删除标识。0为未删除；1为已删除。',
	`create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY (`id`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='admin权限';
*/
type Permission struct {
	// ID id。
	ID int64 `json:"id" db:"id"`

	// Name 名称。
	Name string `json:"title" db:"title"`

	// PermissionCode 权限代码。
	PermissionCode string `json:"permission_code" db:"permission_code"`

	// Deleted 是否已经逻辑删除。
	Deleted bool `json:"-" db:"deleted"`

	// CreateTime 创建时间。
	CreateTime time.Time `json:"create_time" db:"create_time"`

	// UpdateTime 更新时间。
	UpdateTime time.Time `json:"update_time" db:"update_time"`
}
