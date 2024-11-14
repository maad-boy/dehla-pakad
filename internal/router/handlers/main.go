package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/maad-boy/dehla-pakad/internal/router/utils"
	"net/http"
)

func GinHandler[Req any, Res any](handler Handler[Req, Res]) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		req := new(Req)
		err := utils.BindReq(c, req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = handler.Validate(ctx, *req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := handler.Handle(ctx, *req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, res)
	}
}
