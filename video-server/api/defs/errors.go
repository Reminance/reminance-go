package defs

//
type Err struct {
	Error     string
	ErrorCode string
}

type ErrorResponse struct {
	HttpSC int
	Error  Err
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{HttpSC: 400, Error: Err{Error: "Request body is not correct", ErrorCode: "001"}}
	ErrorNotAuthorized          = ErrorResponse{HttpSC: 401, Error: Err{Error: "user authentication is not valid", ErrorCode: "002"}}
	ErrorDBError                = ErrorResponse{HttpSC: 500, Error: Err{Error: "DB ops failed", ErrorCode: "003"}}
	ErrorInternalFaults         = ErrorResponse{HttpSC: 500, Error: Err{Error: "Error Internal Faults", ErrorCode: "004"}}
)
