package helpers

import (
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int
	Status  string
	Message string
	Result  []interface{}
}

func Response(c *gin.Context, code int, status string, Data []interface{}) {

	resp := ResponseData{
		Code:   code,
		Status: status,
		Result: Data,
	}
	c.JSON(code, resp)
}
