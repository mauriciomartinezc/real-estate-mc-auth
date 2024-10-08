package request

import "github.com/mauriciomartinezc/real-estate-mc-auth/domain"

type CompanyUserRequest struct {
	CompanyId string       `json:"company_id" validate:"required"`
	UserId    string       `json:"user_id" validate:"required"`
	Roles     domain.Roles `json:"roles" validate:"required"`
}

type CompanyUserResponse struct {
}
