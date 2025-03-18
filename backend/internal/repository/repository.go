package repository

import (
	"chorebutler/pkg/utils"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type RepositoryContainer struct {
	AdminRepository IAdminRepository
}

type RepositoryFactory struct {
	db  *mongo.Database
	cfg utils.Config
}

func NewRepositoryFactory(db *mongo.Database, cfg utils.Config) *RepositoryFactory {
	return &RepositoryFactory{db: db, cfg: cfg}
}

func (f *RepositoryFactory) CreateRepositories() *RepositoryContainer {
	return &RepositoryContainer{
		AdminRepository: f.CreateAdminRepository(),
	}
}

func (f *RepositoryFactory) CreateAdminRepository() IAdminRepository {
	return NewAdminRepository(f.db, utils.ADMINS)
}
