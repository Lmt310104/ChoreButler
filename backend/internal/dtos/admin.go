package dtos

type LoginAdminRequest struct {
	Username string `json:"username" binding:"required" validate:"required, min=1, max=255"`
	Password string `json:"password" binding:"required" validate:"required, min=1, max=255"`
}

type LoginAdminResponse struct {
	Token string `json:"token"`
}
