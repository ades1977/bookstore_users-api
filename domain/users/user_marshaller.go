package users

import "encoding/json"

type PublicUser struct{
	Id         	int64  	`json:"id,omitempty"`			//dipergunakan untuk mapping
	CreateDate 	string 	`json:"create_date,omitempty"`	//dipergunakan untuk mapping
	Status		string 	`json:"status,omitempty"`
}

type PrivateUser struct{
	Id         	int64  	`json:"id,omitempty"`			//dipergunakan untuk mapping
	FirstName  	string 	`json:"first_name,omitempty"`	//dipergunakan untuk mapping
	LastName   	string 	`json:"last_name,omitemCpty"`	//dipergunakan untuk mapping
	Email      	string 	`json:"email,omitempty"`		//dipergunakan untuk mapping
	CreateDate 	string 	`json:"create_date,omitempty"`	//dipergunakan untuk mapping
	Status		string 	`json:"status,omitempty"`
}

func (user *User) Marshall(isPublic bool) interface{}{
	if isPublic {
		return PublicUser{
			Id : user.Id,
			CreateDate: user.CreateDate,
			Status: user.Status,
		}
	}
	userJson, _ := json.Marshal(user)
	var privateUser  PrivateUser
	json.Unmarshal(userJson,&privateUser)
	return  privateUser
}