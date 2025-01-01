package request

type GetMainMenus struct {
	ID            *int32  `json:"id"`
	UserID        *int32  `json:"user_id"`
	RoleID        *int32  `json:"role_id"`
	Name          *string `json:"name"`
	NameEn        *string `json:"name_en"`
	Url           *string `json:"url"`
	Icon          *string `json:"icon"`
	Tooltip       *string `json:"tooltip"`
	ParentID      *int32  `json:"parent_id"`
	Position      *int32  `json:"position"`
	IsActive      *bool   `json:"is_active"`
	IsExternal    *bool   `json:"is_external"`
	PermissionKey *string `json:"permission_key"`
	CreatedAt     *string `json:"created_at"`
	UpdatedAt     *string `json:"updated_at"`
}

type GetNewMainManus struct {
	MenuID          *int32  `json:"menu_id"`
	MenuName        *string `json:"menu_name"`
	MenuURL         *string `json:"menu_url"`
	MenuIcon        *string `json:"menu_icon"`
	MenuTooltip     *string `json:"menu_tooltip"`
	MenuPosition    *int32  `json:"menu_position"`
	MenuIsActive    *bool   `json:"menu_is_active"`
	SubMenuID       *int32  `json:"sub_menu_id"`
	SubMenuName     *string `json:"sub_menu_name"`
	SubMenuURL      *string `json:"sub_menu_url"`
	SubMenuIcon     *string `json:"sub_menu_icon"`
	SubMenuTooltip  *string `json:"sub_menu_tooltip"`
	SubMenuPosition *int32  `json:"sub_menu_position"`
	SubMenuIsActive *bool   `json:"sub_menu_is_active"`
	UserID          *int32  `json:"user_id"`
	RoleID          *int32  `json:"role_id"`
}

type GetAllMenus struct {
	ID            *int32  `json:"id"`
	UserID        *int32  `json:"user_id"`
	RoleID        *int32  `json:"role_id"`
	Name          *string `json:"name"`
	NameEn        *string `json:"name_en"`
	Url           *string `json:"url"`
	Icon          *string `json:"icon"`
	Tooltip       *string `json:"tooltip"`
	ParentID      *int32  `json:"parent_id"`
	Position      *int32  `json:"position"`
	IsActive      *bool   `json:"is_active"`
	IsExternal    *bool   `json:"is_external"`
	PermissionKey *string `json:"permission_key"`
	CreatedAt     *string `json:"created_at"`
	UpdatedAt     *string `json:"updated_at"`
}

type DeleteMenus struct {
	ID            *int32  `json:"id"`
	UserID        *int32  `json:"user_id"`
	RoleID        *int32  `json:"role_id"`
	Name          *string `json:"name"`
	NameEn        *string `json:"name_en"`
	Url           *string `json:"url"`
	Icon          *string `json:"icon"`
	Tooltip       *string `json:"tooltip"`
	ParentID      *int32  `json:"parent_id"`
	Position      *int32  `json:"position"`
	IsActive      *bool   `json:"is_active"`
	IsExternal    *bool   `json:"is_external"`
	PermissionKey *string `json:"permission_key"`
	CreatedAt     *string `json:"created_at"`
	UpdatedAt     *string `json:"updated_at"`
}

type (
	CreateMenusRows struct {
		ID            *int32  `json:"id"`
		UserID        *int32  `json:"user_id"`
		RoleID        *int32  `json:"role_id"`
		Name          *string `json:"name"`
		NameEn        *string `json:"name_en"`
		Url           *string `json:"url"`
		Icon          *string `json:"icon"`
		Tooltip       *string `json:"tooltip"`
		ParentID      *int32  `json:"parent_id"`
		Position      *int32  `json:"position"`
		IsActive      *bool   `json:"is_active"`
		IsExternal    *bool   `json:"is_external"`
		PermissionKey *string `json:"permission_key"`
		CreatedAt     *string `json:"created_at"`
		UpdatedAt     *string `json:"updated_at"`
	}

	CreateMenus struct {
		CreateMenusRows []CreateMenusRows `json:"cre_rows"`
	}
)

type (
	UpdateMenusRows struct {
		ID            *int32  `json:"id"`
		UserID        *int32  `json:"user_id"`
		RoleID        *int32  `json:"role_id"`
		Name          *string `json:"name"`
		NameEn        *string `json:"name_en"`
		Url           *string `json:"url"`
		Icon          *string `json:"icon"`
		Tooltip       *string `json:"tooltip"`
		ParentID      *int32  `json:"parent_id"`
		Position      *int32  `json:"position"`
		IsActive      *bool   `json:"is_active"`
		IsExternal    *bool   `json:"is_external"`
		PermissionKey *string `json:"permission_key"`
		CreatedAt     *string `json:"created_at"`
		UpdatedAt     *string `json:"updated_at"`
	}

	UpdateMenus struct {
		UpdateMenusRows []UpdateMenusRows `json:"upd_rows"`
	}
)
