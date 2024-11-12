package roles

import (
	"fmt"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-auth/repositories"
	"github.com/mauriciomartinezc/real-estate-mc-common/cache"
	"gorm.io/gorm"
)

func SyncRolesSeeds(db *gorm.DB, cache cache.Cache) {
	roleRepo := repositories.NewRoleRepository(db, cache)
	for _, role := range domain.ROLES {
		_, err := roleRepo.FindBySlug(role.Slug)
		if err != nil {
			errCreate := roleRepo.Create(&role)
			if errCreate != nil {
				fmt.Printf("Error al crear el rol %s: %v\n", role.Name, err)
			}
		}
	}
}
