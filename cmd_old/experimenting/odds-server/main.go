package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/ttudrej/hand_analysis/httpd/handler"
	"github.com/ttudrej/hand_analysis/platform/newsfeed"
)

// ############################################################
func GetHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GET Home Page",
	})
}

// ############################################################
func PostHomePage(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(200, gin.H{
		"message": string(value),
	})
}

// ############################################################
func PutHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "PUT home page",
	})
}

// ############################################################
func DeleteHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "DELETE home page",
	})
}

// ############################################################
func PatchHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "PATCH home page",
	})
}

// ############################################################
func HeadHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Head home page",
	})
}

// ############################################################
func OptionsHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OPTIONS home page",
	})
}

// ############################################################
func QueryStrings(c *gin.Context) {
	name := c.Query("name") // name=<somestring>
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

// ############################################################
func PathParams(c *gin.Context) {
	name := c.Param("name") // name=<somestring>
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

// ############################################################
// ############################################################
// ############################################################

func main() {

	fmt.Println("Hello Wrorld")
	feed := newsfeed.New()

	r := gin.Default()

	r.GET("/", GetHomePage)

	r.GET("/ping", handler.PingGet(feed))
	r.GET("/newsfeed", handler.NewsfeedGet(feed))
	r.POST("/newsfeed", handler.NewsfeedPost(feed))
	r.POST("/potodds", handler.PotOddsPost())
	r.POST("/handoddsf", handler.HandOddsFPost())
	// r.POST("/handoddst", handler.HandOddsTPost())
	// r.POST("/handoddsr", handler.HandOddsRPost())

	r.GET("/query", QueryStrings)         // /query?name=<somestriing>&age=<someint>
	r.GET("/path/:name/:age", PathParams) // /path/tom/50

	r.POST("/", PostHomePage)

	r.PUT("/", PutHomePage)
	r.DELETE("/", DeleteHomePage)
	r.PATCH("/", PatchHomePage)
	r.HEAD("/", HeadHomePage)
	r.OPTIONS("/", OptionsHomePage)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	/*
		feed := newsfeed.New()
		fmt.Println(feed)
		feed.Add(newsfeed.Item{"Hello", "how ya doing?"})
		fmt.Println(feed)
	*/
}
