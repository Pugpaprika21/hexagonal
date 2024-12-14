package response

type GetUsersGroupSetting struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	GroupID   int    `json:"group_id"`
	CreatedAt string `json:"created_at"`
	DeletedAt string `json:"deleted_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetUsersGroupSettingByUserID struct {
	ID         int    `json:"id"`
	GroupName  string `json:"group_name"`
	HasInGroup int    `json:"has_in_group"`
	CreatedAt  string `json:"created_at"`
	DeletedAt  string `json:"deleted_at"`
	UpdatedAt  string `json:"updated_at"`
}
