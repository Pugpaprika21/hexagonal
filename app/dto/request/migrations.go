package request

type GetGoStructProcedure struct {
	Tablename string `json:"tablename"`
	StName    string `json:"stname"`
}

type GetAllGoStructProcedure struct {
	Tablename string `json:"tablename"`
	StName    string `json:"stname"`
}
