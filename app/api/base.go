package api

import (
	"github.com/all-skeleton/gin-skeleton/app/library"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(c *gin.Context, code int, data interface{}) bool {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  library.GetMsg(code),
		"data": data,
	})
	return true
}

func ResponseSuccess(c *gin.Context, data interface{}) bool {
	c.JSON(http.StatusOK, gin.H{
		"code": library.SUCCESS,
		"msg":  library.GetMsg(library.SUCCESS),
		"data": data,
	})
	return true
}

func ResponseError(c *gin.Context, msg string, data interface{}) bool {
	c.JSON(http.StatusOK, gin.H{
		"code": library.ERROR,
		"msg":  msg,
		"data": data,
	})
	return true
}
