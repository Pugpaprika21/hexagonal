package schema

import "database/sql"

type GetParamsDtl struct {
	ParamID       sql.NullInt64  `gorm:"column:param_id"`
	ParamkeyID    sql.NullInt32  `gorm:"column:paramkey_id"`
	ParamparentID sql.NullInt32  `gorm:"column:paramparent_id"`
	ParamkeyTxt   sql.NullString `gorm:"column:paramkey_txt"`
	ParamCode     sql.NullString `gorm:"column:param_code"`
	ParamNameth   sql.NullString `gorm:"column:param_nameth"`
	ParamNameeng  sql.NullString `gorm:"column:param_nameeng"`
	ParamDetail   sql.NullString `gorm:"column:param_detail"`
	ParamOrder    sql.NullInt32  `gorm:"column:param_order"`
	ParamhdrID    sql.NullInt32  `gorm:"column:paramhdr_id"`
	IsActive      sql.NullBool   `gorm:"column:is_active"`
	CreatedProgby sql.NullString `gorm:"column:created_progby"`
	CreatedBy     sql.NullString `gorm:"column:created_by"`
	UpdatedBy     sql.NullString `gorm:"column:updated_by"`
	CreatedAt     sql.NullString `gorm:"column:created_at"`
	UpdatedAt     sql.NullString `gorm:"column:updated_at"`
}

type LovParamsDtl struct {
	ParamID       sql.NullInt64  `gorm:"column:param_id"`
	ParamkeyID    sql.NullInt32  `gorm:"column:paramkey_id"`
	ParamparentID sql.NullInt32  `gorm:"column:paramparent_id"`
	ParamkeyTxt   sql.NullString `gorm:"column:paramkey_txt"`
	ParamCode     sql.NullString `gorm:"column:param_code"`
	ParamNameth   sql.NullString `gorm:"column:param_nameth"`
	ParamNameeng  sql.NullString `gorm:"column:param_nameeng"`
	ParamDetail   sql.NullString `gorm:"column:param_detail"`
	ParamOrder    sql.NullInt32  `gorm:"column:param_order"`
	ParamhdrID    sql.NullInt32  `gorm:"column:paramhdr_id"`
	IsActive      sql.NullBool   `gorm:"column:is_active"`
	CreatedProgby sql.NullString `gorm:"column:created_progby"`
	CreatedBy     sql.NullString `gorm:"column:created_by"`
	UpdatedBy     sql.NullString `gorm:"column:updated_by"`
	CreatedAt     sql.NullString `gorm:"column:created_at"`
	UpdatedAt     sql.NullString `gorm:"column:updated_at"`
}

type CreateParamsDtl struct {
	ParamID       *int64  `gorm:"column:param_id"`
	ParamkeyID    *int32  `gorm:"column:paramkey_id"`
	ParamparentID *int32  `gorm:"column:paramparent_id"`
	ParamkeyTxt   *string `gorm:"column:paramkey_txt"`
	ParamCode     *string `gorm:"column:param_code"`
	ParamNameth   *string `gorm:"column:param_nameth"`
	ParamNameeng  *string `gorm:"column:param_nameeng"`
	ParamDetail   *string `gorm:"column:param_detail"`
	ParamOrder    *int32  `gorm:"column:param_order"`
	ParamhdrID    *int32  `gorm:"column:paramhdr_id"`
	IsActive      *bool   `gorm:"column:is_active"`
	CreatedProgby *string `gorm:"column:created_progby"`
	CreatedBy     *string `gorm:"column:created_by"`
	UpdatedBy     *string `gorm:"column:updated_by"`
	CreatedAt     *string `gorm:"column:created_at"`
	UpdatedAt     *string `gorm:"column:updated_at"`
}

type UpdateParamsDtl struct {
	ParamID       *int64  `gorm:"column:param_id"`
	ParamkeyID    *int32  `gorm:"column:paramkey_id"`
	ParamparentID *int32  `gorm:"column:paramparent_id"`
	ParamkeyTxt   *string `gorm:"column:paramkey_txt"`
	ParamCode     *string `gorm:"column:param_code"`
	ParamNameth   *string `gorm:"column:param_nameth"`
	ParamNameeng  *string `gorm:"column:param_nameeng"`
	ParamDetail   *string `gorm:"column:param_detail"`
	ParamOrder    *int32  `gorm:"column:param_order"`
	ParamhdrID    *int32  `gorm:"column:paramhdr_id"`
	IsActive      *bool   `gorm:"column:is_active"`
	CreatedProgby *string `gorm:"column:created_progby"`
	CreatedBy     *string `gorm:"column:created_by"`
	UpdatedBy     *string `gorm:"column:updated_by"`
	CreatedAt     *string `gorm:"column:created_at"`
	UpdatedAt     *string `gorm:"column:updated_at"`
}
