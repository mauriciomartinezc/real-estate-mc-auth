package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permission struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name string    `json:"name" gorm:"uniqueIndex;not null"`
}

func (p *Permission) BeforeCreate(ctx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}
