package services

import (
	"github.com/ades1977/bookstore_users-api/domain/users"
	"github.com/ades1977/bookstore_users-api/utils/date_utils"
	"github.com/ades1977/bookstore_users-api/utils/errors"
)

func GetUser(userId int64)  ([]users.User, *errors.RestErr) {
	result := &users.User{}
	return result.Get(userId)
}


func Search(status string)  ([]users.User, *errors.RestErr) {
	result := &users.User{}
	return result.FindStatus(status)
}

func GetUserPaging(p1 int64,p2 int64)  ([]users.User, *errors.RestErr) {
	result := &users.User{}
	return result.GetPaging(p1, p2)
}


func CreateUser(user users.User) ([]users.User, *errors.RestErr){
	if err := user.Validate(); err != nil {
		return nil,err
	}

	user.CreateDate=date_utils.GetNowLocalString()
	user.Status = users.StatusRegister

	lastid,err := user.Save()
	//log.Println(lastid)
	if  err != nil{
		return  nil,err
	}

	currentLastIdUser, err := GetUser(lastid)
	if  err != nil{
		return  nil,err
	}
	return currentLastIdUser, nil
}

func UpdateUser(ispartial bool, userId int64,user users.User) ([]users.User, *errors.RestErr){
	currentUser, err := GetUser(userId)
	if err != nil {
		return nil,err
	}

	if ispartial {
		//cek kalau kosong isi current nya
		//log.Println(user.LastName)
		if user.FirstName =="" {
			user.FirstName = currentUser[0].FirstName
		}else{
			currentUser[0].FirstName=user.FirstName
		}
		if user.LastName =="" {
			user.LastName = currentUser[0].LastName
		}else{
			currentUser[0].LastName=user.LastName
		}
		if user.Email =="" {
			user.Email = currentUser[0].Email
		}else{
			currentUser[0].Email=user.Email
		}

	}else{
		//data current tidak kosong
		currentUser[0].FirstName = user.FirstName
		currentUser[0].LastName=user.LastName
		currentUser[0].Email=user.Email
	}

	user.CreateDate = currentUser[0].CreateDate
	if err := currentUser[0].Update(); err != nil{
		return  nil,err
	}

	return  currentUser, nil
}

func DeleteUser(userId int64) ([]users.User, *errors.RestErr){
	currentUser, err := GetUser(userId)
	if err != nil {
		return  nil,err
	}

	if  err := currentUser[0].Delete(); err != nil{
		return   nil,err
	}

	return  currentUser,nil
}