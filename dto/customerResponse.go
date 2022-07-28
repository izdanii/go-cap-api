package dto

type CustomerResponse struct {
	ID          string `json:"ID" db:"customer_id"`
	Name        string `json:"Name"`
	City        string `json:"City"`
	ZipCode     string `json:"ZipCode"`
	DateOfBirth string `json:"date_of_birth" db:"date_of_birth"`
	Status      string `json:"status"`
}
