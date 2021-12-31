package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/projects/COVID_Database/src/services"
)

func TestLabList(c *gin.Context) {

	testLabService := services.TestLabService{}

	testLabList, err := testLabService.GetTestLabs()
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"test_labs": testLabList,
	})

}
