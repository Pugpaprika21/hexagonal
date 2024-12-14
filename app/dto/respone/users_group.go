package respone

type GetUsersGroup struct {
	ID               int    `json:"id"`
	GroupCode        string `json:"group_code"`
	GroupName        string `json:"group_name"`
	GroupDescription string `json:"group_description"`
	CreatedAt        string `json:"created_at"`
	DeletedAt        string `json:"deleted_at"`
	UpdatedAt        string `json:"updated_at"`
}
