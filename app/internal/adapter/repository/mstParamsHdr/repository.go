package mstparamshdr

import (
	"context"
	"pugpaprika/app/internal/core/schema"
	"pugpaprika/app/pkg/sqlx"

	"gorm.io/gorm"
)

type IMstParamsHdrRepository interface {
	GetParamsHdr(ctx context.Context, sql sqlx.Sqlx) ([]schema.GetParamsHdr, error)
	LovParamsHdr(ctx context.Context, sql sqlx.Sqlx) ([]schema.LovParamsHdr, error)
	CreateParamsHdr(ctx context.Context, parmObj []schema.CreateParamsHdr) error
	UpdateParamsHdr(ctx context.Context, parmObj []schema.UpdateParamsHdr, sql sqlx.Sqlx) error
	DeleteParamsHdr(ctx context.Context, sql sqlx.Sqlx) error
}

type mstParamsHdrRepository struct {
	db *gorm.DB
}

func NewMstParamsHdrRepository(db *gorm.DB) IMstParamsHdrRepository {
	return &mstParamsHdrRepository{db: db}
}
