package services

import (
    "encoding/json"
    "fmt"
    "net/http"
    "room-booking-dashboard/models"
    "time"
)

type ExternalAPIService struct {
    httpClient *http.Client
}

func NewExternalAPIService() *ExternalAPIService {
    return &ExternalAPIService{
        httpClient: &http.Client{
            Timeout: 30 * time.Second,
        },
    }
}

func (s *ExternalAPIService) GetBookingList() ([]models.Booking, error) {
    url := "https://66876cc30bc7155dc017a662.mockapi.io/api/dummy-data/bookingList"
    
    resp, err := s.httpClient.Get(url)
    if err != nil {
        return nil, fmt.Errorf("failed to fetch booking list: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("API returned status code: %d", resp.StatusCode)
    }

    var bookings []models.Booking
    if err := json.NewDecoder(resp.Body).Decode(&bookings); err != nil {
        return nil, fmt.Errorf("failed to decode booking response: %w", err)
    }

    return bookings, nil
}

func (s *ExternalAPIService) GetMasterJenisKonsumsi() ([]models.Consumption, error) {
    url := "https://6686cb5583c983911b03a7f3.mockapi.io/api/dummy-data/masterJenisKonsumsi"
    
    resp, err := s.httpClient.Get(url)
    if err != nil {
        return nil, fmt.Errorf("failed to fetch consumption types: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("API returned status code: %d", resp.StatusCode)
    }

    var consumptions []models.Consumption
    if err := json.NewDecoder(resp.Body).Decode(&consumptions); err != nil {
        return nil, fmt.Errorf("failed to decode consumption response: %w", err)
    }

    return consumptions, nil
}