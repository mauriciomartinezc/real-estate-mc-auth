package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permission struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name string    `json:"name" gorm:"unique;index;not null"`
	Slug string    `json:"slug" gorm:"unique;index;not null"`
}

func (p *Permission) BeforeCreate(ctx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}
