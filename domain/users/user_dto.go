package users

import (
	"github.com/ades1977/bookstore_users-api/utils/errors"
	"strings"
)

type User struct {
	Id         int64  `json:"id"`			//dipergunakan untuk mapping
	FirstName  string `json:"first_name"`	//dipergunakan untuk mapping
	LastName   string `json:"last_name"`	//dipergunakan untuk mapping
	Email      string `json:"email"`		//dipergunakan untuk mapping
	CreateDate string `json:"create_date"`	//dipergunakan untuk mapping
}


func (user *User) Validate() *errors.RestErr{
	user.Email=strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email==""{
		return errors.NewBedrequest("Bed Email Request")
	}
	return nil
}