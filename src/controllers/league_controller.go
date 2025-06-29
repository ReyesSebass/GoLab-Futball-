package controllers

import (
    "encoding/json"
    "net/http"

    "football-app/src/models"
    "gorm.io/gorm"
)

func GetLeaguesHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var leagues []models.League
        if err := db.Find(&leagues).Error; err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(map[string]string{"error": "No se pudieron obtener las ligas"})
            return
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(leagues)
    }
}
