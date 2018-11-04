//
//  router.go
//  Service
//
//  Created by Inso on 2018/11/2.
//  Copyright © 2018 Inso. All rights reserved.
//
package routers

import (
	"GinBlog/Service/midware/jwt"
	"GinBlog/Service/pkg/setting"
	"GinBlog/Service/routers/api"
	"GinBlog/Service/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")

	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//article
		//获取所有文章列表
		apiv1.GET("/articles", v1.GetAritcles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		apiv1.POST("/articles", v1.AddArticles)
		apiv1.PUT("/articles/:id", v1.EditArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}
	return r
}
