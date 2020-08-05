package users

import (
	"fmt"
	"github.com/ades1977/bookstore_users-api/domain/users"
	"github.com/ades1977/bookstore_users-api/services"
	"github.com/ades1977/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)



func CreateUser(c *gin.Context) {
	var user users.User
	if err:= c.ShouldBindJSON(&user); err !=nil {
		//return invalid recall first time ( bed Request )
		 restErr := errors.NewBedrequest("Salah Nih : Bed Request JSON")
		c.JSON(restErr.Status,restErr)
		fmt.Println(err)
		return
	}
	result,  saveErr  := services.CreateUser(user)
	if saveErr !=nil{
		//hubungan error ke database ( not found, down, etc )
		c.JSON(saveErr.Status,saveErr)
		return
	}

	//result = errors.NewSaveDBSuccess(fmt.Sprintf("Save database users success %s ", user))
	//fmt.Println(user)

	var response users.UserResponse
	response.Message="Save User Success"
	response.Status=200
	response.Data.Id= result.Id
	response.Data.FirstName= result.FirstName
	response.Data.LastName = result.LastName
	response.Data.Email = result.Email
	response.Data.CreateDate = result.CreateDate
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"),10,64)
	if userErr != nil {
		err := errors.NewBedrequest("Userid should be a number ")
		c.JSON(err.Status,err)
		return
	}
	result, getErr  := services.GetUser(userId)
	if getErr != nil{
		//hubungan error ke database ( not found, down, etc )
		c.JSON(getErr.Status,getErr)
		return
	}

	var response users.UserResponse
	response.Message="Retrive User Success"
	response.Status=200
	response.Data .Id= result.Id
	response.Data.FirstName= result.FirstName
	response.Data.LastName = result.LastName
	response.Data.Email = result.Email
	response.Data.CreateDate = result.CreateDate
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}


func GetUserPaging(c *gin.Context) {
	p1, userErr := strconv.ParseInt(c.Param("page_from"),10,64)
	if userErr != nil {
		err := errors.NewBedrequest("Userid should be a number ")
		c.JSON(err.Status,err)
		return
	}
	p2, userErr := strconv.ParseInt(c.Param("page_to"),10,64)
	if userErr != nil {
		err := errors.NewBedrequest("Userid should be a number ")
		c.JSON(err.Status,err)
		return
	}
	result , getErr := services.GetUserPaging(p1,p2)
	if getErr != nil{
		//hubungan error ke database ( not found, down, etc )
		c.JSON(getErr.Status,getErr)
		return
	}

	var response users.UserResponse
	response.Message="Retrive User Success"
	response.Status=200
	//response.Data = result
	log.Println(result)
	//response.Data =  result
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}

