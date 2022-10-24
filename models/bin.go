package models

import "gorm.io/gorm"

type Bin struct {
	gorm.Model
	Name   string `json:"name"`
	AreaId int
	Area   Area
	Skus   []*Sku `gorm:"many2many:sku_locations;" json:"skus,omitempty"`
}
