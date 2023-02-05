package model

type Item struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Quantity  int    `json:"quantity"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}
