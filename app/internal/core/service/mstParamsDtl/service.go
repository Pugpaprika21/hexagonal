package mstparamsdtl

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/dto/response"
	mstparamsdtl "pugpaprika/app/internal/adapter/repository/mstParamsDtl"
)

type IMstParamsDtlService interface {
	GetParamsDtl(ctx context.Context, req request.GetParamsDtl) ([]response.GetParamsDtl, error)
	LovParamsDtl(ctx context.Context, req request.LovParamsDtl) ([]response.LovParamsDtl, error)
	CreateParamsDtl(ctx context.Context, req []request.CreateParamsDtlRows) error
	UpdateParamsDtl(ctx context.Context, req []request.UpdateParamsDtlRows) error
	DeleteParamsDtl(ctx context.Context, req request.DeleteParamsDtl) error
}

type mstParamsDtlService struct {
	repos mstparamsdtl.IMstParamsDtlRepository
}

func NewMstParamsDtlService(repos mstparamsdtl.IMstParamsDtlRepository) IMstParamsDtlService {
	return &mstParamsDtlService{repos: repos}
}
