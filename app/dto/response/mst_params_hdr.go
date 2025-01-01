package response

type GetParamsHdr struct {
	ParamID       int64  `json:"param_id"`
	ParamkeyTxt   string `json:"paramkey_txt"`
	ParamCode     string `json:"param_code"`
	ParamNameth   string `json:"param_nameth"`
	ParamNameeng  string `json:"param_nameeng"`
	ParamDetail   string `json:"param_detail"`
	ParamOrder    int32  `json:"param_order"`
	IsActive      bool   `json:"is_active"`
	CreatedProgby string `json:"created_progby"`
	CreatedBy     string `json:"created_by"`
	UpdatedBy     string `json:"updated_by"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type LovParamsHdr struct {
	ParamID       int64  `json:"param_id"`
	ParamkeyTxt   string `json:"paramkey_txt"`
	ParamCode     string `json:"param_code"`
	ParamNameth   string `json:"param_nameth"`
	ParamNameeng  string `json:"param_nameeng"`
	ParamDetail   string `json:"param_detail"`
	ParamOrder    int32  `json:"param_order"`
	IsActive      bool   `json:"is_active"`
	CreatedProgby string `json:"created_progby"`
	CreatedBy     string `json:"created_by"`
	UpdatedBy     string `json:"updated_by"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
