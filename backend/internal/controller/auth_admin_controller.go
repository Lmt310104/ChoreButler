package controller

import (
	"chorebutler/pkg/utils"
	"github.com/go-playground/validator/v10"
)

type AuthAdminController struct {
	cfg      utils.Config
	validate *validator.Validate
}
