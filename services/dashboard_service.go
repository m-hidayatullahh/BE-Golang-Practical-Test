package services

import (
    "fmt"
    "room-booking-dashboard/models"
    "room-booking-dashboard/utils"
    "strconv"
    "time"
)

type DashboardService struct {
    externalAPI *ExternalAPIService
}

func NewDashboardService() *DashboardService {
    return &DashboardService{
        externalAPI: NewExternalAPIService(),
    }
}

func (s *DashboardService) GetDashboardSummary(year int, month int) (*models.DashboardSummary, error) {
    // Fetch data from external APIs
    bookings, err := s.externalAPI.GetBookingList()
    if err != nil {
        return nil, fmt.Errorf("failed to get bookings: %w", err)
    }

    consumptions, err := s.externalAPI.GetMasterJenisKonsumsi()
    if err != nil {
        return nil, fmt.Errorf("failed to get consumption types: %w", err)
    }

    // Create consumption map for easy lookup
    consumptionMap := make(map[int]string)
    for _, c := range consumptions {
        consumptionMap[c.ID] = c.Name
    }

    // Filter bookings by period
    filteredBookings := s.filterBookingsByPeriod(bookings, year, month)
    
    // Generate summary
    summary := s.generateSummary(filteredBookings, consumptionMap, year, month)
    
    return summary, nil
}

func (s *DashboardService) filterBookingsByPeriod(bookings []models.Booking, year int, month int) []models.Booking {
    var filtered []models.Booking
    
    for _, booking := range bookings {
        if booking.TanggalPemesanan.Year() == year && int(booking.TanggalPemesanan.Month()) == month {
            filtered = append(filtered, booking)
        }
    }
    
    return filtered
}

func (s *DashboardService) generateSummary(bookings []models.Booking, consumptionMap map[int]string, year int, month int) *models.DashboardSummary {
    // Group by unit and room
    unitMap := make(map[int]map[string][]models.Booking)
    
    for _, booking := range bookings {
        if unitMap[booking.UnitInduk] == nil {
            unitMap[booking.UnitInduk] = make(map[string][]models.Booking)
        }
        
        unitMap[booking.UnitInduk][booking.NamaRuangan] = append(
            unitMap[booking.UnitInduk][booking.NamaRuangan], 
            booking,
        )
    }

    var unitSummaries []models.UnitSummary
    var totalStats models.TotalStats

    // Process each unit
    for unitNumber, rooms := range unitMap {
        var ruanganStats []models.RuanganStats
        
        for roomName, roomBookings := range rooms {
            stats := s.calculateRoomStats(roomBookings, consumptionMap)
            stats.NamaRuangan = roomName
            ruanganStats = append(ruanganStats, stats)
            
            // Add to total stats
            totalStats.TotalNominalKonsumsi += stats.NominalKonsumsi
            totalStats.TotalSnackSiang += stats.SnackSiang
            totalStats.TotalMakanSiangSiang += stats.MakanSiangSiang
            totalStats.TotalSnackSore += stats.SnackSore
            totalStats.RataRataPersentase += stats.PersentasePemakaian
        }
        
        unitSummaries = append(unitSummaries, models.UnitSummary{
            UnitIndukNumber: unitNumber,
            Ruangan:        ruanganStats,
        })
    }

    // Calculate average percentage
    totalRooms := 0
    for _, unit := range unitSummaries {
        totalRooms += len(unit.Ruangan)
    }
    if totalRooms > 0 {
        totalStats.RataRataPersentase = totalStats.RataRataPersentase / float64(totalRooms)
    }

    periode := utils.FormatPeriode(year, month)
    
    return &models.DashboardSummary{
        Periode:    periode,
        UnitInduk:  unitSummaries,
        TotalStats: totalStats,
    }
}

func (s *DashboardService) calculateRoomStats(bookings []models.Booking, consumptionMap map[int]string) models.RuanganStats {
    var stats models.RuanganStats
    
    totalPersentase := 0.0
    snackSiang := 0
    makanSiang := 0  
    snackSore := 0
    
    for _, booking := range bookings {
        stats.NominalKonsumsi += booking.NominalKonsumsi
        totalPersentase += booking.PersentasePemakaian
        stats.TotalTransaksi++
        
        // Count consumption types based on name/ID
        consumptionName := consumptionMap[booking.JenisKonsumsiID]
        switch consumptionName {
        case "Snack Siang":
            snackSiang++
        case "Makan Siang":
            makanSiang++
        case "Snack Sore":
            snackSore++
        }
    }
    
    // Calculate average percentage
    if len(bookings) > 0 {
        stats.PersentasePemakaian = totalPersentase / float64(len(bookings))
    }
    
    stats.SnackSiang = snackSiang
    stats.MakanSiangSiang = makanSiang
    stats.SnackSore = snackSore
    
    return stats
}