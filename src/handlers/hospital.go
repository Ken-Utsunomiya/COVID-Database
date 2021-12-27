package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HospitalList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}
