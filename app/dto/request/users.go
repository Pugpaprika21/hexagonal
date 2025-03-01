package request

type GetUsers struct {
	LazyLoad     LazyLoad `json:"lazy_load"`
	Address      *string  `json:"address"`
	CreatedAt    *string  `json:"created_at"`
	DateOfBirth  *string  `json:"date_of_birth"`
	DeletedAt    *string  `json:"deleted_at"`
	Email        *string  `json:"email"`
	FirstName    *string  `json:"first_name"`
	ID           *int     `json:"id"`
	LastName     *string  `json:"last_name"`
	Password     *string  `json:"password"`
	PasswordHash *string  `json:"password_hash"`
	PhoneNumber  *string  `json:"phone_number"`
	UpdatedAt    *string  `json:"updated_at"`
	Username     *string  `json:"username"`
}

type (
	CreateUsersRows struct {
		Address      *string `json:"address"`
		CreatedAt    *string `json:"created_at"`
		DateOfBirth  *string `json:"date_of_birth"`
		DeletedAt    *string `json:"deleted_at"`
		Email        *string `json:"email"`
		FirstName    *string `json:"first_name"`
		ID           *int    `json:"id"`
		LastName     *string `json:"last_name"`
		Password     *string `json:"password"`
		PasswordHash *string `json:"password_hash"`
		PhoneNumber  *string `json:"phone_number"`
		UpdatedAt    *string `json:"updated_at"`
		Username     *string `json:"username"`
	}

	CreateUsers struct {
		CreateUsersRows []CreateUsersRows `json:"users_rows"`
	}
)

type (
	UpdateUsersRows struct {
		Address      *string `json:"address"`
		CreatedAt    *string `json:"created_at"`
		DateOfBirth  *string `json:"date_of_birth"`
		DeletedAt    *string `json:"deleted_at"`
		Email        *string `json:"email"`
		FirstName    *string `json:"first_name"`
		ID           *int    `json:"id"`
		LastName     *string `json:"last_name"`
		Password     *string `json:"password"`
		PasswordHash *string `json:"password_hash"`
		PhoneNumber  *string `json:"phone_number"`
		UpdatedAt    *string `json:"updated_at"`
		Username     *string `json:"username"`
	}

	UpdateUsers struct {
		UpdateUsersRows []UpdateUsersRows `json:"users_rows"`
	}
)

type DeleteUsers struct {
	Address      string `json:"address"`
	CreatedAt    string `json:"created_at"`
	DateOfBirth  string `json:"date_of_birth"`
	DeletedAt    string `json:"deleted_at"`
	Email        string `json:"email"`
	FirstName    string `json:"first_name"`
	ID           int    `json:"id"`
	LastName     string `json:"last_name"`
	Password     string `json:"password"`
	PasswordHash string `json:"password_hash"`
	PhoneNumber  string `json:"phone_number"`
	UpdatedAt    string `json:"updated_at"`
	Username     string `json:"username"`
}
