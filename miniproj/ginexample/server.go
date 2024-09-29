package ginexample

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TestObj struct {
	Name string `json:"name"`
}

var testError = errors.New("test error")

func GinServer() {
	router := gin.Default()
	router.GET("/authget", testHandler).Use(authHandler)
	router.GET("/test2", testHandler)
	router.GET("/hello", testHandler)
	router.DELETE("/testdelete", deleteHandler).GET("/testget", testHandler)
	err := router.Run(":4040")
	if err != nil {
		return
	}
}

func deleteHandler(ctx *gin.Context) {
	t := TestObj{
		Name: "test name",
	}
	ctx.XML(http.StatusAccepted, t)
}

func testHandler(ctx *gin.Context) {
	t := TestObj{
		Name: "test name",
	}
	ctx.XML(http.StatusOK, t)
}

func authHandler(ctx *gin.Context) {
	fmt.Println("auth")
}
