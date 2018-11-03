//
//  tag.go
//  Service 
//
//  Created by Inso on 2018/11/2.
//  Copyright © 2018 Inso. All rights reserved.
//
package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model

	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag)  {
	//从pageNum起, pageSize数据
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int)  {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByName(name string) bool  {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.Id > 0 {
		return true
	}
	return false
}

func AddTag(name string, state int, createdBy string) bool  {

	db.Create(&Tag{
		Name: name,
		State: state,
		CreatedBy: createdBy,
	})
	return true
}

func EditTag(c *gin.Context)  {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Required(id, "id").Message("ID不能为空")
}

//这是gorm自带的callback, gorm支持的回调方法
/*
	创建：BeforeSave、BeforeCreate、AfterCreate、AfterSave
	更新：BeforeSave、BeforeUpdate、AfterUpdate、AfterSave
	删除：BeforeDelete、AfterDelete
	查询：AfterFind
 */
func (tag *Tag)BeforeCreate(scope *gorm.Scope) error  {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (tag *Tag)BeforeUpdate(scope *gorm.Scope) error  {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

