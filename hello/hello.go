package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func getMessage(c *gin.Context) {
	hostname, _ := os.Hostname()
	c.String(http.StatusOK, fmt.Sprintf("Hello World!\nHostname: %s", hostname))
}

func main() {
	router := gin.Default()
	router.GET("/", getMessage)
	router.Run(":8080")
}
