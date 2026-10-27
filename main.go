package main

import (
	"log"

	"golang-rest-api/bootstrap"
	"golang-rest-api/config"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
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

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
