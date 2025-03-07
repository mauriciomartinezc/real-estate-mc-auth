package users

import (
	"fmt"
	"github.com/bxcodec/faker/v4"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-auth/repositories"
	"github.com/mauriciomartinezc/real-estate-mc-common/cache"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

func CreateUserSeeds(db *gorm.DB, cache cache.Cache, count int) {
	userRepo := repositories.NewUserRepository(db, cache)
	profileRepo := repositories.NewProfileRepository(db, cache)
	createDefaultUser(userRepo, profileRepo)
	generateUsers(userRepo, profileRepo, count)
}

func createDefaultUser(userRepository repositories.UserRepository, profileRepository repositories.ProfileRepository) {
	userAdmin, _ := userRepository.FindByEmail("super.admin@realestate.com", false)
	if userAdmin == nil {
		roleAdmin := new(domain.Role)
		roleAdmin.ID = domain.ROLES["SUPER_ADMIN"].ID
		firstName := "Super Admin"
		profile := &domain.Profile{
			FirstName: &firstName,
		}
		user := &domain.User{
			Email:    "super.admin@realestate.com",
			Password: "eg9k'_VBnY~VG3ibgnTqn3",
		}
		profile, _ = profileRepository.Create(user, profile)
		profileId := profile.ID.String()
		user.ProfileId = &profileId
		err := userRepository.Create(user)
		if err != nil {
			fmt.Printf("Error when create default user: %s\n", user.Email)
			return
		}
	}
}

func generateUsers(userRepository repositories.UserRepository, profileRepository repositories.ProfileRepository, count int) {
	for i := 0; i < count; i++ {
		roleData := getRandomRole()
		role := new(domain.Role)
		role.ID = roleData.ID
		firstName, lastName := generateFullName()
		email := generateEmail()
		profile := &domain.Profile{
			FirstName: &firstName,
			LastName:  &lastName,
		}
		user := &domain.User{
			Email:    email,
			Password: "Password",
		}
		profile, _ = profileRepository.Create(user, profile)
		profileId := profile.ID.String()
		user.ProfileId = &profileId
		err := userRepository.Create(user)
		if err != nil {
			return
		}
	}
}

func generateFullName() (string, string) {
	firstName := faker.FirstName()
	lastName := faker.LastName()
	return firstName, lastName
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
	return domain.ROLES[randomKey]
}
