package respone

type GetUsers struct {
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
