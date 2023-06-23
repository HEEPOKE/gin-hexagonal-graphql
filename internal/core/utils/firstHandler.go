package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloHandler godoc
// @Summary Hello example
// @Description do Hello
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} string "Hello"
// @Router / [get]
func HandleFirst(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}
