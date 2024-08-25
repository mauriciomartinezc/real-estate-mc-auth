package domain

import (
	"github.com/google/uuid"
)

type Role struct {
	ID          uuid.UUID    `json:"id,omitempty" gorm:"type:uuid;primaryKey"`
	Name        string       `json:"name,omitempty" gorm:"unique;index;not null"`
	Slug        string       `json:"slug,omitempty" gorm:"unique;index;not null"`
	Permissions []Permission `json:"permissions,omitempty" gorm:"many2many:role_permissions;"`
}

var ROLES = map[string]Role{
	"SUPER_ADMIN": {
		ID:   uuid.MustParse("dc56c7fd-e4f4-43ce-bb2b-e6c12e95495d"),
		Name: "Super Admin",
		Slug: "super_admin",
		Permissions: []Permission{
			{Name: "Super Admin", Slug: "super_admin"},
		},
	},
	"ADMIN": {
		ID:   uuid.MustParse("4b919037-45c1-4856-bf7f-85548b923243"),
		Name: "Administrador de Empresa",
		Slug: "admin",
		Permissions: []Permission{
			{Name: "Administrador de Empresa", Slug: "admin"},
		},
	},
	"FINANCE": {
		ID:   uuid.MustParse("60b98916-9bd9-43e2-bc27-f6e9ce4e1948"),
		Name: "Financiero",
		Slug: "finance",
		Permissions: []Permission{
			{Name: "Financiero", Slug: "finance"},
		},
	},
	"AGENT": {
		ID:   uuid.MustParse("f63f651c-d32b-46e4-8dfc-68f4d5b4e199"),
		Name: "Agente Inmobiliario",
		Slug: "agent",
		Permissions: []Permission{
			{Name: "Agente Inmobiliario", Slug: "agent"},
		},
	},
	"BROKER": {
		ID:   uuid.MustParse("c87cf265-5703-437b-9e4f-1eda20b9d846"),
		Name: "Broker",
		Slug: "broker",
		Permissions: []Permission{
			{Name: "Broker", Slug: "broker"},
		},
	},
	"USER": {
		ID:   uuid.MustParse("436871f2-f6a7-4920-9b4f-f2b65c425ee8"),
		Name: "Usuario",
		Slug: "user",
		Permissions: []Permission{
			{Name: "Usuario", Slug: "user"},
		},
	},
}
