package bootstrap

import (
	"golang-rest-api/city"
	"golang-rest-api/products"
	"golang-rest-api/role"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, clients *Clients) *gin.Engine {
	router := gin.Default()

	productRepo := products.NewRepository(db)
	productService := products.NewService(productRepo)
	productHandler := products.NewHandler(productService)

	roleRepo := role.NewRepository(db)
	roleService := role.NewService(roleRepo)
	roleHandler := role.NewHandler(roleService)

	cityService := city.NewService(clients.City)
	cityHandler := city.NewHandler(cityService)

	api := router.Group("/api")
	{
		// Products
		api.GET("/products", productHandler.GetProducts)
		api.GET("/products/:id", productHandler.GetProduct)
		api.POST("/products", productHandler.PostProduct)

		// Role
		api.GET("/role", roleHandler.GetRoleByID)
		api.POST("/role", roleHandler.Store)

		// City (External API integration)
		api.GET("/city", cityHandler.GetCities)
		api.GET("/city/:id", cityHandler.GetCities)
	}

	return router
}
