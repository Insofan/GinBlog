//
//  models.go
//  Service
//
//  Created by Inso on 2018/11/2.
//  Copyright © 2018 Inso. All rights reserved.
//
package models

import (
	"GinBlog/Service/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

type Model struct {
	Id          int `gorm:"primary_key" json:"id"`
	CreateOn    int `json:"create_on"`
	ModifiedOne int `json:"modified_one"`
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatalf(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("Type").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("user").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbName))

	if err != nil {
		log.Println(err)
	}

	//这里用了table prefix
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	//全局设置表名不可以为复数形式。
	db.SingularTable(true)

	//SetMaxOpenConns用于设置最大打开的连接数
	//SetMaxIdleConns用于设置闲置的连接数
	db.DB().SetMaxIdleConns(100)
	db.DB().SetMaxOpenConns(100)

}

func CloseDB() {
	defer db.Close()
}
