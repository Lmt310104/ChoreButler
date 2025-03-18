package repository

import "go.mongodb.org/mongo-driver/v2/mongo"

type MongoDBRepository struct {
	database     *mongo.Database
	databaseName string
}
