package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type User struct{}

//login
func (u *User) Login(c *gin.Context) {
	fmt.Println("ni hao")
	return
}
