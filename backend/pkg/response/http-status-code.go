package response

const (
	SuccessCode        = 20001
	ErrorCode          = 50001 // internal server error
	AdminAuthErrorCode = 40002
	AdminValidateError = 40003
)

var msg = map[int]string{
	SuccessCode:        "Success",
	ErrorCode:          "Error",
	AdminAuthErrorCode: "Mật khẩu hoặc tên đăng nhập không đúng",
}
