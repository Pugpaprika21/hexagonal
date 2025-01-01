package mstparamshdr

import (
	mstparamshdr "pugpaprika/app/internal/core/service/mstParamsHdr"

	"github.com/labstack/echo/v4"
)

type IMstParamsHdrHandler interface {
	GetParamsHdr(c echo.Context) error
	LovParamsHdr(c echo.Context) error
	CreateParamsHdr(c echo.Context) error
	UpdateParamsHdr(c echo.Context) error
	DeleteParamsHdr(c echo.Context) error
}

type mstParamsHdrHandler struct {
	serv mstparamshdr.IMstParamsHdrService
}

func NewMstParamsHdrHandler(serv mstparamshdr.IMstParamsHdrService) IMstParamsHdrHandler {
	return &mstParamsHdrHandler{serv: serv}
}
