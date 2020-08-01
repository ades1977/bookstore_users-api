package users

import (
	"fmt"
	"github.com/ades1977/bookstore_users-api/domain/users"
	"github.com/ades1977/bookstore_users-api/services"
	"github.com/ades1977/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
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
	result, saveErr  := services.CreateUser(user)
	if saveErr !=nil{
		//hubungan error ke database ( not found, down, etc )
		c.JSON(saveErr.Status,saveErr)
		return
	}
	fmt.Println(user)
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"),10,64)
	if userErr != nil {
		err := errors.NewBedrequest("Userid should be a number ")
		c.JSON(err.Status,err)
		return
	}
	user, getErr  := services.GetUser(userId)
	if getErr != nil{
		//hubungan error ke database ( not found, down, etc )
		c.JSON(getErr.Status,getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}
