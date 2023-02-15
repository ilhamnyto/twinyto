package response

import "net/http"

type Response struct {
	StatusCode int `json:"status_code"`
	Message string `json:"message"`
	Payload struct {
		Data interface{} `json:"data,omitempty"`
	} `json:"payload,omitempty"`
}

var (
	generalSuccess = Response{
		StatusCode: http.StatusOK,
		Message: "SUCCESS",
	}

	createdSuccess = Response{
		StatusCode: http.StatusCreated,
		Message: "SUCCESS_CREATED",
	}
)

func GeneralSuccess() *Response {
	succ := generalSuccess
	return &succ
}

func GeneralSuccessWithCustomMessageAndPayload(message string, payload interface{}) *Response {
	succ := generalSuccess
	succ.Message = message
	succ.Payload.Data = payload

	return &succ
}

func CreatedSuccess() *Response {
	succ := createdSuccess
	return &succ
}

func CreatedSuccessWithPayload(payload interface{}) *Response {
	succ := createdSuccess
	succ.Payload.Data = payload
	return &succ
}