package response

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ErrorResponse(err error, data ErrData) Response {
	return Response{Success: false, Message: err.Error(), Data: data}
}

func ErrorCustomRepsonse(message string) Response {
	return Response{Success: false, Message: message}
}
func ErrorResponseWithoutData(err error) Response {
	return Response{Success: false, Message: err.Error()}
}
func SuccessResponse(data interface{}) Response {
	return Response{Success: true, Data: data}
}

func SuccessResponseWithMessage(data interface{}, message string) Response {
	return Response{Success: true, Message: message, Data: data}
}
