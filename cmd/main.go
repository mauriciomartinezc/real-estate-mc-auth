package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-auth/handlers"
	"github.com/mauriciomartinezc/real-estate-mc-auth/routes"
	"github.com/mauriciomartinezc/real-estate-mc-auth/seeds/roles"
	"github.com/mauriciomartinezc/real-estate-mc-auth/seeds/users"
	"github.com/mauriciomartinezc/real-estate-mc-common/config"
	"github.com/mauriciomartinezc/real-estate-mc-common/discovery/consul"
	"github.com/mauriciomartinezc/real-estate-mc-common/middlewares"
	"github.com/mauriciomartinezc/real-estate-mc-common/utils"
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

	cacheClient := config.NewCacheClient()

	// Consul
	discoveryClient := consul.NewConsultApi()
	discoveryClient.RegisterService("mc-auth")

	// Seeds
	roles.SyncRolesSeeds(db, cacheClient)
	users.CreateUserSeeds(db, cacheClient, 0)

	e := echo.New()
	utils.RouteHealth(e)
	e.Use(middlewares.LanguageHandler())
	handlers.InitValidate()
	routes.SetupRoutes(e, db, cacheClient)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	log.Println("Registered routes:")
	for _, r := range e.Routes() {
		// Filtrar la ruta no deseada
		if r.Method == "echo_route_not_found" {
			continue
		}
		fmt.Printf("%s %s\n", r.Method, r.Path)
	}

	return e.Start(":" + os.Getenv("SERVER_PORT"))
}
