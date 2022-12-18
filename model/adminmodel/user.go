package adminmodel

import "time"

/*
	User admin用户。

CREATE TABLE `admin_user` (

	`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
	`user_name` varchar(30) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
	`nice_name` varchar(60) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
	`password_sha256` varchar(128) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码的sha256摘要',
	`secure_salt` varchar(64) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '计算摘要用的盐',
	`deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '逻辑删除标识。0为未删除；1为已删除。',
	`create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY (`id`),
	UNIQUE KEY `idx_admin_user_username` (`user_name`) USING BTREE

) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='admin用户';
*/
type User struct {
	// ID id。
	ID int64 `json:"id" db:"id"`

	// UserName 用户名。
	UserName string `json:"user_name" db:"user_name"`

	// NiceName 昵称
	NiceName string `json:"nice_name" db:"nice_name"`

	// PasswordSha256 密码sha256摘要。sha256({password}{secure_salt})
	PasswordSha256 string `json:"-" db:"password_sha256"`

	// SecureSalt 计算摘要用的盐。
	SecureSalt string `json:"-" db:"secure_salt"`

	// Deleted 是否已经逻辑删除。
	Deleted bool `json:"-" db:"deleted"`

	// CreateTime 创建时间。
	CreateTime time.Time `json:"create_time" db:"create_time"`

	// UpdateTime 更新时间。
	UpdateTime time.Time `json:"update_time" db:"update_time"`
}
