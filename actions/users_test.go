package actions

import (
	"github.com/projects/ai/models"
)

func (as *ActionSuite) Test_Users_Index() {
	res := as.HTML("/users").Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_Users_Show() {
	user := models.User{
		Name: "Testname",
		Age:  31,
	}
	models.DB.Create(&user)
	res := as.HTML("/users/" + user.ID.String()).Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Testname")
	as.Contains(res.Body.String(), "31")
}
