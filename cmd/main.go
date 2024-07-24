package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-auth/config"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-auth/handler"
	"github.com/mauriciomartinezc/real-estate-mc-auth/middleware"
	"github.com/mauriciomartinezc/real-estate-mc-auth/repository"
	"github.com/mauriciomartinezc/real-estate-mc-auth/seeds/roles"
	"github.com/mauriciomartinezc/real-estate-mc-auth/seeds/users"
	"github.com/mauriciomartinezc/real-estate-mc-auth/service"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

func main() {
	err := config.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	dsn := config.GetDSN()

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Error,
			Colorful: true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&domain.User{}, &domain.Role{}, &domain.Permission{})
	if err != nil {
		log.Fatalf("failed to auto migrate models: %v", err)
	}

	roles.SyncRolesSeeds(db)
	users.CreateUserSeeds(db, 0)

	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	permissionRepo := repository.NewPermissionRepository(db)

	userService := service.NewUserService(userRepo, roleRepo, permissionRepo)
	roleService := service.NewRoleService(roleRepo)
	permissionService := service.NewPermissionService(permissionRepo)

	e := echo.New()

	e.Use(middleware.LanguageHandler())

	api := e.Group("/api")

	handler.NewUserHandler(api, userService)
	handler.NewRoleHandler(api, roleService)
	handler.NewPermissionHandler(api, permissionService)

	// swagger documentation
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	log.Fatal(e.Start(":" + os.Getenv("SERVER_PORT")))
}
