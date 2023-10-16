package response

type Response struct {
	Code    uint64      `json:"code" example:"2000"`
	Message string      `json:"message" example:"Success."`
	Data    interface{} `json:"data,omitempty"`
}

type ErrResponse struct {
	Code    uint64 `json:"code" example:"4000"`
	Message string `json:"message" example:"Fail."`
	Error   string `json:"error,omitempty" example:"'<Field>' must be REQUIRED field but the input is ''"`
}

func NewResponse(code uint64, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func NewErrResponse(code uint64, message string, err string) *ErrResponse {
	return &ErrResponse{
		Code:    code,
		Message: message,
		Error:   err,
	}
}
