package mstparamsdtl

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/internal/core/schema"
	"pugpaprika/app/pkg/sqlx"
	"strings"
)

func (m *mstParamsDtlService) UpdateParamsDtl(ctx context.Context, rec []request.UpdateParamsDtlRows) error {
	var sql sqlx.Sqlx
	var now string = sqlx.Now()
	parmObj := make([]schema.UpdateParamsDtl, len(rec))

	for i, rec := range rec {
		parmObj[i] = schema.UpdateParamsDtl{
			ParamID:       sqlx.Nil(rec.ParamID),
			ParamkeyTxt:   sqlx.Nil(rec.ParamkeyTxt),
			ParamCode:     sqlx.Nil(rec.ParamCode),
			ParamNameth:   sqlx.Nil(rec.ParamNameth),
			ParamNameeng:  sqlx.Nil(rec.ParamNameeng),
			ParamDetail:   sqlx.Nil(rec.ParamDetail),
			ParamOrder:    sqlx.Nil(rec.ParamOrder),
			IsActive:      sqlx.Nil(rec.IsActive),
			CreatedProgby: sqlx.Nil(rec.CreatedProgby),
			CreatedBy:     sqlx.Nil(rec.CreatedBy),
			UpdatedBy:     sqlx.Nil(rec.UpdatedBy),
			CreatedAt:     sqlx.Nil(rec.CreatedAt),
			UpdatedAt:     sqlx.Nil(&now),
		}

		if rec.ParamID != nil && *rec.ParamID != 0 {
			sql.WhereClause = append(sql.WhereClause, "param_id = ?")
			sql.Args = append(sql.Args, rec.ParamID)
		}

		if rec.ParamhdrID != nil && *rec.ParamhdrID != 0 {
			sql.WhereClause = append(sql.WhereClause, "paramhdr_id = ?")
			sql.Args = append(sql.Args, rec.ParamhdrID)
		}

		if rec.ParamkeyID != nil && *rec.ParamkeyID != 0 {
			sql.WhereClause = append(sql.WhereClause, "paramkey_id = ?")
			sql.Args = append(sql.Args, rec.ParamkeyID)
		}

		if rec.ParamparentID != nil && *rec.ParamparentID != 0 {
			sql.WhereClause = append(sql.WhereClause, "paramparent_id = ?")
			sql.Args = append(sql.Args, rec.ParamparentID)
		}

		if rec.ParamkeyTxt != nil && *rec.ParamkeyTxt != "" {
			sql.WhereClause = append(sql.WhereClause, "paramkey_txt = ?")
			sql.Args = append(sql.Args, rec.ParamkeyTxt)
		}

		if rec.ParamCode != nil && *rec.ParamCode != "" {
			sql.WhereClause = append(sql.WhereClause, "param_code = ?")
			sql.Args = append(sql.Args, rec.ParamCode)
		}

		sql.WhereClause = append(sql.WhereClause, "is_active = ?")
		sql.Args = append(sql.Args, true)

		if len(sql.WhereClause) > 0 {
			sql.Stmt += " Where " + strings.Join(sql.WhereClause, " and ")
		}

		if err := m.repos.UpdateParamsDtl(ctx, parmObj, sql); err != nil {
			return err
		}
	}

	return nil
}
