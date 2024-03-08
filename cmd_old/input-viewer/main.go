package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ttudrej/pokertrainer/pkg/httpd/handler"
)

// ############################################################

// ############################################################
// ############################################################
// ############################################################

func main() {

	fmt.Println("Hello Wrorld")

	r := gin.Default()

	r.POST("/handstate", handler.HandStatePost())

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
