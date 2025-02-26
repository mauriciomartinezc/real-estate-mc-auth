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

// SetupRoutes inicializa las rutas principales
func SetupRoutes(e *echo.Echo, db *gorm.DB, cache cache.Cache) {
	api := e.Group("/api")
	api.Use(middleware.JWTAuth)

	company(api, db, cache)
	companyUser(api, db, cache)
	role(api, db, cache)
	permission(api, db, cache)
	profile(api, db, cache)
	user(api, db, cache) // Usuario no requiere autenticación para registro/login
}

// Rutas de Empresas (Companies)
func company(g *echo.Group, db *gorm.DB, cache cache.Cache) {
	group := g.Group("/companies")

	repo := repositories.NewCompanyRepository(db, cache)
	service := services.NewCompanyService(repo)
	handler := handlers.NewCompanyHandler(service)

	group.POST("", handler.CreateCompany)
	group.GET("/:uuid", handler.FindCompany)
	group.PUT("/:uuid", handler.UpdateCompany)
	group.GET("/me", handler.CompaniesMe)
}

// Rutas de Usuarios en Empresas (Company Users)
func companyUser(g *echo.Group, db *gorm.DB, cache cache.Cache) {
	group := g.Group("/companyUsers")

	repoUser := repositories.NewUserRepository(db, cache)
	serviceUser := services.NewUserService(repoUser)

	repoProfile := repositories.NewProfileRepository(db, cache)
	serviceProfile := services.NewProfileService(repoProfile)

	repoCompany := repositories.NewCompanyRepository(db, cache)
	serviceCompany := services.NewCompanyService(repoCompany)

	repo := repositories.NewCompanyUserRepository(db, cache)
	service := services.NewCompanyUserService(repo)
	handler := handlers.NewCompanyUserHandler(service, serviceUser, serviceProfile, serviceCompany)

	group.POST("/users", handler.CreateUser)
	group.GET("/:uuid", handler.FindById)
	group.POST("", handler.AddUserToCompany)
	group.PUT("/:uuid", handler.UpdateCompanyUser)
	group.DELETE("/:uuid", handler.DeleteCompanyUser)
}

// Rutas de Roles
func role(g *echo.Group, db *gorm.DB, cache cache.Cache) {
	group := g.Group("/roles")

	repo := repositories.NewRoleRepository(db, cache)
	service := services.NewRoleService(repo)
	handler := handlers.NewRoleHandler(service)

	group.POST("", handler.CreateRole)
}

// Rutas de Permisos
func permission(g *echo.Group, db *gorm.DB, cache cache.Cache) {
	group := g.Group("/permissions")

	repo := repositories.NewPermissionRepository(db, cache)
	service := services.NewPermissionService(repo)
	handler := handlers.NewPermissionHandler(service)

	group.POST("", handler.CreatePermission)
}

// Rutas de Perfiles (Profiles)
func profile(g *echo.Group, db *gorm.DB, cache cache.Cache) {
	group := g.Group("/profiles")

	repoUser := repositories.NewUserRepository(db, cache)
	serviceUser := services.NewUserService(repoUser)

	repo := repositories.NewProfileRepository(db, cache)
	service := services.NewProfileService(repo)
	handler := handlers.NewProfileHandler(service, serviceUser)

	group.POST("", handler.Create)
	group.GET("", handler.MeProfile)
	group.PUT("/:uuid", handler.Update)
}

// Rutas de Usuarios y Autenticación
func user(g *echo.Group, db *gorm.DB, cache cache.Cache) {
	auth := g.Group("/auth") // Rutas sin autenticación

	repo := repositories.NewUserRepository(db, cache)
	service := services.NewUserService(repo)
	handler := handlers.NewUserHandler(service)

	auth.POST("/register", handler.Register)
	auth.POST("/login", handler.Login)

	// Rutas protegidas para usuarios autenticados
	protected := g.Group("/user")
	protected.Use(middleware.JWTAuth)
	protected.POST("/resetPassword", handler.ResetPassword)
}
