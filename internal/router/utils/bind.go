package utils

import "github.com/gin-gonic/gin"

func BindReq(c *gin.Context, req interface{}) error {
	return c.BindQuery(req)
}
