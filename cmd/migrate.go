package cmd

import (
	"github.com/spf13/cobra"
	"go-wms-poc/config"
	"go-wms-poc/database"
	"go-wms-poc/models"
)

func init() {
	rootCmd.AddCommand(migrate)
}

var migrate = &cobra.Command{
	Use: "migrate",
	Run: func(cmd *cobra.Command, args []string) {
		db := database.NewConnection(config.Config)

		migrationErr := db.AutoMigrate(
			&models.Area{},
			&models.Bin{},
			&models.Sku{},
			&models.SkuLocation{},
		)

		if migrationErr != nil {
			panic("error executing migration script")
		}
	},
}
