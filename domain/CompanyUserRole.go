package domain

type CompanyUserRole struct {
	CompanyUserId string      `json:"company_user_id" gorm:"type:uuid;not null"`
	RoleId        string      `json:"role_id" gorm:"type:uuid;not null"`
	CompanyUser   CompanyUser `json:"company_user" gorm:"foreignKey:CompanyUserID"`
	Role          Role        `json:"role" gorm:"foreignKey:RoleId"`
}

type CompanyUserRoles []CompanyUserRole
