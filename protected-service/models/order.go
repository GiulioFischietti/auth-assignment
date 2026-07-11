package models

type Order struct {
	ID string `bson:"_id"`

	Customer CustomerInfo `bson:"customer"`

	Payment PaymentInfo `bson:"payment"`

	OrderStatus string `bson:"order_status"`

	Items []Item `bson:"items"`
}

type CustomerInfo struct {
	ID       int64  `bson:"id"`
	Username string `bson:"username"`
}

type PaymentInfo struct {
	Status string `bson:"status"`
}

type Item struct {
	Product  string  `bson:"product"`
	Quantity int     `bson:"quantity"`
	Price    float64 `bson:"price"`
}
