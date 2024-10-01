package domain

import (
	"github.com/google/uuid"
	commonDomain "github.com/mauriciomartinezc/real-estate-mc-common/domain"
	"gorm.io/gorm"
	"time"
)

type Company struct {
	ID                 uuid.UUID          `json:"id,omitempty" gorm:"type:uuid;primaryKey"`
	Name               string             `json:"name,omitempty" gorm:"not null" validate:"required"`
	ContactEmail       string             `json:"contact_email,omitempty" gorm:"not null" validate:"required"`
	ContactPhone       string             `json:"contact_phone,omitempty" gorm:"not null" validate:"required"`
	BillingAddress     string             `json:"billing_address,omitempty" gorm:"not null" validate:"required"`
	TaxID              string             `json:"tax_id,omitempty" gorm:"not null" validate:"required"`
	CityId             string             `json:"city_id,omitempty" gorm:"type:uuid;not null" validate:"required"`
	City               *commonDomain.City `json:"city,omitempty" gorm:"-"`
	DigitalSignatureID string             `json:"digital_signature_id,omitempty" gorm:"null"`
	CreatedAt          int64              `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt          int64              `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
	CompanyUsers       CompanyUsers       `json:"company_users,omitempty" gorm:"foreignKey:CompanyId"`
}

type Companies []Company

func (c *Company) BeforeCreate(ctx *gorm.DB) (err error) {
	c.ID = uuid.New()
	c.CreatedAt = time.Now().Unix()
	c.UpdatedAt = time.Now().Unix()
	return
}

func (c *Company) BeforeUpdate(ctx *gorm.DB) (err error) {
	c.UpdatedAt = time.Now().Unix()
	return
}
