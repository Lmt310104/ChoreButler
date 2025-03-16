package response

const (
	SuccessCode = 20001
	ErrorCode   = 40001
)

var msg = map[int]string{
	SuccessCode: "Success",
	ErrorCode:   "Error",
}
