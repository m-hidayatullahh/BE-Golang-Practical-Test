package utils

import (
    "fmt"
    "time"
)

func FormatPeriode(year int, month int) string {
    monthNames := []string{
        "Januari", "Februari", "Maret", "April", "Mei", "Juni",
        "Juli", "Agustus", "September", "Oktober", "November", "Desember",
    }
    
    if month < 1 || month > 12 {
        return fmt.Sprintf("%d", year)
    }
    
    return fmt.Sprintf("%s %d", monthNames[month-1], year)
}

func ParseDate(dateStr string) (time.Time, error) {
    layouts := []string{
        "2006-01-02T15:04:05Z",
        "2006-01-02T15:04:05.000Z",
        "2006-01-02",
        "2006-01-02 15:04:05",
    }
    
    for _, layout := range layouts {
        if t, err := time.Parse(layout, dateStr); err == nil {
            return t, nil
        }
    }
    
    return time.Time{}, fmt.Errorf("unable to parse date: %s", dateStr)
}