package users

import (
	"github.com/ades1977/bookstore_users-api/databases/mysql/users_db"
	"github.com/ades1977/bookstore_users-api/utils/date_utils"
	"github.com/ades1977/bookstore_users-api/utils/errors"
	"github.com/ades1977/bookstore_users-api/utils/mysql_utils"
)

const(
	queryInsertUser = 	"INSERT INTO users(first_name,last_name,email,password,status, create_date) values(?,?,?,?,?,?);"
	querySelectUser	= 	"SELECT id,first_name,last_name,email, create_date FROM users WHERE id=? limit 1 ;"
	queryPagingUser	=	"select id,first_name,last_name,email, create_date from users order by id desc  limit ?,? ;"
	queryUpdateUser	=	"UPDATE USERS set first_name=?,last_name=?,email=? where id=?"
	queryDeleteUser	=	"DELETE FROM USERS where id=?"
	queryFindUserByStatus	= 	"SELECT id,first_name,last_name,email, create_date,status FROM users WHERE status=? ;"
)


var (
	userDB=make(map[int64]*User)
)

func (user *User) Get(userId int64) ([]User, *errors.RestErr){
	stm, err := users_db.Client.Prepare(querySelectUser)
	if err != nil{
		return  nil, errors.NewInternalServerError(err.Error())
	}
	defer stm.Close()
	var zuser User
	rows := stm.QueryRow(userId)  //untuk single row
	if err := rows.Scan(&zuser.Id, &zuser.FirstName, &zuser.LastName, &zuser.Email, &zuser.CreateDate);
		err != nil {
		return nil,mysql_utils.ParseError(err)
	}
	results := make([]User , 0)
	results = append(results, zuser)
	return results,nil
}

func (user *User) FindStatus(status string) ([]User, *errors.RestErr) {
	stm, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil{
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stm.Close()
	rows, err := stm.Query(status)  //untuk single row
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	defer rows.Close()
	results := make([]User , 0)
	for  rows.Next() {
		var user User
		if rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreateDate, &user.Status);
			err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}
	if len(results)==0 {
		return nil, errors.NewNotFound("Data Not Found")
	}
	return results,nil
}



func (user User) GetPaging(p1 int64,p2 int64) ([]User, *errors.RestErr){
	stm, err := users_db.Client.Prepare(queryPagingUser)
	if err != nil{
		return  nil,errors.NewInternalServerError(err.Error())
	}

	defer stm.Close()
	rows, err := stm.Query(p1,p2) //untuk banyak row
	if err != nil{
		return nil, mysql_utils.ParseError(err)
	}
	defer rows.Close()
	//log.Println(queryPagingUser)

	results := make([]User , 0)
	for  rows.Next() {
		//var zuser User
		if rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreateDate );
			err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
		//log.Println(user.FirstName)
	}

	if len(results)==0 {
		return nil, errors.NewNotFound("Data Not Found")
	}

	return results,nil

}


func (user *User) Save() (int64, *errors.RestErr){
	stm, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil{
		return  0,errors.NewInternalServerError(err.Error())
	}
	defer stm.Close() //defer untuk mengecek terkoneksi atau tidak
	user.CreateDate =  date_utils.GetNowLocalString()
	insertResult, saveErr :=  stm.Exec(user.FirstName,user.LastName,user.Email,user.Password, user.Status,user.CreateDate)
	if saveErr != nil {
			return 0,mysql_utils.ParseError(saveErr)
	}
	userId, err := insertResult.LastInsertId()
	if err != nil{
		return  0,errors.NewInternalServerError("Error when trying to get last user id")
	}

	user.Id=userId
	return userId,nil
}

func (user *User) Update() *errors.RestErr{
	stm, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil{
		return  errors.NewInternalServerError(err.Error())
	}
	defer stm.Close() //defer untuk mengecek terkoneksi atau tidak

	_, saveErr :=  stm.Exec(user.FirstName,user.LastName,user.Email,user.Id)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	return nil

}

func (user *User) Delete() *errors.RestErr{
	stm, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil{
		return  errors.NewInternalServerError(err.Error())
	}
	defer stm.Close() //defer untuk mengecek terkoneksi atau tidak

	_, saveErr :=  stm.Exec(user.Id)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	return nil

}