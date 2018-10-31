//
//  main.go
//  GinBlog
//
//  Created by Inso on 2018/10/31.
//  Copyright © 2018 Inso. All rights reserved.
//
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8000")
}
