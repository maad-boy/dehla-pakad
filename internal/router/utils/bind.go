package utils

import "github.com/gin-gonic/gin"

func BindReq(c *gin.Context, req interface{}) error {
	err := c.BindHeader(req)
	if err != nil {
		return err
	}

	err = c.ShouldBind(req)
	if err != nil {
		return err
	}

	err = c.BindQuery(req)
	if err != nil {
		return err
	}

	err = c.BindUri(req)
	if err != nil {
		return err
	}

	return nil
}
