package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"github.com/projects/COVID_Database/src/middlewares"
	"github.com/projects/COVID_Database/src/services"
	"github.com/projects/COVID_Database/src/validators"
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

func TestLabRegister(c *gin.Context) {

	testLabService := services.TestLabService{}

	registerRequest := validators.TestLabRegisterRequest{}
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		_ = c.Error(errors.New(middlewares.BadRequest)).SetType(gin.ErrorTypePublic)
		return
	}

	testLab, err := testLabService.AddTestLab(registerRequest)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusCreated, testLab)

}

func TestLabDelete(c *gin.Context) {

	testLabService := services.TestLabService{}

	testLabId, _ := strconv.Atoi(c.Param("test_lab_id"))
	if err := testLabService.DeleteTestLab(testLabId); err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.Status(http.StatusNoContent)
}
