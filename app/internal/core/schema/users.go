package schema

import "database/sql"

type GetUsers struct {
	RowNum       sql.NullInt64  `gorm:"column:row_num"`
	Address      sql.NullString `gorm:"column:address"`
	CreatedAt    sql.NullString `gorm:"column:created_at"`
	DateOfBirth  sql.NullString `gorm:"column:date_of_birth"`
	DeletedAt    sql.NullString `gorm:"column:deleted_at"`
	Email        sql.NullString `gorm:"column:email"`
	FirstName    sql.NullString `gorm:"column:first_name"`
	ID           sql.NullInt64  `gorm:"column:id"`
	LastName     sql.NullString `gorm:"column:last_name"`
	Password     sql.NullString `gorm:"column:password"`
	PasswordHash sql.NullString `gorm:"column:password_hash"`
	PhoneNumber  sql.NullString `gorm:"column:phone_number"`
	UpdatedAt    sql.NullString `gorm:"column:updated_at"`
	Username     sql.NullString `gorm:"column:username"`
	FullName     sql.NullString `gorm:"column:full_name"`
}

type CreateUsers struct {
	ID           *int    `gorm:"column:id"`
	Address      *string `gorm:"column:address"`
	CreatedAt    *string `gorm:"column:created_at"`
	DateOfBirth  *string `gorm:"column:date_of_birth"`
	DeletedAt    *string `gorm:"column:deleted_at"`
	Email        *string `gorm:"column:email"`
	FirstName    *string `gorm:"column:first_name"`
	LastName     *string `gorm:"column:last_name"`
	Password     *string `gorm:"column:password"`
	PasswordHash *string `gorm:"column:password_hash"`
	PhoneNumber  *string `gorm:"column:phone_number"`
	UpdatedAt    *string `gorm:"column:updated_at"`
	Username     *string `gorm:"column:username"`
}

type UpdateUsers struct {
	ID           *int    `gorm:"column:id"`
	Address      *string `gorm:"column:address"`
	CreatedAt    *string `gorm:"column:created_at"`
	DateOfBirth  *string `gorm:"column:date_of_birth"`
	DeletedAt    *string `gorm:"column:deleted_at"`
	Email        *string `gorm:"column:email"`
	FirstName    *string `gorm:"column:first_name"`
	LastName     *string `gorm:"column:last_name"`
	Password     *string `gorm:"column:password"`
	PasswordHash *string `gorm:"column:password_hash"`
	PhoneNumber  *string `gorm:"column:phone_number"`
	UpdatedAt    *string `gorm:"column:updated_at"`
	Username     *string `gorm:"column:username"`
}
