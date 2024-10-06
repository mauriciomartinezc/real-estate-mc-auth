package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-auth/handler"
	"github.com/mauriciomartinezc/real-estate-mc-auth/repository"
	"github.com/mauriciomartinezc/real-estate-mc-auth/seeds/roles"
	"github.com/mauriciomartinezc/real-estate-mc-auth/seeds/users"
	"github.com/mauriciomartinezc/real-estate-mc-auth/service"
	"github.com/mauriciomartinezc/real-estate-mc-common/config"
	"github.com/mauriciomartinezc/real-estate-mc-common/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("application failed: %v", err)
	}
}

func run() error {
	if err := config.LoadEnv(); err != nil {
		return fmt.Errorf("failed to load environment: %w", err)
	}

	if err := config.ValidateEnvironments(); err != nil {
		return fmt.Errorf("invalid environment configuration: %w", err)
	}

	dsn, err := config.GetDSN()
	if err != nil {
		return fmt.Errorf("failed to get DSN: %w", err)
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Error,
			Colorful: true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.AutoMigrate(&domain.Profile{}, &domain.User{}, &domain.Role{}, &domain.Permission{}, &domain.Company{}, &domain.CompanyUser{}); err != nil {
		return fmt.Errorf("failed to auto migrate models: %w", err)
	}

	// Seeds
	roles.SyncRolesSeeds(db)
	users.CreateUserSeeds(db, 0)

	// Repositories and Services
	userRepo := repository.NewUserRepository(db)
	profileRepo := repository.NewProfileRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	permissionRepo := repository.NewPermissionRepository(db)
	companyRepo := repository.NewCompanyRepository(db)
	companyUserRepo := repository.NewCompanyUserRepository(db)

	userService := service.NewUserService(userRepo)
	profileService := service.NewProfileService(profileRepo)
	roleService := service.NewRoleService(roleRepo)
	permissionService := service.NewPermissionService(permissionRepo)
	companyService := service.NewCompanyService(companyRepo)
	companyUserService := service.NewCompanyUserService(companyUserRepo)

	e := echo.New()
	e.Use(middleware.LanguageHandler())

	api := e.Group("/api")
	handler.NewUserHandler(api, userService)
	handler.NewProfileHandler(api, profileService, userService)
	handler.NewRoleHandler(api, roleService)
	handler.NewPermissionHandler(api, permissionService)
	handler.NewCompanyHandler(api, companyService)
	handler.NewCompanyUserHandler(api, companyUserService, userService, companyService)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return e.Start(":" + os.Getenv("SERVER_PORT"))
}
