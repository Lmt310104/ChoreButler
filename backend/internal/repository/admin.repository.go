package repository

import (
	"chorebutler/internal/dtos"
	"chorebutler/internal/schemas"
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type IAdminRepository interface {
	CheckAdminExists(req dtos.LoginAdminRequest) (schemas.Admin, error)
	GetAdminProfile(adminID bson.ObjectID) (schemas.Admin, error)
}

type AdminRepository struct {
	db              *mongo.Database
	collection_name string
}

func (a AdminRepository) CheckAdminExists(req dtos.LoginAdminRequest) (schemas.Admin, error) {
	var admin schemas.Admin
	if err := a.db.Collection(a.collection_name).FindOne(context.TODO(), bson.M{"username": req.Username, "password": req.Password}).Decode(&admin); err != nil {
		return admin, err
	}
	return admin, nil
}

func (a AdminRepository) GetAdminProfile(adminID bson.ObjectID) (schemas.Admin, error) {
	var admin schemas.Admin
	if err := a.db.Collection(a.collection_name).FindOne(context.TODO(), bson.M{"id": adminID}).Decode(&admin); err != nil {
		return admin, err
	}
	return admin, nil
}

func NewAdminRepository(db *mongo.Database, collection_name string) IAdminRepository {
	return &AdminRepository{
		db:              db,
		collection_name: collection_name,
	}
}
