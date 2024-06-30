package ginexample

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	*http.Client
}

func New() *http.Client {
	c := http.DefaultClient
	return c
}

func CallService() {
	fmt.Println("test")
	c := New()
	resp, err := c.Get("http://localhost:4040/test2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Body)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Body : %s", body)
}
