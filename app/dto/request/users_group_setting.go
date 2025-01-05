package request

type GetUsersGroupSetting struct {
	ID        *int    `json:"id"`
	UserID    *int    `json:"user_id"`
	GroupID   *int    `json:"group_id"`
	CreatedAt *string `json:"created_at"`
	DeletedAt *string `json:"deleted_at"`
	UpdatedAt *string `json:"updated_at"`
}

type GetUsersGroupSettingByUserID struct {
	ID         *int    `json:"id"`
	GroupName  *string `json:"group_name"`
	HasInGroup *int    `json:"has_in_group"`
	UserID     *int    `json:"user_id"`
	CreatedAt  *string `json:"created_at"`
	DeletedAt  *string `json:"deleted_at"`
	UpdatedAt  *string `json:"updated_at"`
}

type (
	CreateUsersGroupSettingRows struct {
		ID        *int    `json:"id"`
		UserID    *int    `json:"user_id"`
		GroupID   *int    `json:"group_id"`
		CreatedAt *string `json:"created_at"`
		DeletedAt *string `json:"deleted_at"`
		UpdatedAt *string `json:"updated_at"`
	}

	CreateUsersGroupSetting struct {
		CreateUsersGroupSettingRows []CreateUsersGroupSettingRows `json:"users_group_setting_rows"`
	}
)

type (
	UpdateUsersGroupSettingRows struct {
		ID        *int    `json:"id"`
		UserID    *int    `json:"user_id"`
		GroupID   *int    `json:"group_id"`
		CreatedAt *string `json:"created_at"`
		DeletedAt *string `json:"deleted_at"`
		UpdatedAt *string `json:"updated_at"`
	}

	UpdateUsersGroupSetting struct {
		UpdateUsersGroupSettingRows []UpdateUsersGroupSettingRows `json:"users_group_setting_rows"`
	}
)

type DeleteUsersSettingGroup struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	GroupID   int    `json:"group_id"`
	CreatedAt string `json:"created_at"`
	DeletedAt string `json:"deleted_at"`
	UpdatedAt string `json:"updated_at"`
}
