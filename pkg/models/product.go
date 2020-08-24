package models

//Product model
type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
}
