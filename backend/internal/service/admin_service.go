package service

import (
	"chorebutler/internal/dtos"
	"chorebutler/internal/repository"
	"chorebutler/internal/schemas"
	"chorebutler/pkg/sanctum"
)

type AdminService struct {
	repo         repository.IAdminRepository
	sanctumToken *sanctum.Token
}

type IAdminService interface {
	CheckAdminExists(req dtos.LoginAdminRequest) (schemas.Admin, error)
	VerifyPassword(password, hashedPassword string) bool
}

func NewAdminService(repo repository.IAdminRepository, sanctumToken *sanctum.Token) IAdminService {
	return &AdminService{
		repo:         repo,
		sanctumToken: sanctumToken,
	}
}

func (as *AdminService) CheckAdminExists(req dtos.LoginAdminRequest) (schemas.Admin, error) {
	admin, err := as.repo.CheckAdminExists(req)
	if err != nil {
		return schemas.Admin{}, err
	}
	return admin, nil
}

func (as *AdminService) VerifyPassword(password, hashedPassword string) bool {
	return as.sanctumToken.Crypto.VerifyPassword(hashedPassword, password)
}

var _ IAdminService = (*AdminService)(nil)
