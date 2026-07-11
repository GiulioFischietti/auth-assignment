package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"protected-service/models"
)

type OrderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(
	db *mongo.Database,
) *OrderRepository {

	return &OrderRepository{

		collection: db.Collection("orders"),
	}
}

func (r *OrderRepository) FindByCustomerID(
	ctx context.Context,
	userID int64,
) ([]models.Order, error) {

	cursor, err :=
		r.collection.Find(
			ctx,
			bson.M{
				"customer.id": userID,
			},
		)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var orders []models.Order

	err =
		cursor.All(
			ctx,
			&orders,
		)

	return orders, err
}
