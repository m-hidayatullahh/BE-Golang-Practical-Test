package models

type ConsumptionResponse struct {
    Success bool          `json:"success"`
    Data    []Consumption `json:"data"`
}

type Consumption struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}
