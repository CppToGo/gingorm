package main

import (
	"gingorm/src/Framwork"
	"net/http"

	"github.com/gin-gonic/gin"
)


type TestHandle struct {
	Framwork.Handle
}

func (h *TestHandle) SetRootRouter() {
	h.RootRouter = "/test"
}

func (h *TestHandle) Process(group *gin.RouterGroup) error {
	statuCode := http.StatusOK

	group.GET("/:words", func(c *gin.Context) {
		words := c.Param("words")
		c.String(statuCode, "%s", words)
	})

	return nil
}