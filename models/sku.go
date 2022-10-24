package models

import "gorm.io/gorm"

type Sku struct {
	gorm.Model
	Name      string `json:"name"`
	Locations []*Bin `gorm:"many2many:sku_locations;" json:"locations"`
}
