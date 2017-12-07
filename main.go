package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/flag", PrintFlag)
	router.Run(":8080")
}

func PrintFlag(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		s := string(body)
		fmt.Println(s)
	}

	cookie, err := c.Request.Cookie("FLAG")
	if err != nil {
		fmt.Println(cookie.Value)
	}

}
