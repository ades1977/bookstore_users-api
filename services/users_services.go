package services

import (
	"github.com/ades1977/bookstore_users-api/domain/users"
	"github.com/ades1977/bookstore_users-api/utils/errors"
)

func GetUser(userId int64)  (*users.User, *errors.RestErr) {
	result := users.User{Id: userId}
	if err := result.Get(); err !=nil {
		return nil,err
	}
	 return &result ,nil
}

func GetUserPaging(p1 int64,p2 int64)  (*users.User, *errors.RestErr) {
	result := users.User{PageFrom: p1,PageTo: p2}
	if err := result.GetPaging(); err !=nil {
		return nil, err
	}
	return &result ,nil
}


func CreateUser(user users.User) (*users.User, *errors.RestErr){
	if err := user.Validate(); err != nil {
		return nil,err
	}

	if err := user.Save(); err != nil{
		return  nil,err
	}
	return &user, nil
}