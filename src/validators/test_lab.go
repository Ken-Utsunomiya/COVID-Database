package validators

import "github.com/projects/COVID_Database/src/models"

type TestLabRegisterRequest struct {
	Name       string `json:"name" binding:"required,max=30"`
	TestMethod string `json:"test_method" binding:"required,max=20"`
}

func RegisterRequestToTestLabModel(request TestLabRegisterRequest) models.TestLab {

	testLab := models.TestLab{}
	testLab.Name = request.Name
	testLab.TestMethod = request.TestMethod
	return testLab

}
