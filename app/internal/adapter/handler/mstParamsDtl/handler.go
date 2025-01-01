package mstparamsdtl

import (
	mstparamsdtl "pugpaprika/app/internal/core/service/mstParamsDtl"

	"github.com/labstack/echo/v4"
)

type IMstParamsDtlHandler interface {
	GetParamsDtl(c echo.Context) error
	LovParamsDtl(c echo.Context) error
	CreateParamsDtl(c echo.Context) error
	UpdateParamsDtl(c echo.Context) error
	DeleteParamsDtl(c echo.Context) error
}

type mstParamsDtlHandler struct {
	serv mstparamsdtl.IMstParamsDtlService
}

func NewMstParamsDtlHandler(serv mstparamsdtl.IMstParamsDtlService) IMstParamsDtlHandler {
	return &mstParamsDtlHandler{serv: serv}
}
