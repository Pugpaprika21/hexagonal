package schema

import "database/sql"

type GetUsersGroupSetting struct {
	ID        sql.NullInt64  `gorm:"column:id"`
	UserID    sql.NullInt64  `gorm:"column:user_id"`
	GroupID   sql.NullInt64  `gorm:"column:group_id"`
	CreatedAt sql.NullString `gorm:"column:created_at"`
	DeletedAt sql.NullString `gorm:"column:deleted_at"`
	UpdatedAt sql.NullString `gorm:"column:updated_at"`
}

type GetUsersGroupSettingByUserID struct {
	ID         sql.NullInt64  `gorm:"column:id"`
	GroupName  sql.NullString `gorm:"column:group_name"`
	HasInGroup sql.NullInt64  `gorm:"column:has_in_group"`
	CreatedAt  sql.NullString `gorm:"column:created_at"`
	DeletedAt  sql.NullString `gorm:"column:deleted_at"`
	UpdatedAt  sql.NullString `gorm:"column:updated_at"`
}

type CreateUsersGroupSetting struct {
	ID        *int    `gorm:"column:id"`
	UserID    *int    `gorm:"column:user_id"`
	GroupID   *int    `gorm:"column:group_id"`
	CreatedAt *string `gorm:"column:created_at"`
	DeletedAt *string `gorm:"column:deleted_at"`
	UpdatedAt *string `gorm:"column:updated_at"`
}

type UpdateUsersGroupSetting struct {
	ID        *int    `gorm:"column:id"`
	UserID    *int    `gorm:"column:user_id"`
	GroupID   *int    `gorm:"column:group_id"`
	CreatedAt *string `gorm:"column:created_at"`
	DeletedAt *string `gorm:"column:deleted_at"`
	UpdatedAt *string `gorm:"column:updated_at"`
}
