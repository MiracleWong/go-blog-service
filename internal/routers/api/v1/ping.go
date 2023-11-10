package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Ping struct {
}

func NewPing() Ping {
	return Ping{}
}

func (t Ping) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong",
	})
}
