package users_db

import (
	"database/sql"
	"fmt"
	"log"
	_  "github.com/go-sql-driver/mysql"  //pergunakan _ jika tidak belum ada referensinya
	"os"
)

const (
	mysql_users_username="mysql_users_username"
	mysql_users_password="mysql_users_password"
	mysql_users_host="mysql_users_host"
	mysql_users_schema="mysql_users_schema"
)


var(
	Client *sql.DB

	username = os.Getenv("mysql_users_username")
	password = os.Getenv("mysql_users_password")
	host 	 = os.Getenv("mysql_users_host")
	schema 	 = os.Getenv("mysql_users_schema")

)

func init(){

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username,//userid
		password,	//password
		host, //host
		schema,	//nama database
		)

	var err error
	Client, err = sql.Open("mysql",dataSourceName)

	//don't log database connected in production environment
	//log.Println(fmt.Sprintf("About to Connect %s",dataSourceName))
	if err != nil {
		panic(err)
	}
	if err = Client.Ping() ; err !=nil{
		panic(err)
	}
	log.Println("Database connected succefuly")

}
