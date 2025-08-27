package models

import "time"

type BookingResponse struct {
    Success bool      `json:"success"`
    Data    []Booking `json:"data"`
}

type Booking struct {
    ID                    string    `json:"id"`
    NomorTransaksi        string    `json:"nomorTransaksi"`
    TanggalPemesanan      time.Time `json:"tanggalPemesanan"`
    TanggalMulaiSewa      time.Time `json:"tanggalMulaiSewa"`
    TanggalSelesaiSewa    time.Time `json:"tanggalSelesaiSewa"`
    NamaRuangan           string    `json:"namaRuangan"`
    UnitInduk             int       `json:"unitInduk"`
    NominalKonsumsi       float64   `json:"nominalKonsumsi"`
    JenisKonsumsiID       int       `json:"jenisKonsumsiId"`
    StatusPemesanan       string    `json:"statusPemesanan"`
    PersentasePemakaian   float64   `json:"persentasePemakaian"`
}