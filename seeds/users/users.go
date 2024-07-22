package users

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-auth/repository"
	"gorm.io/gorm"
)

func CreateUserSeeds(db *gorm.DB, count int) {
	userRepo := repository.NewUserRepository(db)
	userAdmin, _ := userRepo.FindByEmail("super.admin@realestate.com", false)
	if userAdmin == nil {
		createDefaultUser(db)
		for i := 0; i < count; i++ {
			name := generateFullName()
			email := generateEmail()
			fmt.Println(name)
			user := &domain.User{
				Name:     name,
				Email:    email,
				Password: "Password",
			}
			err := userRepo.Create(user)
			if err != nil {
				fmt.Printf("Error when create user: %s\n", user.Email)
				return
			}
		}
	}

}

func generateFullName() string {
	firstName := faker.FirstName()
	lastName := faker.LastName()
	return firstName + " " + lastName
}

func generateEmail() string {
	return faker.Email()
}

func createDefaultUser(db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	user := &domain.User{
		Name:     "Super Admin",
		Email:    "super.admin@realestate.com",
		Password: "eg9k'_VBnY~VG3ibgnTqn3",
	}
	err := userRepo.Create(user)
	if err != nil {
		fmt.Printf("Error when create default user: %s\n", user.Email)
		return
	}
}
