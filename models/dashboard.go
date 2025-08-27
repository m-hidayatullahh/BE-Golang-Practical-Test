package models

type DashboardSummary struct {
    Periode    string       `json:"periode"`
    UnitInduk  []UnitSummary `json:"unitInduk"`
    TotalStats TotalStats   `json:"totalStats"`
}

type UnitSummary struct {
    UnitIndukNumber int            `json:"unitIndukNumber"`
    Ruangan        []RuanganStats `json:"ruangan"`
}

type RuanganStats struct {
    NamaRuangan             string  `json:"namaRuangan"`
    PersentasePemakaian     float64 `json:"persentasePemakaian"`
    NominalKonsumsi         float64 `json:"nominalKonsumsi"`
    SnackSiang              int     `json:"snackSiang"`
    MakanSiangSiang         int     `json:"makanSiangSiang"`
    SnackSore               int     `json:"snackSore"`
    TotalTransaksi          int     `json:"totalTransaksi"`
}

type TotalStats struct {
    TotalNominalKonsumsi    float64 `json:"totalNominalKonsumsi"`
    TotalSnackSiang         int     `json:"totalSnackSiang"`
    TotalMakanSiangSiang    int     `json:"totalMakanSiangSiang"`
    TotalSnackSore          int     `json:"totalSnackSore"`
    RataRataPersentase      float64 `json:"rataRataPersentase"`
}

type ApiResponse struct {
    Success bool        `json:"success"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}