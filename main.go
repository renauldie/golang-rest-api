package main

import (
	"log"

	"golang-rest-api/bootstrap"
	"golang-rest-api/internal/config"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	appConfig := config.LoadAppConfig()

	cfg := config.LoadConfig()
	db, err := cfg.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect DB: %v", err)
	}

	//_ = db.AutoMigrate(&products.Product{}, &role.Role{})

	redisConfig := config.LoadRedisConfig()
	redisClient, err := redisConfig.InitRedis()
	if err != nil {
		log.Fatalf("Failed to connect Redis: %v", err)
	}

	defer redisClient.Close()

	clientConfig := config.LoadClientConfig()
	clients := bootstrap.InitClients(clientConfig)

	router := bootstrap.SetupRouter(db, clients)

	log.Printf("Server is running on port %s [%s]", appConfig.Port, appConfig.Env)
	if err := router.Run(appConfig.Address()); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
