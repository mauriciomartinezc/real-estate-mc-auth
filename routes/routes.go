package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-auth/handlers"
	"github.com/mauriciomartinezc/real-estate-mc-auth/middleware"
	"github.com/mauriciomartinezc/real-estate-mc-auth/repositories"
	"github.com/mauriciomartinezc/real-estate-mc-auth/services"
	"github.com/mauriciomartinezc/real-estate-mc-common/cache"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB, cache cache.Cache) {
	g := e.Group("api")
	company(g, db, cache)
	companyUser(g, db, cache)
	role(g, db, cache)
	permission(g, db, cache)
	profile(g, db, cache)
}

func company(g *echo.Group, db *gorm.DB, cache cache.Cache) {
	repo := repositories.NewCompanyRepository(db, cache)
	service := services.NewCompanyService(repo)
	handler := handlers.NewCompanyHandler(service)

	g.Use(middleware.JWTAuth)

	g.POST("/companies", handler.CreateCompany)
	g.GET("/companies/:uuid", handler.FindCompany)
	g.PUT("/companies/:uuid", handler.UpdateCompany)
	g.GET("/companies/me", handler.CompaniesMe)
}

func companyUser(g *echo.Group, db *gorm.DB, cache cache.Cache) {
	repoUser := repositories.NewUserRepository(db, cache)
	serviceUser := services.NewUserService(repoUser)

	repoProfile := repositories.NewProfileRepository(db, cache)
	serviceProfile := services.NewProfileService(repoProfile)

	repoCompany := repositories.NewCompanyRepository(db, cache)
	serviceCompany := services.NewCompanyService(repoCompany)

	repo := repositories.NewCompanyUserRepository(db, cache)
	service := services.NewCompanyUserService(repo)
	handler := handlers.NewCompanyUserHandler(service, serviceUser, serviceProfile, serviceCompany)

	g.Use(middleware.JWTAuth)

	g.POST("/users", handler.CreateUser)
	g.GET("/companyUsers/:uuid", handler.FindById)
	g.POST("/companyUsers", handler.AddUserToCompany)
	g.PUT("/companyUsers/:uuid", handler.UpdateCompanyUser)
	g.DELETE("/companyUsers/:uuid", handler.DeleteCompanyUser)
}

func role(g *echo.Group, db *gorm.DB, cache cache.Cache) {
	repo := repositories.NewRoleRepository(db, cache)
	service := services.NewRoleService(repo)
	handler := handlers.NewRoleHandler(service)

	g.Use(middleware.JWTAuth)
	g.POST("/roles", handler.CreateRole)
}

func permission(g *echo.Group, db *gorm.DB, cache cache.Cache) {
	repo := repositories.NewPermissionRepository(db, cache)
	service := services.NewPermissionService(repo)
	handler := handlers.NewPermissionHandler(service)

	g.Use(middleware.JWTAuth)

	g.POST("/permissions", handler.CreatePermission)
}

func profile(g *echo.Group, db *gorm.DB, cache cache.Cache) {
	g = g.Group("/auth")

	repoUser := repositories.NewUserRepository(db, cache)
	serviceUser := services.NewUserService(repoUser)

	repo := repositories.NewProfileRepository(db, cache)
	service := services.NewProfileService(repo)
	handler := handlers.NewProfileHandler(service, serviceUser)

	g.Use(middleware.JWTAuth)
	g.POST("/profiles", handler.Create)
	g.GET("/profiles", handler.MeProfile)
	g.PUT("/profiles/:uuid", handler.Update)
}

func user(g *echo.Group, db *gorm.DB, cache cache.Cache) {
	repo := repositories.NewUserRepository(db, cache)
	service := services.NewUserService(repo)
	handler := handlers.NewUserHandler(service)

	groupRoute := g.Group("/auth")
	groupRoute.POST("/register", handler.Register)
	groupRoute.POST("/login", handler.Login)

	g.Use(middleware.JWTAuth)
	g.POST("/resetPassword", handler.ResetPassword)
}
