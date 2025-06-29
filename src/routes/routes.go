package routes

import (
    "net/http"

    "football-app/src/controllers"
    "gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) {
    http.HandleFunc("/leagues", controllers.GetLeaguesHandler(db))

    // Si usás archivos estáticos como CSS o imágenes
    http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("views/assets"))))

    // Si usás plantillas HTML
    http.Handle("/", http.FileServer(http.Dir("views")))
}
