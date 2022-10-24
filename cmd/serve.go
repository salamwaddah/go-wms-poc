package cmd

import (
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

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		config := config.Config

		db := database.NewConnection(config)

		e := echo.New()
		h := handler.New(db)

		route := e.Group("")

		route.GET("areas", h.FindAreas)
		route.GET("areas/:id", h.FindAreaByID)

		route.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "wms in go!")
		})

		e.Logger.Fatal(e.Start(":" + config.AppPort))
	},
}
