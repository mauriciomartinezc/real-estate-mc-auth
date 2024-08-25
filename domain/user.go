package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primaryKey"`
	Email     string    `json:"email,omitempty" gorm:"unique;index;not null"`
	Password  string    `json:"password,omitempty" gorm:"not null"`
	ProfileId uuid.UUID `json:"profile_id,omitempty" gorm:"type:uuid;default:null"`
	Profile   Profile   `json:"profile" gorm:"foreign:ProfileId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Roles     []Role    `json:"roles,omitempty" gorm:"many2many:user_roles;"`
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
