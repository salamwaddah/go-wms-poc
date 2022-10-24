package handler

import (
	"github.com/labstack/echo/v4"
	"go-wms-poc/models"
	"go-wms-poc/response"
)

func (h *Handler) FindBins(ctx echo.Context) error {
	var bins []*models.Bin
	p := response.NewPagination(ctx)
	h.db.Scopes(response.Paginate(bins, p, h.db)).Preload("Area").Find(&bins)

	return ctx.JSON(200, response.NewResponse(p, bins))
}

func (h *Handler) FindABinByID(ctx echo.Context) error {
	var bin *models.Bin
	h.db.Preload("Area").Find(&bin, ctx.Param("id"))

	return ctx.JSON(200, bin)
}
