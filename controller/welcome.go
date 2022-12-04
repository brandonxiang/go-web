package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.String(http.StatusOK, "Welcome!\n")
}