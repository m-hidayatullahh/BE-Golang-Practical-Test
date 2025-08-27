### Testing API Practical Test

### Dashboard Summary
```bash
# Get current month summary
curl "http://localhost:8080/api/v1/dashboard/summary"

# Get specific month summary
curl "http://localhost:8080/api/v1/dashboard/summary?year=2024&month=1"
```

### Raw Data
```bash
# Get raw booking data
curl "http://localhost:8080/api/v1/dashboard/raw-bookings"

# Get consumption types
curl "http://localhost:8080/api/v1/dashboard/consumption-types"
```

### Health Check
```bash
curl "http://localhost:8080/api/v1/health"
```

### Response Format
## Dashboard Summary Response
```bash
{
  "success": true,
  "message": "Dashboard summary retrieved successfully",
  "data": {
    "periode": "Januari 2024",
    "unitInduk": [
      {
        "unitIndukNumber": 1,
        "ruangan": [
          {
            "namaRuangan": "Nama Ruangan 1",
            "persentasePemakaian": 86.34,
            "nominalKonsumsi": 35000000,
            "snackSiang": 140,
            "makanSiangSiang": 280,
            "snackSore": 140,
            "totalTransaksi": 25
          }
        ]
      }
    ],
    "totalStats": {
      "totalNominalKonsumsi": 175000000,
      "totalSnackSiang": 700,
      "totalMakanSiangSiang": 1400,
      "totalSnackSore": 700,
      "rataRataPersentase": 86.34
    }
  }
}
```

## Features

✅ Dashboard Summary: Mengaggregasi data booking per unit dan ruangan <br>
✅ Period Filtering: Filter berdasarkan tahun dan bulan <br>
✅ Statistics Calculation: Menghitung persentase pemakaian, total konsumsi, dll <br>
✅ External API Integration: Mengambil data dari API eksternal <br>
✅ Error Handling: Proper error handling dan response format <br>
✅ CORS Support: Mendukung cross-origin requests <br>
✅ Health Check: Endpoint untuk monitoring <br>
✅ Clean Architecture: Pemisahan concern yang jelas

## Build & Run
```
docker build -t room-booking-dashboard .

docker run -p 8080:8080 room-booking-dashboard
```
