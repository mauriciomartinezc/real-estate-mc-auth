package request

import "github.com/mauriciomartinezc/real-estate-mc-auth/domain"

type UserRequest struct {
	Email     string       `json:"email" validate:"required,email"`
	Password  string       `json:"password" validate:"required"`
	FirstName string       `json:"first_name" validate:"required"`
	LastName  string       `json:"last_name" validate:"required"`
	CityId    string       `json:"city_id" validate:"required"`
	Roles     domain.Roles `json:"roles" validate:"required"`
	CompanyId string       `json:"company_id" validate:"required"`
}
