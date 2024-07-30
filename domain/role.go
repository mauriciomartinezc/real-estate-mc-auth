package domain

import (
	"github.com/google/uuid"
)

type Role struct {
	ID          uuid.UUID    `gorm:"type:uuid;primaryKey"`
	Name        string       `json:"name" gorm:"unique;index;not null"`
	Slug        string       `json:"slug" gorm:"unique;index;not null"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}

var ROLES = map[string]Role{
	"SUPER_ADMIN": {
		Name: "Super Admin",
		Slug: "super_admin",
		Permissions: []Permission{
			{Name: "Super Admin", Slug: "super_admin"},
		},
	},
	"ADMIN": {
		Name: "Administrador de Empresa",
		Slug: "admin",
		Permissions: []Permission{
			{Name: "Administrador de Empresa", Slug: "admin"},
		},
	},
	"FINANCE": {
		Name: "Financiero",
		Slug: "finance",
		Permissions: []Permission{
			{Name: "Financiero", Slug: "finance"},
		},
	},
	"AGENT": {
		Name: "Agente Inmobiliario",
		Slug: "agent",
		Permissions: []Permission{
			{Name: "Agente Inmobiliario", Slug: "agent"},
		},
	},
	"BROKER": {
		Name: "Broker",
		Slug: "broker",
		Permissions: []Permission{
			{Name: "Broker", Slug: "broker"},
		},
	},
	"USER": {
		Name: "Usuario",
		Slug: "user",
		Permissions: []Permission{
			{Name: "Usuario", Slug: "user"},
		},
	},
}
