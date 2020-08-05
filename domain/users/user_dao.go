package users

import (
	"fmt"
	"github.com/ades1977/bookstore_users-api/databases/mysql/users_db"
	"github.com/ades1977/bookstore_users-api/utils/date_utils"
	"github.com/ades1977/bookstore_users-api/utils/errors"
	"log"
	"strings"
)

const(
	errorNorows 	= 	"no rows in result set"
	queryInsertUser = 	"INSERT INTO users(first_name,last_name,email, create_date) values(?,?,?,?);"
	querySelectUser	= 	"SELECT * FROM users WHERE id=? limit 1 ;"
	queryPagingUser	=	"select * from users order by id desc  limit ?,? ;"
)


var (
	userDB=make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr{
	stm, err := users_db.Client.Prepare(querySelectUser)
	if err != nil{
		return  errors.NewInternalServerError(err.Error())
	}
	defer stm.Close()
	rows := stm.QueryRow(user.Id)  //untuk single row
	if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreateDate);
			err != nil   {
				if strings.Contains(err.Error(),errorNorows){
					return errors.NewNotFound(fmt.Sprintf("User Id %d not found", user.Id))
				}
		return  errors.NewInternalServerError(fmt.Sprintf("Error when trying user %d :  %s ",user.Id, err.Error()))
	}
	return nil
}


func (user *User) GetPaging() *errors.RestErr{
	stm, err := users_db.Client.Prepare(queryPagingUser)
	if err != nil{
		return  errors.NewInternalServerError(err.Error())
	}

	defer stm.Close()
	rows, err := stm.Query(user.PageFrom,user.PageTo) //untuk banyak row
	if err != nil{
		return  errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	u := User{}
	res := []User{}
	for rows.Next() {
		var er =   rows.Scan(&u.Id, &u.FirstName, &u.LastName,
				&u.Email,&u.CreateDate)
		if er != nil {
			return  errors.NewNotFound(fmt.Sprintf("Paging Data user Error : %s ",err.Error()))
		}
		res = append(res, u)
	}
	//user = res
	log.Println(res)
	return nil
}


func (user *User) Save() *errors.RestErr{
	stm, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil{
		return  errors.NewInternalServerError(err.Error())
	}
	defer stm.Close() //defer untuk mengecek terkoneksi atau tidak
	user.CreateDate =  date_utils.GetNowLocalString()
	insertResult, err := stm.Exec(user.FirstName,user.LastName,user.Email,user.CreateDate)
	if err != nil {
		if strings.Contains(err.Error(), "1062") {
			return errors.NewBedrequest(fmt.Sprintf("Email %s Already register ",user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error when trying Save user %s ", err.Error()))
	}
	userId, err := insertResult.LastInsertId()
	if err != nil{
		return  errors.NewInternalServerError("Error when trying to get last user id")
	}
	user.Id=userId
	return nil
}