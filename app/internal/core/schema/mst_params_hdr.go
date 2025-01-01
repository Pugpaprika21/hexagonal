package schema

import "database/sql"

type GetParamsHdr struct {
	ParamID       sql.NullInt64  `gorm:"column:param_id"`
	ParamkeyTxt   sql.NullString `gorm:"column:paramkey_txt"`
	ParamCode     sql.NullString `gorm:"column:param_code"`
	ParamNameth   sql.NullString `gorm:"column:param_nameth"`
	ParamNameeng  sql.NullString `gorm:"column:param_nameeng"`
	ParamDetail   sql.NullString `gorm:"column:param_detail"`
	ParamOrder    sql.NullInt32  `gorm:"column:param_order"`
	IsActive      sql.NullBool   `gorm:"column:is_active"`
	CreatedProgby sql.NullString `gorm:"column:created_progby"`
	CreatedBy     sql.NullString `gorm:"column:created_by"`
	UpdatedBy     sql.NullString `gorm:"column:updated_by"`
	CreatedAt     sql.NullString `gorm:"column:created_at"`
	UpdatedAt     sql.NullString `gorm:"column:updated_at"`
}

type LovParamsHdr struct {
	ParamID       sql.NullInt64  `gorm:"column:param_id"`
	ParamkeyTxt   sql.NullString `gorm:"column:paramkey_txt"`
	ParamCode     sql.NullString `gorm:"column:param_code"`
	ParamNameth   sql.NullString `gorm:"column:param_nameth"`
	ParamNameeng  sql.NullString `gorm:"column:param_nameeng"`
	ParamDetail   sql.NullString `gorm:"column:param_detail"`
	ParamOrder    sql.NullInt32  `gorm:"column:param_order"`
	IsActive      sql.NullBool   `gorm:"column:is_active"`
	CreatedProgby sql.NullString `gorm:"column:created_progby"`
	CreatedBy     sql.NullString `gorm:"column:created_by"`
	UpdatedBy     sql.NullString `gorm:"column:updated_by"`
	CreatedAt     sql.NullString `gorm:"column:created_at"`
	UpdatedAt     sql.NullString `gorm:"column:updated_at"`
}

type CreateParamsHdr struct {
	ParamID       *int64  `gorm:"column:param_id"`
	ParamkeyTxt   *string `gorm:"column:paramkey_txt"`
	ParamCode     *string `gorm:"column:param_code"`
	ParamNameth   *string `gorm:"column:param_nameth"`
	ParamNameeng  *string `gorm:"column:param_nameeng"`
	ParamDetail   *string `gorm:"column:param_detail"`
	ParamOrder    *int32  `gorm:"column:param_order"`
	IsActive      *bool   `gorm:"column:is_active"`
	CreatedProgby *string `gorm:"column:created_progby"`
	CreatedBy     *string `gorm:"column:created_by"`
	UpdatedBy     *string `gorm:"column:updated_by"`
	CreatedAt     *string `gorm:"column:created_at"`
	UpdatedAt     *string `gorm:"column:updated_at"`
}

type UpdateParamsHdr struct {
	ParamID       *int64  `gorm:"column:param_id"`
	ParamkeyTxt   *string `gorm:"column:paramkey_txt"`
	ParamCode     *string `gorm:"column:param_code"`
	ParamNameth   *string `gorm:"column:param_nameth"`
	ParamNameeng  *string `gorm:"column:param_nameeng"`
	ParamDetail   *string `gorm:"column:param_detail"`
	ParamOrder    *int32  `gorm:"column:param_order"`
	IsActive      *bool   `gorm:"column:is_active"`
	CreatedProgby *string `gorm:"column:created_progby"`
	CreatedBy     *string `gorm:"column:created_by"`
	UpdatedBy     *string `gorm:"column:updated_by"`
	CreatedAt     *string `gorm:"column:created_at"`
	UpdatedAt     *string `gorm:"column:updated_at"`
}
