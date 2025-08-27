package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "room-booking-dashboard/handlers"

    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
)

// @title Room Booking Dashboard API
// @version 1.0
// @description API untuk dashboard summary pemesanan ruangan
// @host localhost:8080
// @BasePath /api/v1

func main() {
    r := mux.NewRouter()
    
    // CORS middleware
    corsOptions := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
    corsOrigins := handlers.AllowedOrigins([]string{"*"})
    corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
    
    // Initialize handlers
    dashboardHandler := handlers.NewDashboardHandler()
    
    // API Routes
    api := r.PathPrefix("/api/v1").Subrouter()
    
    // Dashboard endpoints
    api.HandleFunc("/dashboard/summary", dashboardHandler.GetDashboardSummary).Methods("GET")
    api.HandleFunc("/dashboard/raw-bookings", dashboardHandler.GetRawBookings).Methods("GET")
    api.HandleFunc("/dashboard/consumption-types", dashboardHandler.GetConsumptionTypes).Methods("GET")
    
    // Health check
    api.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"status": "healthy", "service": "room-booking-dashboard"}`))
    }).Methods("GET")
    
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    
    fmt.Printf("🚀 Room Booking Dashboard API running on port %s\n", port)
    fmt.Printf("📊 Dashboard Summary: http://localhost:%s/api/v1/dashboard/summary\n", port)
    fmt.Printf("📋 Raw Bookings: http://localhost:%s/api/v1/dashboard/raw-bookings\n", port)
    fmt.Printf("🍽️  Consumption Types: http://localhost:%s/api/v1/dashboard/consumption-types\n", port)
    fmt.Printf("❤️  Health Check: http://localhost:%s/api/v1/health\n", port)
    
    log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(corsOptions, corsOrigins, corsMethods)(r)))
}