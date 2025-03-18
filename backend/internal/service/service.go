package service

import (
	"chorebutler/internal/repository"
	"chorebutler/pkg/sanctum"
	"chorebutler/pkg/utils"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ServiceContainer struct {
	AdminService IAdminService
}

type ServiceFactory struct {
	repo         *repository.IAdminRepository
	sanctumToken *sanctum.Token
}

func NewServiceFactory(db *mongo.Database, cfg utils.Config, sanctum *sanctum.Token) *ServiceFactory {
	repoFactory := repository.NewRepositoryFactory(db, cfg)
	repos := repoFactory.CreateRepositories()
	return &ServiceFactory{
		repo:         &repos.AdminRepository,
		sanctumToken: sanctum,
	}
}
func (f *ServiceFactory) CreateServices() *ServiceContainer {
	return &ServiceContainer{
		AdminService: f.createAdminService(),
	}
}

func (f *ServiceFactory) createAdminService() IAdminService {
	return NewAdminService(*f.repo, f.sanctumToken)
}
