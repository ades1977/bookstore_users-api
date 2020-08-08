package mysql_utils

import (
	"fmt"
	"github.com/ades1977/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

const (
	errorNorows 	= 	"no rows in result set"
)

func ParseError(err error) *errors.RestErr{
	sqlErr, ok := err.(*mysql.MySQLError)
	log.Println(sqlErr)
	if !ok{
		if strings.Contains(err.Error() ,  errorNorows){
			return errors.NewNotFound(fmt.Sprintf("Record not found  %s",err.Error()))
		}
		return  errors.NewInternalServerError("Error Parsing Database response ")
	}

	switch sqlErr.Number {
	case 1062:
		return  errors.NewBedrequest("Invalid Data / Already Exist ")
	}
	return  errors.NewInternalServerError("Error processing request ")
}