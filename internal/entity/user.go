package entity

type User struct {
	ID       string       `json:"id"`
	Encoding [256]float64 `json:"encoding"`
}
