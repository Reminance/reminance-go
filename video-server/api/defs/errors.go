package defs

//
type Err struct {
	Error string
	ErrorCode string
}

type ErrorResponse struct {
	HttpSC int
	Error Err
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{HttpSC:400, Error:Err{Error:"Request body is not correct", ErrorCode:"001"}}
	ErrorNotAuthorized = ErrorResponse{HttpSC:401, Error:Err{Error:"user authentication is not valid", ErrorCode:"002"}}
)