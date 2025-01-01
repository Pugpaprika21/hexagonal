package mstparamshdr

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/internal/core/schema"
	"pugpaprika/app/pkg/sqlx"
)

func (m *mstParamsHdrService) CreateParamsHdr(ctx context.Context, req []request.CreateParamsHdrRows) error {
	parmObj := make([]schema.CreateParamsHdr, len(req))
	for i, rec := range req {
		parmObj[i] = schema.CreateParamsHdr{
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
			UpdatedAt:     sqlx.Nil(rec.UpdatedAt),
		}
	}

	if err := m.repos.CreateParamsHdr(ctx, parmObj); err != nil {
		return err
	}

	return nil
}
