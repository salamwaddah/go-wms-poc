package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go-wms-poc/models"
	"go-wms-poc/response"
)

func (h *Handler) FindSkus(ctx echo.Context) error {
	var skus []*models.Sku
	p := response.NewPagination(ctx)
	h.db.Scopes(response.Paginate(skus, p, h.db)).Preload("Locations").Find(&skus)

	return ctx.JSON(200, response.NewResponse(p, skus))
}

type AssignBody struct {
	BinId    int `json:"bin_id" validate:"required,numeric"`
	Quantity int `json:"quantity" validate:"required,numeric"`
}

func (h *Handler) AssignSkuToLocation(ctx echo.Context) (err error) {
	var sku *models.Sku
	var bin *models.Bin

	reqBody := new(AssignBody)

	if bindErr := ctx.Bind(&reqBody); bindErr != nil {
		return bindErr
	}

	if err = ctx.Validate(reqBody); err != nil {
		return err
	}

	// first verify the SKU exists
	result := h.db.Debug().Find(&sku, ctx.Param("id"))
	if result.RowsAffected < 1 {
		return ctx.JSON(404, `{"message" : "SKU does not exist"}`)
	}

	// then verify the destination BIN exists
	result = h.db.Debug().Find(&bin, reqBody.BinId)
	if result.RowsAffected < 1 {
		return ctx.JSON(400, fmt.Sprintf(`{"message" : "%s"}`, "BIN does not exist"))
	}

	skuLocation := &models.SkuLocation{
		SkuID:    int(sku.ID),
		BinID:    int(sku.ID),
		Quantity: reqBody.Quantity,
	}

	result = h.db.Create(&skuLocation)
	if result.Error != nil {
		return ctx.JSON(400, fmt.Sprintf(`{"message" : "%s"}`, result.Error.Error()))
	}

	return ctx.JSON(200, skuLocation)
}
