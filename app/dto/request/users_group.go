package request

type GetUsersGroup struct {
	ID               int    `json:"id"`
	GroupCode        string `json:"group_code"`
	GroupName        string `json:"group_name"`
	GroupDescription string `json:"group_description"`
	CreatedAt        string `json:"created_at"`
	DeletedAt        string `json:"deleted_at"`
	UpdatedAt        string `json:"updated_at"`
}

type (
	CreateUsersGroupRows struct {
		ID               *int    `json:"id"`
		GroupCode        *string `json:"group_code"`
		GroupName        *string `json:"group_name"`
		GroupDescription *string `json:"group_description"`
		CreatedAt        *string `json:"created_at"`
		DeletedAt        *string `json:"deleted_at"`
		UpdatedAt        *string `json:"updated_at"`
	}

	CreateUsersGroup struct {
		CreateUsersGroupRows []CreateUsersGroupRows `json:"users_group_rows"`
	}
)

type (
	UpdateUsersGroupRows struct {
		ID               *int    `json:"id"`
		GroupCode        *string `json:"group_code"`
		GroupName        *string `json:"group_name"`
		GroupDescription *string `json:"group_description"`
		CreatedAt        *string `json:"created_at"`
		DeletedAt        *string `json:"deleted_at"`
		UpdatedAt        *string `json:"updated_at"`
	}

	UpdateUsersGroup struct {
		UpdateUsersGroupRows []UpdateUsersGroupRows `json:"users_group_rows"`
	}
)

type DeleteUsersGroup struct {
	ID               int    `json:"id"`
	GroupCode        string `json:"group_code"`
	GroupName        string `json:"group_name"`
	GroupDescription string `json:"group_description"`
	CreatedAt        string `json:"created_at"`
	DeletedAt        string `json:"deleted_at"`
	UpdatedAt        string `json:"updated_at"`
}
