package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

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

func TestLabRetrieve(c *gin.Context) {

	testLabService := services.TestLabService{}

	testLabId, _ := strconv.Atoi(c.Param("test_lab_id"))
	testLab, err := testLabService.GetTestLab(testLabId)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, testLab)

}
