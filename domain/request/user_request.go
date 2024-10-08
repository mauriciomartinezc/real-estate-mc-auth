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

type ChangePasswordRequest struct {
	UserId               string `json:"user_id" validate:"required"`
	OldPassword          string `json:"old_password" validate:"required"`
	NewPassword          string `json:"new_password" validate:"required"`
	ConfirmedNewPassword string `json:"confirmed_new_password" validate:"required"`
}
