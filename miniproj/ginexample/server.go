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
	router.Use(authHandler)
	router.GET("/test2", testHandler)
	err := router.Run(":4040")
	if err != nil {
		return
	}
	fmt.Println("calling service")
	CallService()
	fmt.Println("end service")
}

func testHandler(ctx *gin.Context) {
	t := TestObj{
		Name: "test name",
	}

	ctx.JSON(http.StatusOK, t)

}

func authHandler(ctx *gin.Context) {
	fmt.Println("auth")
}
