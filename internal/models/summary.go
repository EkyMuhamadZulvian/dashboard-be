package models

import "time"

// Summary merepresentasikan tabel Summary di MariaDB
type Summary struct {
	ID            int       `json:"id"`
	Date          time.Time `json:"date"`
	LocationCode  string    `json:"location_code"`
	Name          string    `json:"name"`
	VehicleType   string    `json:"vehicle_type"`
	Month         string    `json:"month"`
	Year          string    `json:"year"`
	Traffic       int       `json:"traffic"`
	Amount        float64   `json:"amount"`
	Region        string    `json:"region"`
	IdCategory    string    `json:"id_category"`
	IssuerName    string    `json:"issuer_name"`
	PaymentMethod string    `json:"payment_method"`
	GateInCode    string    `json:"gate_in_code"`
	GateOutCode   string    `json:"gate_out_code"`
}
