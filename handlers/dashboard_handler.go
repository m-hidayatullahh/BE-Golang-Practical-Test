package handlers

import (
    "encoding/json"
    "net/http"
    "room-booking-dashboard/models"
    "room-booking-dashboard/services"
    "strconv"
    "time"
)

type DashboardHandler struct {
    dashboardService *services.DashboardService
}

func NewDashboardHandler() *DashboardHandler {
    return &DashboardHandler{
        dashboardService: services.NewDashboardService(),
    }
}

// @Summary Get Dashboard Summary
// @Description Get room booking dashboard summary for a specific period
// @Tags dashboard
// @Produce json
// @Param year query int false "Year (default: current year)"
// @Param month query int false "Month (default: current month)"
// @Success 200 {object} models.ApiResponse
// @Failure 400 {object} models.ApiResponse
// @Failure 500 {object} models.ApiResponse
// @Router /api/v1/dashboard/summary [get]
func (h *DashboardHandler) GetDashboardSummary(w http.ResponseWriter, r *http.Request) {
    // Parse query parameters
    now := time.Now()
    year := now.Year()
    month := int(now.Month())
    
    if yearParam := r.URL.Query().Get("year"); yearParam != "" {
        if y, err := strconv.Atoi(yearParam); err == nil {
            year = y
        }
    }
    
    if monthParam := r.URL.Query().Get("month"); monthParam != "" {
        if m, err := strconv.Atoi(monthParam); err == nil && m >= 1 && m <= 12 {
            month = m
        }
    }
    
    // Get dashboard summary
    summary, err := h.dashboardService.GetDashboardSummary(year, month)
    if err != nil {
        h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to get dashboard summary", err)
        return
    }
    
    response := models.ApiResponse{
        Success: true,
        Message: "Dashboard summary retrieved successfully",
        Data:    summary,
    }
    
    h.sendJSONResponse(w, http.StatusOK, response)
}

// @Summary Get Raw Booking Data
// @Description Get raw booking data from external API
// @Tags dashboard
// @Produce json
// @Success 200 {object} models.ApiResponse
// @Failure 500 {object} models.ApiResponse
// @Router /api/v1/dashboard/raw-bookings [get]
func (h *DashboardHandler) GetRawBookings(w http.ResponseWriter, r *http.Request) {
    externalAPI := services.NewExternalAPIService()
    bookings, err := externalAPI.GetBookingList()
    if err != nil {
        h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to get raw booking data", err)
        return
    }
    
    response := models.ApiResponse{
        Success: true,
        Message: "Raw booking data retrieved successfully",
        Data:    bookings,
    }
    
    h.sendJSONResponse(w, http.StatusOK, response)
}

// @Summary Get Master Consumption Types
// @Description Get master consumption types from external API
// @Tags dashboard
// @Produce json
// @Success 200 {object} models.ApiResponse
// @Failure 500 {object} models.ApiResponse
// @Router /api/v1/dashboard/consumption-types [get]
func (h *DashboardHandler) GetConsumptionTypes(w http.ResponseWriter, r *http.Request) {
    externalAPI := services.NewExternalAPIService()
    consumptions, err := externalAPI.GetMasterJenisKonsumsi()
    if err != nil {
        h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to get consumption types", err)
        return
    }
    
    response := models.ApiResponse{
        Success: true,
        Message: "Consumption types retrieved successfully",
        Data:    consumptions,
    }
    
    h.sendJSONResponse(w, http.StatusOK, response)
}

func (h *DashboardHandler) sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(data)
}

func (h *DashboardHandler) sendErrorResponse(w http.ResponseWriter, statusCode int, message string, err error) {
    response := models.ApiResponse{
        Success: false,
        Message: message,
    }
    
    if err != nil {
        response.Message += ": " + err.Error()
    }
    
    h.sendJSONResponse(w, statusCode, response)
}