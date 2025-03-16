package service

import (
	"chorebutler/internal/dtos"
	"chorebutler/internal/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type AdminService struct {
	db mongo.Database
}

type IAdminService interface {
	CheckAdminExists(req dtos.LoginAdminRequest)
	VerifyPassword(password, hashedPassword string) bool
	CreateToken(admin models.User) (string, error)
	GetAdminProfile(adminID string) (models.User, error)
	GetDashboardData()
	GetUserDashboardData()
}

//func NewAdminService(db mongo.Database) IAdminService {
//	return &AdminService{
//		db: db
//	}
//
//}
//
//func (as *AdminService) CheckAdminExists (req dtos.LoginAdminRequest)(models.Admin, error){
//	return as.db.Collection("admins").FindOne()
//}

var _IAdminService = (*AdminService)(nil)
