package schema

import "database/sql"

type GetUsersGroup struct {
	ID               sql.NullInt64  `gorm:"column:id"`
	GroupCode        sql.NullString `gorm:"column:group_code"`
	GroupName        sql.NullString `gorm:"column:group_name"`
	GroupDescription sql.NullString `gorm:"column:group_description"`
	CreatedAt        sql.NullString `gorm:"column:created_at"`
	DeletedAt        sql.NullString `gorm:"column:deleted_at"`
	UpdatedAt        sql.NullString `gorm:"column:updated_at"`
}

type CreateUsersGroup struct {
	ID               *int    `gorm:"column:id"`
	GroupCode        *string `gorm:"column:group_code"`
	GroupName        *string `gorm:"column:group_name"`
	GroupDescription *string `gorm:"column:group_description"`
	CreatedAt        *string `gorm:"column:created_at"`
	DeletedAt        *string `gorm:"column:deleted_at"`
	UpdatedAt        *string `gorm:"column:updated_at"`
}

type UpdateUsersGroup struct {
	ID               *int    `gorm:"column:id"`
	GroupCode        *string `gorm:"column:group_code"`
	GroupName        *string `gorm:"column:group_name"`
	GroupDescription *string `gorm:"column:group_description"`
	CreatedAt        *string `gorm:"column:created_at"`
	DeletedAt        *string `gorm:"column:deleted_at"`
	UpdatedAt        *string `gorm:"column:updated_at"`
}
