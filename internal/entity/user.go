package entity

type User struct {
	ID       string       `json:"id" bson:"_id"`
	Encoding [256]float64 `json:"encoding" bson:"encoding"`
}
