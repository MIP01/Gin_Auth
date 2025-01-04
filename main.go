package main

import (
    "log"
	"go_auth/route"
	"go_auth/config"

	"github.com/gin-gonic/gin"
)


func main() {
    // Inisialisasi konfigurasi dan koneksi database
    config.InitConfig()

    // Inisialisasi router
    r := gin.Default()

    // Set up routes
    api := r.Group("/api/v1")
    route.SetupRoutes(api)

    // Informasi URL API
	log.Println("Server is running at http://localhost:8080")

    // Jalankan server pada port 8080
    r.Run(":8080")
}