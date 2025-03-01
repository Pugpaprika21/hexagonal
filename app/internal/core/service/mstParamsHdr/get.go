package mstparamshdr

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/dto/response"
	"pugpaprika/app/pkg/sqlx"
	"strings"
)

func (m *mstParamsHdrService) GetParamsHdr(ctx context.Context, req request.GetParamsHdr) ([]response.GetParamsHdr, error) {
	var sql sqlx.Sqlx

	sql.Stmt = `select param_id ,paramkey_txt ,param_code ,param_nameth ,param_nameeng ,param_detail ,param_order from mst_params_hdr `

	if req.ParamID != nil && *req.ParamID != 0 {
		sql.WhereClause = append(sql.WhereClause, "param_id = ?")
		sql.Args = append(sql.Args, req.ParamID)
	}

	if req.ParamkeyTxt != nil && *req.ParamkeyTxt != "" {
		sql.WhereClause = append(sql.WhereClause, "paramkey_txt = ?")
		sql.Args = append(sql.Args, req.ParamkeyTxt)
	}

	if req.ParamCode != nil && *req.ParamCode != "" {
		sql.WhereClause = append(sql.WhereClause, "param_code = ?")
		sql.Args = append(sql.Args, req.ParamCode)
	}

	sql.WhereClause = append(sql.WhereClause, "is_active = ?")
	sql.Args = append(sql.Args, true)

	if len(sql.WhereClause) > 0 {
		sql.Stmt += " Where " + strings.Join(sql.WhereClause, " and ")
	}

	sql.Stmt += ` order by param_id desc;`

	rows, err := m.repos.GetParamsHdr(ctx, sql)
	if err != nil {
		return nil, err
	}

	resp := make([]response.GetParamsHdr, len(rows))
	for i, rec := range rows {
		resp[i] = response.GetParamsHdr{
			ParamID:       rec.ParamID.Int64,
			ParamkeyTxt:   rec.ParamkeyTxt.String,
			ParamCode:     rec.ParamCode.String,
			ParamNameth:   rec.ParamNameth.String,
			ParamNameeng:  rec.ParamNameeng.String,
			ParamDetail:   rec.ParamDetail.String,
			ParamOrder:    rec.ParamOrder.Int32,
			IsActive:      rec.IsActive.Bool,
			CreatedProgby: rec.CreatedProgby.String,
			CreatedBy:     rec.CreatedBy.String,
			UpdatedBy:     rec.UpdatedBy.String,
			CreatedAt:     rec.CreatedAt.String,
			UpdatedAt:     rec.UpdatedAt.String,
		}

	}

	return resp, nil
}
