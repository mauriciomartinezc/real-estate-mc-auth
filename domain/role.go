package domain

import (
	"github.com/google/uuid"
)

type Role struct {
	ID          uuid.UUID    `json:"id,omitempty" gorm:"type:uuid;primaryKey"`
	Name        string       `json:"name,omitempty" gorm:"unique;index;not null" validate:"required"`
	Slug        string       `json:"slug,omitempty" gorm:"unique;index;not null" validate:"required"`
	Permissions []Permission `json:"permissions,omitempty" gorm:"many2many:role_permissions;"`
}

type Roles []Role

var ROLES = map[string]Role{
	"SUPER_ADMIN": {
		ID:   uuid.MustParse("806cac72-30b0-4c56-a511-28a7d2077f94"),
		Name: "Super Admin",
		Slug: "super_admin",
		Permissions: []Permission{
			{Name: "Super Admin", Slug: "super_admin"},
		},
	},
	"ADMIN": {
		ID:   uuid.MustParse("ae3c9743-5745-488e-bbdd-3d1cb69db1eb"),
		Name: "Administrador de Empresa",
		Slug: "admin",
		Permissions: []Permission{
			{Name: "Administrador de Empresa", Slug: "admin"},
		},
	},
	"FINANCE": {
		ID:   uuid.MustParse("7e51c41f-a4be-4339-8dc1-6a0f00c91709"),
		Name: "Financiero",
		Slug: "finance",
		Permissions: []Permission{
			{Name: "Financiero", Slug: "finance"},
		},
	},
	"AGENT": {
		ID:   uuid.MustParse("04dec8e2-c013-41c2-8ed2-9b281fb07692"),
		Name: "Agente Inmobiliario",
		Slug: "agent",
		Permissions: []Permission{
			{Name: "Agente Inmobiliario", Slug: "agent"},
		},
	},
	"BROKER": {
		ID:   uuid.MustParse("530b7eaa-e7ee-4da0-af46-6fbc45096069"),
		Name: "Broker",
		Slug: "broker",
		Permissions: []Permission{
			{Name: "Broker", Slug: "broker"},
		},
	},
	"USER": {
		ID:   uuid.MustParse("033bbdf5-b041-4ec5-9011-8f1314de74e5"),
		Name: "Usuario",
		Slug: "user",
		Permissions: []Permission{
			{Name: "Usuario", Slug: "user"},
		},
	},
}
