package users

import (
	"github.com/ades1977/bookstore_users-api/utils/errors"
	"strings"
)

//constanta harus huruf besar pada awal
const(
	StatusRegister = "REGISTER"
	StatusUnRegister = "UNREGISTER"
	StatusBanned="BANNED"
	StatusUnBanned="UNBANNED"
)

type User struct {
	Id         	int64  	`json:"id,omitempty"`			//dipergunakan untuk mapping
	FirstName  	string 	`json:"first_name,omitempty"`	//dipergunakan untuk mapping
	LastName   	string 	`json:"last_name,omitempty"`	//dipergunakan untuk mapping
	Email      	string 	`json:"email,omitempty"`		//dipergunakan untuk mapping
	CreateDate 	string 	`json:"create_date,omitempty"`	//dipergunakan untuk mapping
	Status		string 	`json:"status,omitempty"`
	Password	string	`json:"-"`	// - tidak menampilkan isinya secara komplit
	PageFrom  	int64 	`json:"page_from,omitempty"`	//dipergunakan untuk mapping
	PageTo  	int64 	`json:"page_to,omitempty"`	//dipergunakan untuk mapping, omitempty boleh kosong
}

type UserResponse struct {
	Status  int    				`json:"status,omitempty"`
	Message string 				`json:"message,omitempty"`
	Data    *User 				`json:"data"`
}

type UserResp struct {
	Status  int    				`json:"status,omitempty"`
	Message string 				`json:"message,omitempty"`
	Data    []User 				`json:"data"`
}

func (user *User) Validate() *errors.RestErr{
	user.FirstName=strings.TrimSpace(strings.ToLower(user.FirstName))
	user.LastName=strings.TrimSpace(strings.ToLower(user.LastName))
	user.CreateDate=strings.TrimSpace(strings.ToLower(user.CreateDate))
	user.Email=strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email==""{
		return errors.NewBedrequest("Bed Email Request")
	}


	user.Password=strings.TrimSpace(user.Password)
	if user.Password==""{
		return errors.NewBedrequest("Invalid Password")
	}
	return nil
}

