package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique;index;not null"`
	Password  string    `json:"password" gorm:"not null"`
	Roles     []Role    `gorm:"many2many:user_roles;"`
	CreatedAt int64
	UpdatedAt int64
}

func (u *User) BeforeCreate(ctx *gorm.DB) (err error) {
	u.ID = uuid.New()
	u.CreatedAt = time.Now().Unix()
	u.UpdatedAt = time.Now().Unix()
	return
}

func (u *User) BeforeUpdate(ctx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now().Unix()
	return
}
