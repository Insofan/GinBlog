//
//  tag.go
//  Service 
//
//  Created by Inso on 2018/11/2.
//  Copyright © 2018 Inso. All rights reserved.
//
package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"
	"GinBlog/Service/pkg/e"
	"GinBlog/Service/models"
	"GinBlog/Service/pkg/util"
	"GinBlog/Service/pkg/setting"
	"net/http"
	"github.com/astaxie/beego/validation"
	"fmt"
)

func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")

	valid := validation.Validation{}
	//用来校验是否合法
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTagByName(name)	 {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			//logging.Info(err.Key, err.Message)
			fmt.Println("Info: key " + err.Key+ " msg "+ err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg" : e.GetMsg(code),
		"data": make(map[string]string),
	})
}
func EditTag(c *gin.Context) {

}
func DeleteTag(c *gin.Context) {

}
