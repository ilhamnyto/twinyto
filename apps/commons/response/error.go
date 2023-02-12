package response

import "net/http"

type CustomError struct {
	Code string `json:"code"`
	StatusCode int `json:"status_code"`
	Message string `json:"message"`
	AdditionalInfo interface{} `json:"additional_info,omitempty"`
}

var (
	generalError = CustomError{
		Code: "ERR",
		StatusCode: http.StatusInternalServerError,
		Message: "ERROR",
	}

	repositoryError = CustomError {
		Code: "ERR_REPO",
		StatusCode: http.StatusInternalServerError,
		Message: "REPOSITORY ERROR",
	}

	notFoundError = CustomError{
		Code: "ERR_NOT_FOUND",
		StatusCode: http.StatusNotFound,
		Message: "NOT FOUND",
	}

	unAuthorizedError = CustomError{
		Code: "ERR_NOT_AUTHORIZED",
		StatusCode: http.StatusUnauthorized,
		Message: "UNAUTHORIZED ERROR",
	}
)

func GeneralError() *CustomError {
	err := generalError
	return &err
}

func GeneralErrorWithAdditionalInfo(info interface{}) *CustomError {
	err := generalError
	err.AdditionalInfo = info
	return &err
}

func RepositoryError() *CustomError {
	err := repositoryError
	return &err
}

func RepositoryErrorWithAdditionalInfo(info interface{}) *CustomError {
	err := repositoryError
	err.AdditionalInfo = info
	return &err
}

func NotFoundError() *CustomError {
	err := notFoundError
	return &err
}

func NotFoundErrorWithAdditionalInfo(info interface{}) *CustomError {
	err := notFoundError
	err.AdditionalInfo = info
	return &err
}

func UnauthorizedError() *CustomError {
	err := unAuthorizedError
	return &err
}

func UnauthorizedErrorWithAdditionalInfo(info interface{}) *CustomError {
	err := unAuthorizedError
	err.AdditionalInfo = info
	return &err
}