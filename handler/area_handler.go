package handler

import (
	"github.com/labstack/echo/v4"
	"go-wms-poc/models"
	"go-wms-poc/response"
)

func (h *Handler) FindAreas(ctx echo.Context) error {
	var areas []*models.Area
	p := response.NewPagination(ctx)
	h.db.Scopes(response.Paginate(areas, p, h.db)).Find(&areas)

	return ctx.JSON(200, response.NewResponse(p, areas))
}

func (h *Handler) FindAreaByID(ctx echo.Context) error {
	var area *models.Area
	h.db.Find(&area, ctx.Param("id"))

	return ctx.JSON(200, area)
}
