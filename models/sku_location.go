package models

import "time"

type SkuLocation struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	BinID     int       `gorm:"uniqueIndex:idx_bin_sku;" json:"bin_id"`
	SkuID     int       `gorm:"uniqueIndex:idx_bin_sku;" json:"sku_id"`
	Quantity  int       `json:"quantity"`

	Sku Sku `json:"sku,omitempty"`
	Bin Bin `json:"bin,omitempty"`
}
