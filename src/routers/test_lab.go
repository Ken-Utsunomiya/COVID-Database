package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/projects/COVID_Database/src/handlers"
)

func TestLabRoutes(rg *gin.RouterGroup) {

	testLab := rg.Group("/test_labs")
	{
		testLab.GET("/", handlers.TestLabList)
	}
}
