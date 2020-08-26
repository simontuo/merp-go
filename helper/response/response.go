package response

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Err(c *gin.Context, code int, detail string) {
	c.JSON(code, gin.H{
		"code":   code,
		"detail": detail,
		"status": "error",
	})
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}

type ErrMes struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
	Status string `json:"status"`
}

func MicroErr(c *gin.Context, err error) {
	var msg ErrMes
	e := json.Unmarshal([]byte(err.Error()), &msg)
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":   http.StatusInternalServerError,
			"detail": e.Error(),
			"status": "error",
		})
	} else {
		c.JSON(msg.Code, gin.H{
			"code":   msg.Code,
			"detail": msg.Detail,
			"Status": msg.Status,
		})
	}
}

func MsgOK(c *gin.Context, data interface{}, message string) {
	m := make(map[string]interface{})
	m["message"] = message
	m["data"] = data
	m["Status"] = "success"

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": m,
	})
}
