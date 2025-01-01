package request

type GetParamsDtl struct {
	ParamID       *int64  `json:"param_id"`
	ParamkeyID    *int32  `json:"paramkey_id"`
	ParamparentID *int32  `json:"paramparent_id"`
	ParamkeyTxt   *string `json:"paramkey_txt"`
	ParamCode     *string `json:"param_code"`
	ParamNameth   *string `json:"param_nameth"`
	ParamNameeng  *string `json:"param_nameeng"`
	ParamDetail   *string `json:"param_detail"`
	ParamOrder    *int32  `json:"param_order"`
	ParamhdrID    *int32  `json:"paramhdr_id"`
	IsActive      *bool   `json:"is_active"`
	CreatedProgby *string `json:"created_progby"`
	CreatedBy     *string `json:"created_by"`
	UpdatedBy     *string `json:"updated_by"`
	CreatedAt     *string `json:"created_at"`
	UpdatedAt     *string `json:"updated_at"`
}

type LovParamsDtl struct {
	ParamID       *int64  `json:"param_id"`
	ParamkeyID    *int32  `json:"paramkey_id"`
	ParamparentID *int32  `json:"paramparent_id"`
	ParamkeyTxt   *string `json:"paramkey_txt"`
	ParamCode     *string `json:"param_code"`
	ParamNameth   *string `json:"param_nameth"`
	ParamNameeng  *string `json:"param_nameeng"`
	ParamDetail   *string `json:"param_detail"`
	ParamOrder    *int32  `json:"param_order"`
	ParamhdrID    *int32  `json:"paramhdr_id"`
	IsActive      *bool   `json:"is_active"`
	CreatedProgby *string `json:"created_progby"`
	CreatedBy     *string `json:"created_by"`
	UpdatedBy     *string `json:"updated_by"`
	CreatedAt     *string `json:"created_at"`
	UpdatedAt     *string `json:"updated_at"`
}

type (
	CreateParamsDtlRows struct {
		ParamID       *int64  `json:"param_id"`
		ParamkeyID    *int32  `json:"paramkey_id"`
		ParamparentID *int32  `json:"paramparent_id"`
		ParamkeyTxt   *string `json:"paramkey_txt"`
		ParamCode     *string `json:"param_code"`
		ParamNameth   *string `json:"param_nameth"`
		ParamNameeng  *string `json:"param_nameeng"`
		ParamDetail   *string `json:"param_detail"`
		ParamOrder    *int32  `json:"param_order"`
		ParamhdrID    *int32  `json:"paramhdr_id"`
		IsActive      *bool   `json:"is_active"`
		CreatedProgby *string `json:"created_progby"`
		CreatedBy     *string `json:"created_by"`
		UpdatedBy     *string `json:"updated_by"`
		CreatedAt     *string `json:"created_at"`
		UpdatedAt     *string `json:"updated_at"`
	}

	CreateParamsDtl struct {
		CreateParamsDtlRows []CreateParamsDtlRows `json:"cre_rows"`
	}
)

type (
	UpdateParamsDtlRows struct {
		ParamID       *int64  `json:"param_id"`
		ParamkeyID    *int32  `json:"paramkey_id"`
		ParamparentID *int32  `json:"paramparent_id"`
		ParamkeyTxt   *string `json:"paramkey_txt"`
		ParamCode     *string `json:"param_code"`
		ParamNameth   *string `json:"param_nameth"`
		ParamNameeng  *string `json:"param_nameeng"`
		ParamDetail   *string `json:"param_detail"`
		ParamOrder    *int32  `json:"param_order"`
		ParamhdrID    *int32  `json:"paramhdr_id"`
		IsActive      *bool   `json:"is_active"`
		CreatedProgby *string `json:"created_progby"`
		CreatedBy     *string `json:"created_by"`
		UpdatedBy     *string `json:"updated_by"`
		CreatedAt     *string `json:"created_at"`
		UpdatedAt     *string `json:"updated_at"`
	}

	UpdateParamsDtl struct {
		UpdateParamsDtlRows []UpdateParamsDtlRows `json:"upd_rows"`
	}
)

type DeleteParamsDtl struct {
	ParamID       *int64  `json:"param_id"`
	ParamkeyID    *int32  `json:"paramkey_id"`
	ParamparentID *int32  `json:"paramparent_id"`
	ParamkeyTxt   *string `json:"paramkey_txt"`
	ParamCode     *string `json:"param_code"`
	ParamNameth   *string `json:"param_nameth"`
	ParamNameeng  *string `json:"param_nameeng"`
	ParamDetail   *string `json:"param_detail"`
	ParamOrder    *int32  `json:"param_order"`
	ParamhdrID    *int32  `json:"paramhdr_id"`
	IsActive      *bool   `json:"is_active"`
	CreatedProgby *string `json:"created_progby"`
	CreatedBy     *string `json:"created_by"`
	UpdatedBy     *string `json:"updated_by"`
	CreatedAt     *string `json:"created_at"`
	UpdatedAt     *string `json:"updated_at"`
}
