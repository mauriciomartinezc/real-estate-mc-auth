package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type CompanyUser struct {
	ID        uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primaryKey"`
	UserId    string    `json:"user_id,omitempty" gorm:"type:uuid;not null" validate:"required"`
	User      *User     `json:"user,omitempty" gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CompanyId string    `json:"company_id,omitempty" gorm:"type:uuid;not null"`
	Company   *Company  `json:"company,omitempty" gorm:"foreignKey:CompanyId"`
	Roles     *Roles    `json:"roles,omitempty" gorm:"many2many:company_user_roles;"`
	CreatorId string    `json:"creator_id,omitempty" gorm:"type:uuid;not null"`
	Creator   *User     `json:"creator,omitempty" gorm:"foreignKey:CreatorId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UpdaterId *string   `json:"updater_id,omitempty" gorm:"type:uuid"`
	Updater   *User     `json:"updater,omitempty" gorm:"foreignKey:UpdaterId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt int64     `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt int64     `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}

type CompanyUsers []CompanyUser

func (c *CompanyUser) BeforeCreate(ctx *gorm.DB) (err error) {
	c.ID = uuid.New()
	c.CreatedAt = time.Now().Unix()
	c.UpdatedAt = time.Now().Unix()
	return
}

func (c *CompanyUser) BeforeUpdate(ctx *gorm.DB) (err error) {
	c.UpdatedAt = time.Now().Unix()
	return
}
