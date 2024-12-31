package schema

import "database/sql"

type GetMainMenus struct {
	ID            sql.NullInt32  `gorm:"column:id"`
	UserID        sql.NullInt32  `gorm:"column:user_id"`
	RoleID        sql.NullInt32  `gorm:"column:role_id"`
	Name          sql.NullString `gorm:"column:name"`
	NameEn        sql.NullString `gorm:"column:name_en"`
	Url           sql.NullString `gorm:"column:url"`
	Icon          sql.NullString `gorm:"column:icon"`
	Tooltip       sql.NullString `gorm:"column:tooltip"`
	ParentID      sql.NullInt32  `gorm:"column:parent_id"`
	Position      sql.NullInt32  `gorm:"column:position"`
	IsActive      sql.NullBool   `gorm:"column:is_active"`
	IsExternal    sql.NullBool   `gorm:"column:is_external"`
	PermissionKey sql.NullString `gorm:"column:permission_key"`
	CreatedAt     sql.NullString `gorm:"column:created_at"`
	UpdatedAt     sql.NullString `gorm:"column:updated_at"`
}

type GetParentMenus struct {
	ID            sql.NullInt32  `gorm:"column:id"`
	UserID        sql.NullInt32  `gorm:"column:user_id"`
	RoleID        sql.NullInt32  `gorm:"column:role_id"`
	Name          sql.NullString `gorm:"column:name"`
	NameEn        sql.NullString `gorm:"column:name_en"`
	Url           sql.NullString `gorm:"column:url"`
	Icon          sql.NullString `gorm:"column:icon"`
	Tooltip       sql.NullString `gorm:"column:tooltip"`
	ParentID      sql.NullInt32  `gorm:"column:parent_id"`
	Position      sql.NullInt32  `gorm:"column:position"`
	IsActive      sql.NullBool   `gorm:"column:is_active"`
	IsExternal    sql.NullBool   `gorm:"column:is_external"`
	PermissionKey sql.NullString `gorm:"column:permission_key"`
	CreatedAt     sql.NullString `gorm:"column:created_at"`
	UpdatedAt     sql.NullString `gorm:"column:updated_at"`
}

type GetChildMenus struct {
	ID            sql.NullInt32  `gorm:"column:id"`
	UserID        sql.NullInt32  `gorm:"column:user_id"`
	RoleID        sql.NullInt32  `gorm:"column:role_id"`
	Name          sql.NullString `gorm:"column:name"`
	NameEn        sql.NullString `gorm:"column:name_en"`
	Url           sql.NullString `gorm:"column:url"`
	Icon          sql.NullString `gorm:"column:icon"`
	Tooltip       sql.NullString `gorm:"column:tooltip"`
	ParentID      sql.NullInt32  `gorm:"column:parent_id"`
	Position      sql.NullInt32  `gorm:"column:position"`
	IsActive      sql.NullBool   `gorm:"column:is_active"`
	IsExternal    sql.NullBool   `gorm:"column:is_external"`
	PermissionKey sql.NullString `gorm:"column:permission_key"`
	CreatedAt     sql.NullString `gorm:"column:created_at"`
	UpdatedAt     sql.NullString `gorm:"column:updated_at"`
}

type GetNewMainManus struct {
	MenuID          sql.NullInt32  `gorm:"column:menu_id"`
	MenuName        sql.NullString `gorm:"column:menu_name"`
	MenuURL         sql.NullString `gorm:"column:menu_url"`
	MenuIcon        sql.NullString `gorm:"column:menu_icon"`
	MenuTooltip     sql.NullString `gorm:"column:menu_tooltip"`
	MenuPosition    sql.NullInt32  `gorm:"column:menu_position"`
	MenuIsActive    sql.NullBool   `gorm:"column:menu_is_active"`
	SubMenuID       sql.NullInt32  `gorm:"column:sub_menu_id"`
	SubMenuName     sql.NullString `gorm:"column:sub_menu_name"`
	SubMenuURL      sql.NullString `gorm:"column:sub_menu_url"`
	SubMenuIcon     sql.NullString `gorm:"column:sub_menu_icon"`
	SubMenuTooltip  sql.NullString `gorm:"column:sub_menu_tooltip"`
	SubMenuPosition sql.NullInt32  `gorm:"column:sub_menu_position"`
	SubMenuIsActive sql.NullBool   `gorm:"column:sub_menu_is_active"`
	UserID          sql.NullInt32  `gorm:"column:user_id"`
	RoleID          sql.NullInt32  `gorm:"column:role_id"`
}

type CreateMenus struct {
	ID            *int32  `gorm:"column:id"`
	UserID        *int32  `gorm:"column:user_id"`
	RoleID        *int32  `gorm:"column:role_id"`
	Name          *string `gorm:"column:name"`
	NameEn        *string `gorm:"column:name_en"`
	Url           *string `gorm:"column:url"`
	Icon          *string `gorm:"column:icon"`
	Tooltip       *string `gorm:"column:tooltip"`
	ParentID      *int32  `gorm:"column:parent_id"`
	Position      *int32  `gorm:"column:position"`
	IsActive      *bool   `gorm:"column:is_active"`
	IsExternal    *bool   `gorm:"column:is_external"`
	PermissionKey *string `gorm:"column:permission_key"`
	CreatedAt     *string `gorm:"column:created_at"`
	UpdatedAt     *string `gorm:"column:updated_at"`
}

type UpdateMenus struct {
	ID            *int32  `gorm:"column:id"`
	UserID        *int32  `gorm:"column:user_id"`
	RoleID        *int32  `gorm:"column:role_id"`
	Name          *string `gorm:"column:name"`
	NameEn        *string `gorm:"column:name_en"`
	Url           *string `gorm:"column:url"`
	Icon          *string `gorm:"column:icon"`
	Tooltip       *string `gorm:"column:tooltip"`
	ParentID      *int32  `gorm:"column:parent_id"`
	Position      *int32  `gorm:"column:position"`
	IsActive      *bool   `gorm:"column:is_active"`
	IsExternal    *bool   `gorm:"column:is_external"`
	PermissionKey *string `gorm:"column:permission_key"`
	CreatedAt     *string `gorm:"column:created_at"`
	UpdatedAt     *string `gorm:"column:updated_at"`
}
