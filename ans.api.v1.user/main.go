package main

import (
	_ "answersys/db"

	"answersys/ans.api.v1.user/controllers"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-web"
)

func main() {
	service := web.NewService(
		web.Name("com.judashi.api.v1.user"),
	)
	router := gin.Default()
	routerGroup := router.Group("/v1/user")
	user := &controllers.User{}

	routerGroup.POST("login", user.Login)

	service.Handle("/", router)
	err := service.Run()
	if err != nil {
		panic(err)
	}
}
