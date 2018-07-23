package models

import (
	"github.com/gobuffalo/suite"
)

type ActionSuite struct {
	*suite.Action
}

func (as *ActionSuite) Test_User() {

	username := "John Smith"
	age := int8(18)

	user := User{
		Name: username,
		Age:  age,
	}

	as.Equal("sa", user.Name)
}
