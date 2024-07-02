package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// making struct json
type api struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
var data api


func main() {
	r := gin.Default()
	r.GET("/get", getValues)
	r.POST("/post", postValues)
	r.PUT("/put", putValues)
	r.DELETE("/delete", deleteValues)
	r.Run(":9080") // listen and serve on 0.0.0.0:8080
}

// 1:getValues function
func getValues(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": data,
	})
}

// 2: postValues Function
func postValues(c *gin.Context) {
	err := c.BindJSON(&data) // here we bind the data
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Message ": "Something  Wrong",
		})
		return 

	} // StatusOK('k' is capital of OK )
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

// update by using the Put
func putValues(c *gin.Context) {
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Message":"Something wrong " ,
		})
		return
	}
	c.JSON(200,gin.H{
		"message":data,
	})

}

// Now delete method
func deleteValues(c *gin.Context) {
	//It clear the data 
	data=api{}
	
	c.JSON(200,gin.H{
		"message":data,
	})

}