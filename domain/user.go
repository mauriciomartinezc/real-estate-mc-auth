package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID                    uuid.UUID     `json:"id,omitempty" gorm:"type:uuid;primaryKey"`
	Email                 string        `json:"email,omitempty" gorm:"unique;index;not null" validate:"required,email"`
	Password              string        `json:"password,omitempty" gorm:"not null" validate:"required"`
	ProfileId             *string       `json:"profile_id,omitempty" gorm:"type:uuid;default:null"`
	Profile               *Profile      `json:"profile" gorm:"foreign:ProfileId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CompanyUsers          *CompanyUsers `json:"company_users" gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreateForUser         bool          `json:"create_for_user" gorm:"type:bool;default:false"`
	RequiredResetPassword bool          `json:"required_reset_password" gorm:"type:bool;default:false"`
	CreatedAt             int64
	UpdatedAt             int64
}

type Users []User

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
