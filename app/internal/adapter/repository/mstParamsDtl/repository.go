package mstparamsdtl

import (
	"context"
	"pugpaprika/app/internal/core/schema"
	"pugpaprika/app/pkg/sqlx"

	"gorm.io/gorm"
)

type IMstParamsDtlRepository interface {
	GetParamsDtl(ctx context.Context, sql sqlx.Sqlx) ([]schema.GetParamsDtl, error)
	LovParamsDtl(ctx context.Context, sql sqlx.Sqlx) ([]schema.LovParamsDtl, error)
	CreateParamsDtl(ctx context.Context, parmObj []schema.CreateParamsDtl) error
	UpdateParamsDtl(ctx context.Context, parmObj []schema.UpdateParamsDtl, sql sqlx.Sqlx) error
	DeleteParamsDtl(ctx context.Context, sql sqlx.Sqlx) error
}

type mstParamsDtlRepository struct {
	db *gorm.DB
}

func NewMstParamsDtlRepository(db *gorm.DB) IMstParamsDtlRepository {
	return &mstParamsDtlRepository{db: db}
}
