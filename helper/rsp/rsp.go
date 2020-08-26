package rsp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func EmptyDataRsp(c *gin.Context) error {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": make([]string, 0),
	})
	return nil
}

func EmptyItemsRsp(c *gin.Context) error {
	data := make(map[string][]string)
	data["items"] = []string{}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
	return nil
}

func OkRsp(c *gin.Context, data interface{}) error {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
	return nil
}

func OkStatusRsp(c *gin.Context, data interface{}) error {
	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"status": data,
	})
	return nil
}

func OkPageRsp(c *gin.Context, data interface{}, total int) error {
	d := make(map[string]interface{})
	d["items"] = data
	d["total"] = total

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": d,
	})
	return nil
}

func OkMsgRsp(c *gin.Context, data interface{}, message string) error {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    data,
		"message": message,
	})
	return nil
}
