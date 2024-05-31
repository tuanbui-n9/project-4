package response

const (
	ErrorCodeSuccess  = 20001 // Success
	ErrorParamInvalid = 20003 // Invalid parameter
)

var msg = map[int]string{
	ErrorCodeSuccess:  "Success",
	ErrorParamInvalid: "Invalid parameter",
}
