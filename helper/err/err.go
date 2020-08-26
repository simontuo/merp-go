package err

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// api错误的结构体
type APIException struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Request string `json:"request"`
}

// 实现接口
func (e *APIException) Error() string {
	return e.Message
}

func newAPIException(code int, msg string) *APIException {
	return &APIException{
		Code:    code,
		Message: msg,
	}
}

func HandleNotFound(c *gin.Context) {
	handleErr := NotFound()
	handleErr.Request = c.Request.Method + " " + c.Request.URL.String()
	c.JSON(handleErr.Code, handleErr)
	return
}

func HandleMethodNotAllowed(c *gin.Context) {
	handleErr := MethodNotAllowed()
	handleErr.Request = c.Request.Method + " " + c.Request.URL.String()
	c.JSON(handleErr.Code, handleErr)
	return
}

// 500 错误处理
func ServerError() *APIException {
	return newAPIException(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

// 404 错误
func NotFound() *APIException {
	return newAPIException(http.StatusNotFound, http.StatusText(http.StatusNotFound))
}

// 404 错误
func MethodNotAllowed() *APIException {
	return newAPIException(http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
}

// 未知错误
func UnknownError(message string) *APIException {
	return newAPIException(http.StatusForbidden, message)
}

// 参数错误
func ParameterError(message string) *APIException {
	return newAPIException(http.StatusBadRequest, message)
}

// 401错误
func UnauthorizedError(message string) *APIException {
	return newAPIException(http.StatusUnauthorized, message)
}
