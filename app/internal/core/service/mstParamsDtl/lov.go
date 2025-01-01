package mstparamsdtl

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/dto/response"
	"pugpaprika/app/pkg/sqlx"
	"strings"
)

func (m *mstParamsDtlService) LovParamsDtl(ctx context.Context, req request.LovParamsDtl) ([]response.LovParamsDtl, error) {
	var sql sqlx.Sqlx
	var resp []response.LovParamsDtl

	sql.Stmt = `select param_id, paramkey_id, paramparent_id, paramkey_txt, param_code, param_nameth, param_nameeng, param_detail, param_order from mst_params_dtl `

	if req.ParamID != nil && *req.ParamID != 0 {
		sql.WhereClause = append(sql.WhereClause, "param_id = ?")
		sql.Args = append(sql.Args, req.ParamID)
	}

	if req.ParamhdrID != nil && *req.ParamhdrID != 0 {
		sql.WhereClause = append(sql.WhereClause, "paramhdr_id = ?")
		sql.Args = append(sql.Args, req.ParamhdrID)
	}

	if req.ParamkeyID != nil && *req.ParamkeyID != 0 {
		sql.WhereClause = append(sql.WhereClause, "paramkey_id = ?")
		sql.Args = append(sql.Args, req.ParamkeyID)
	}

	if req.ParamparentID != nil && *req.ParamparentID != 0 {
		sql.WhereClause = append(sql.WhereClause, "paramparent_id = ?")
		sql.Args = append(sql.Args, req.ParamparentID)
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

	rows, err := m.repos.LovParamsDtl(ctx, sql)
	if err != nil {
		return nil, err
	}

	if len(rows) > 0 {
		for _, rec := range rows {
			data := response.LovParamsDtl{
				ParamID:       rec.ParamID.Int64,
				ParamkeyID:    rec.ParamkeyID.Int32,
				ParamparentID: rec.ParamparentID.Int32,
				ParamkeyTxt:   rec.ParamkeyTxt.String,
				ParamCode:     rec.ParamCode.String,
				ParamNameth:   rec.ParamNameth.String,
				ParamNameeng:  rec.ParamNameeng.String,
				ParamDetail:   rec.ParamDetail.String,
				ParamOrder:    rec.ParamOrder.Int32,
				ParamhdrID:    rec.ParamhdrID.Int32,
				IsActive:      rec.IsActive.Bool,
				CreatedProgby: rec.CreatedProgby.String,
				CreatedBy:     rec.CreatedBy.String,
				UpdatedBy:     rec.UpdatedBy.String,
				CreatedAt:     rec.CreatedAt.String,
				UpdatedAt:     rec.UpdatedAt.String,
			}
			resp = append(resp, data)
		}
	}

	return resp, nil
}
