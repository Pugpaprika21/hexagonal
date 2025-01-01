package mstparamshdr

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/pkg/sqlx"
	"strings"
)

func (m *mstParamsHdrService) DeleteParamsHdr(ctx context.Context, req request.DeleteParamsHdr) error {
	var sql sqlx.Sqlx

	sql.Stmt = "delete from mst_params_hdr "

	if *req.ParamID != 0 {
		sql.WhereClause = append(sql.WhereClause, "param_id = ?")
		sql.Args = append(sql.Args, req.ParamID)
	}

	if *req.ParamkeyTxt != "" {
		sql.WhereClause = append(sql.WhereClause, "paramkey_txt = ?")
		sql.Args = append(sql.Args, req.ParamkeyTxt)
	}

	if *req.ParamCode != "" {
		sql.WhereClause = append(sql.WhereClause, "param_code = ?")
		sql.Args = append(sql.Args, req.ParamCode)
	}

	if len(sql.WhereClause) > 0 {
		sql.Stmt += " Where " + strings.Join(sql.WhereClause, " and ")
	}

	if err := m.repos.DeleteParamsHdr(ctx, sql); err != nil {
		return err
	}

	return nil
}
