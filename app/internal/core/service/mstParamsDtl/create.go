package mstparamsdtl

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/internal/core/schema"
	"pugpaprika/app/pkg/sqlx"
)

func (m *mstParamsDtlService) CreateParamsDtl(ctx context.Context, req []request.CreateParamsDtlRows) error {
	parmObj := make([]schema.CreateParamsDtl, len(req))
	for i, rec := range req {
		parmObj[i] = schema.CreateParamsDtl{
			ParamID:       sqlx.Nil(rec.ParamID),
			ParamkeyID:    sqlx.Nil(rec.ParamkeyID),
			ParamparentID: sqlx.Nil(rec.ParamparentID),
			ParamkeyTxt:   sqlx.Nil(rec.ParamkeyTxt),
			ParamCode:     sqlx.Nil(rec.ParamCode),
			ParamNameth:   sqlx.Nil(rec.ParamNameth),
			ParamNameeng:  sqlx.Nil(rec.ParamNameeng),
			ParamDetail:   sqlx.Nil(rec.ParamDetail),
			ParamOrder:    sqlx.Nil(rec.ParamOrder),
			ParamhdrID:    sqlx.Nil(rec.ParamhdrID),
			IsActive:      sqlx.Nil(rec.IsActive),
			CreatedProgby: sqlx.Nil(rec.CreatedProgby),
			CreatedBy:     sqlx.Nil(rec.CreatedBy),
			UpdatedBy:     sqlx.Nil(rec.UpdatedBy),
			CreatedAt:     sqlx.Nil(rec.CreatedAt),
			UpdatedAt:     sqlx.Nil(rec.UpdatedAt),
		}
	}

	if err := m.repos.CreateParamsDtl(ctx, parmObj); err != nil {
		return err
	}

	return nil
}
