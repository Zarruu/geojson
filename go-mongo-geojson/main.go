// main.go
package main

import (
	"log"

	"go-mongo-geojson/config"
	"go-mongo-geojson/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Muat file .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}
    
	// Hubungkan ke database dan simpan client-nya
	mongoClient := config.ConnectDB()

	// PANGGIL FUNGSI INISIALISASI CONTROLLER DI SINI
	controllers.InitPlaceCollection(mongoClient)

	// Inisialisasi router Gin
	router := gin.Default()

	 router.Use(cors.Default())
	 
	// Definisikan rute API
	api := router.Group("/api")
	{
		api.POST("/places", controllers.CreatePlace)
		api.GET("/places", controllers.GetPlaces)
		api.PUT("/places/:id", controllers.UpdatePlace)
		api.DELETE("/places/:id", controllers.DeletePlace)
	}

	// Jalankan server
	router.Run(":8080") // Default port 8080
}