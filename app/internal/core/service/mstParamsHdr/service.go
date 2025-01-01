package mstparamshdr

import (
	"context"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/dto/response"
	mstparamshdr "pugpaprika/app/internal/adapter/repository/mstParamsHdr"
)

type IMstParamsHdrService interface {
	GetParamsHdr(ctx context.Context, req request.GetParamsHdr) ([]response.GetParamsHdr, error)
	LovParamsHdr(ctx context.Context, req request.LovParamsHdr) ([]response.LovParamsHdr, error)
	CreateParamsHdr(ctx context.Context, req []request.CreateParamsHdrRows) error
	UpdateParamsHdr(ctx context.Context, req []request.UpdateParamsHdrRows) error
	DeleteParamsHdr(ctx context.Context, req request.DeleteParamsHdr) error
}

type mstParamsHdrService struct {
	repos mstparamshdr.IMstParamsHdrRepository
}

func NewMstParamsHdrService(repos mstparamshdr.IMstParamsHdrRepository) IMstParamsHdrService {
	return &mstParamsHdrService{repos: repos}
}
