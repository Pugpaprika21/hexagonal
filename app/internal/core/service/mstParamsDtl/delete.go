package mstparamsdtl

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/pkg/sqlx"
	"strings"
)

func (m *mstParamsDtlService) DeleteParamsDtl(ctx context.Context, req request.DeleteParamsDtl) error {
	var sql sqlx.Sqlx

	sql.Stmt = "delete from mst_params_dtl "

	if *req.ParamID != 0 {
		sql.WhereClause = append(sql.WhereClause, "param_id = ?")
		sql.Args = append(sql.Args, req.ParamID)
	}

	if *req.ParamhdrID != 0 {
		sql.WhereClause = append(sql.WhereClause, "paramhdr_id = ?")
		sql.Args = append(sql.Args, req.ParamhdrID)
	}

	if *req.ParamkeyID != 0 {
		sql.WhereClause = append(sql.WhereClause, "paramkey_id = ?")
		sql.Args = append(sql.Args, req.ParamkeyID)
	}

	if *req.ParamparentID != 0 {
		sql.WhereClause = append(sql.WhereClause, "paramparent_id = ?")
		sql.Args = append(sql.Args, req.ParamparentID)
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

	if err := m.repos.DeleteParamsDtl(ctx, sql); err != nil {
		return err
	}

	return nil
}
