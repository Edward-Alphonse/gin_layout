package router

import (
	"github.com/gin-gonic/gin"

	"gin_layout/api"
	"gin_layout/internal/biz/handler"
)

func Hello(c *gin.Context) {
	req := &api.HelloRequest{}
	h := handler.NewHelloHandler(c, req)
	if err := h.Handle(); err != nil {
		c.JSON(500, gin.H{})
		return
	}
	c.JSON(200, gin.H{
		"message": h.Resp,
	})
}
