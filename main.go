package main

import (
	"log"
	"net/http"
	"os"

	"football-app/src/routes"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Cargar variables desde .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌❌❌ Error loading .env file ❌❌❌" + "," + "Reason:" + err.Error())
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌❌❌ Error connecting to the database ❌❌❌" + "," + "Reason:" + err.Error())
	}

	// Configurar rutas
	routes.SetupRoutes(db)

	// Servidor
	log.Println("🏃‍♂️🏃‍♂️🏃‍♀️ Server running......")
	log.Println("🌏 Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("❌❌❌ Error starting server ❌❌❌" + "," + "Reason:" + err.Error())
	}
}
