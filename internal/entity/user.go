package entity

type EncodingType [256]float64

type User struct {
	ID       string       `json:"id" bson:"_id"`
	Encoding EncodingType `json:"encoding" bson:"encoding"`
}
