package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permission struct {
	ID   uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primaryKey"`
	Name string    `json:"name,omitempty" gorm:"unique;index;not null"`
	Slug string    `json:"slug,omitempty" gorm:"unique;index;not null"`
}

func (p *Permission) BeforeCreate(ctx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}
