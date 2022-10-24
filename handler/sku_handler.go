package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go-wms-poc/models"
	"go-wms-poc/response"
)

// todo sku location pivot data in response

func (h *Handler) FindSkus(ctx echo.Context) error {
	var skus []*models.Sku
	p := response.NewPagination(ctx)
	h.db.Scopes(response.Paginate(skus, p, h.db)).Preload("Locations").Find(&skus)

	return ctx.JSON(200, response.NewResponse(p, skus))
}

func (h *Handler) AssignSkuToLocation(ctx echo.Context) error {
	var sku *models.Sku
	var bin *models.Bin

	// first verify the SKU exists
	result := h.db.Debug().Find(&sku, ctx.Param("id"))
	if result.RowsAffected < 1 {
		return ctx.JSON(404, `{"message" : "SKU does not exist"}`)
	}

	input := struct {
		BinId    int `json:"bin_id"`
		Quantity int `json:"quantity"`
	}{BinId: 0, Quantity: 0}

	err := ctx.Bind(&input)
	if err != nil {
		return err
	}

	// then verify the destination BIN exists
	result = h.db.Debug().Find(&bin, input.BinId)
	if result.RowsAffected < 1 {
		return ctx.JSON(400, fmt.Sprintf(`{"message" : "%s"}`, "BIN does not exist"))
	}

	// todo create or update?
	skuLocation := &models.SkuLocation{
		SkuID:    int(sku.ID),
		BinID:    int(sku.ID),
		Quantity: input.Quantity,
	}

	result = h.db.Create(&skuLocation)
	if result.Error != nil {
		return ctx.JSON(400, fmt.Sprintf(`{"message" : "%s"}`, result.Error.Error()))
	}

	return ctx.JSON(200, skuLocation)
}
