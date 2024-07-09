package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID          uuid.UUID    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name        string       `json:"name" gorm:"uniqueIndex;not null"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}

func (r *Role) BeforeCreate(ctx *gorm.DB) (err error) {
	r.ID = uuid.New()
	return
}
