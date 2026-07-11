package models

type Customer struct {
	ID            int64   `bson:"_id"`
	Username      string  `bson:"username"`
	PaymentStatus string  `bson:"payment_status"`
	Orders        []Order `bson:"orders"`
}
