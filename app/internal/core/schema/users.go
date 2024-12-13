package schema

import "database/sql"

type GetUsers struct {
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
}
