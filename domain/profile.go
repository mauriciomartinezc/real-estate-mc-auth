package domain

import (
	"github.com/google/uuid"
	commonDomain "github.com/mauriciomartinezc/real-estate-mc-common/domain"
	"gorm.io/gorm"
	"time"
)

type Profile struct {
	ID        uuid.UUID          `json:"id,omitempty" gorm:"type:uuid;primary_key"`
	FirstName *string            `json:"first_name,omitempty" gorm:"type:varchar(255)" validate:"required"`
	LastName  *string            `json:"last_name,omitempty" gorm:"type:varchar(255)" validate:"required"`
	CityId    *string            `json:"city_id,omitempty" gorm:"type:uuid;default:null" validate:"required"`
	City      *commonDomain.City `json:"city,omitempty" gorm:"-"`
	CreatedAt int64
	UpdatedAt int64
}

func (p *Profile) BeforeCreate(ctx *gorm.DB) (err error) {
	p.ID = uuid.New()
	p.CreatedAt = time.Now().Unix()
	p.UpdatedAt = time.Now().Unix()
	return
}

func (p *Profile) BeforeUpdate(ctx *gorm.DB) (err error) {
	p.UpdatedAt = time.Now().Unix()
	return
}
