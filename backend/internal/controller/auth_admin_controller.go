package controller

import (
	"chorebutler/internal/dtos"
	"chorebutler/internal/service"
	"chorebutler/pkg/response"
	"chorebutler/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthAdminController struct {
	cfg          utils.Config
	validate     *validator.Validate
	AdminService service.IAdminService
}

func NewAuthAdminController(cfg utils.Config, validate *validator.Validate, adminService service.IAdminService) *AuthAdminController {
	return &AuthAdminController{
		cfg:          cfg,
		validate:     validate,
		AdminService: adminService,
	}
}

// Login handles the login request for admin users
// @Summary Login as an admin
// @Description Authenticates an admin user and returns a token along with user information
// @Tags admin/auth
// @Accept json
// @Produce json
// @Param request body dtos.LoginAdminRequest true "The login request"
// @Success 200 {object} response.ResponseData "Login successful"
// @Failure 400 {object} response.ResponseData "Invalid request body or input validation failed"
// @Failure 401 {object} response.ResponseData "Invalid credentials"
// @Failure 500 {object} response.ResponseData "Internal server error"
// @Router /admin/auth/login [post]
func (ac *AuthAdminController) Login(ctx *gin.Context) {
	var req dtos.LoginAdminRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(ctx, response.AdminAuthErrorCode, "Yêu cầu không hợp lệ")
		return
	}

	if err := ac.validate.Struct(req); err != nil {
		response.ErrorResponse(ctx, response.AdminValidateError, "Dữ liệu đầu vào không hợp lệ")
		return
	}

	admin, err := ac.AdminService.CheckAdminExists(req)
	if err != nil {
		response.ErrorResponse(
			ctx, response.AdminAuthErrorCode, "Mật khẩu hoặc tên đăng nhập không đúng")
		return
	}

	if !ac.AdminService.VerifyPassword(req.Password, admin.Password) {
		response.ErrorResponse(
			ctx, response.AdminAuthErrorCode, "Mật khẩu hoặc tên đăng nhập không đúng")
		return
	}
}
