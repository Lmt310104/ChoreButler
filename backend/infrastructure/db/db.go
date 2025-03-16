package db

import (
	"chorebutler/pkg/utils"
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func NewDatabaseInstance(cfg utils.Config) *mongo.Client {
	_, cancel := context.WithTimeout(context.Background(), cfg.MongoTimeout)
	defer cancel()
	mongoDbConnection := cfg.MongoDBUri

	db, err := mongo.Connect(options.Client().ApplyURI(mongoDbConnection))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to MongoDB")
	}

	mongoDB := db.Database(cfg.MongoDBName)
	// Seed admin user
	if err := seedAdmin(mongoDB, cfg); err != nil {
		log.Fatal().Err(err).Msg("Failed to seed admin user")
	}
	
	return db
}
func seedAdmin(db *mongo.Database, cfg utils.Config) error {
	count, _ := db.Collection("users").CountDocuments(context.TODO(), bson.M{"email": "admin@gmail.com"})
	if count > 0 {
		return nil
	}

	//Generate hash password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), cfg.BcryptCost)
	if err != nil {
		return err
	}
	_, err = db.Collection("users").InsertOne(context.TODO(), bson.M{
		"full_name": "admin_name",
		"email":     "admin@gmail.com",
		"role":      "admin",
		"password":  string(hashPassword), // Convert sang string
	})
	if err != nil {
		return err
	}
	return nil
}
