package users

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-auth/repository"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

func CreateUserSeeds(db *gorm.DB, count int) {
	userRepo := repository.NewUserRepository(db)
	createDefaultUser(db, userRepo)
	generateUsers(userRepo, db, count)
}

func createDefaultUser(db *gorm.DB, userRepository repository.UserRepository) {
	userAdmin, _ := userRepository.FindByEmail("super.admin@realestate.com", false)
	if userAdmin == nil {
		roleRepo := repository.NewRoleRepository(db)
		roleAdmin, _ := roleRepo.FindBySlug(domain.ROLES["SUPER_ADMIN"].Slug)
		fmt.Println(roleAdmin)
		user := &domain.User{
			Name:     "Super Admin",
			Email:    "super.admin@realestate.com",
			Password: "eg9k'_VBnY~VG3ibgnTqn3",
			Roles:    []domain.Role{*roleAdmin},
		}
		err := userRepository.Create(user)
		if err != nil {
			fmt.Printf("Error when create default user: %s\n", user.Email)
			return
		}
	}
}

func generateUsers(userRepository repository.UserRepository, db *gorm.DB, count int) {
	roleRepo := repository.NewRoleRepository(db)
	for i := 0; i < count; i++ {
		randomRole := getRandomRole()
		role, _ := roleRepo.FindBySlug(randomRole.Slug)
		name := generateFullName()
		email := generateEmail()
		fmt.Println(name)
		user := &domain.User{
			Name:     name,
			Email:    email,
			Password: "Password",
			Roles:    []domain.Role{*role},
		}
		err := userRepository.Create(user)
		if err != nil {
			return
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

func getRandomRole() domain.Role {
	rand.Seed(time.Now().UnixNano())
	keys := make([]string, 0, len(domain.ROLES))
	for k := range domain.ROLES {
		keys = append(keys, k)
	}
	randomKey := keys[rand.Intn(len(keys))]
	fmt.Println(randomKey)
	return domain.ROLES[randomKey]
}
