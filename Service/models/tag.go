//
//  tag.go
//  Service 
//
//  Created by Inso on 2018/11/2.
//  Copyright © 2018 Inso. All rights reserved.
//
package models

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

