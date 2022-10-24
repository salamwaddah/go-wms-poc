package handler

import "gorm.io/gorm"

func New(db *gorm.DB) *Handler {
	return &Handler{
		db: db,
	}
}

type Handler struct {
	db *gorm.DB
}
