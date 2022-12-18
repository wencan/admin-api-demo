package adminmodel

import "time"

/*
	RolePermission 角色的权限。

CREATE TABLE `admin_role_permission` (

	`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
	`role_id` bigint NOT NULL DEFAULT '0' COMMENT 'admin角色id',
	`permission_id` bigint NOT NULL DEFAULT '0' COMMENT 'admin权限id',
	`deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '逻辑删除标识。0为未删除；1为已删除。',
	`create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY (`id`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='admin角色的权限';
*/
type RolePermission struct {
	// ID id。
	ID int64 `json:"id" db:"id"`

	// RoleID 角色id。
	RoleID int64 `json:"role_id" db:"role_id"`

	// PermissionID 权限id。
	PermissionID int64 `json:"permission_id" db:"permission_id"`

	// Deleted 是否已经逻辑删除。
	Deleted bool `json:"-" db:"deleted"`

	// CreateTime 创建时间。
	CreateTime time.Time `json:"create_time" db:"create_time"`

	// UpdateTime 更新时间。
	UpdateTime time.Time `json:"update_time" db:"update_time"`
}
