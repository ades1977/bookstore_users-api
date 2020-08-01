package users

import (
	"fmt"
	"github.com/ades1977/bookstore_users-api/utils/errors"
)

var (
	userDB=make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr{
	result := userDB[user.Id]
	if result == nil {
		return  errors.NewNotFound(fmt.Sprintf("user %d not found",user.Id))
	}
	user.Id=result.Id
	user.FirstName=result.FirstName
	user.LastName=result.LastName
	user.Email=result.Email
	user.CreateDate=result.CreateDate

	return nil
}

func (user *User) Save() *errors.RestErr{
	current := userDB[user.Id]
	if current != nil {
		if current.Email==user.Email{
			return errors.NewBedrequest(fmt.Sprintf("Email %s Already Register",user.Email))
		}
		return errors.NewBedrequest(fmt.Sprintf("User %d Already Exist",user.Id))
	}
	userDB[user.Id]= user
	return nil
}