package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/projects/COVID_Database/src/handlers"
)

func TestLabRoutes(rg *gin.RouterGroup) {

	testLab := rg.Group("/test_labs")
	{
		testLab.GET("/", handlers.TestLabList)
		testLab.GET("/:test_lab_id", handlers.TestLabRetrieve)
		testLab.POST("/", handlers.TestLabRegister)
		testLab.DELETE("/:test_lab_id", handlers.TestLabDelete)
	}
}
