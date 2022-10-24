package cmd

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"go-wms-poc/config"
	"go-wms-poc/database"
	"go-wms-poc/handler"
	"net/http"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

// Validate Data: https://echo.labstack.com/guide/request/#validate-data

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		config := config.Config

		db := database.NewConnection(config)

		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}
		h := handler.New(db)

		route := e.Group("")

		route.GET("areas", h.FindAreas)
		route.GET("areas/:id", h.FindAreaByID)

		route.GET("bins", h.FindBins)
		route.GET("bins/:id", h.FindABinByID)

		route.GET("skus", h.FindSkus)
		route.POST("skus/:id/assign", h.AssignSkuToLocation)

		route.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "wms in go!")
		})

		e.Logger.Fatal(e.Start(":" + config.AppPort))
	},
}
