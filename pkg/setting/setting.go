//
//  setting.go
//  GinBlog 
//
//  Created by Inso on 2018/10/31.
//  Copyright © 2018 Inso. All rights reserved.
//

package setting

import (
	"github.com/go-ini/ini"
	"time"
	"log"
)

var (
	//Cfg *ini.File
	Cfg *ini.File

	RunMode string
	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration

	PageSize int
	JwtSecret string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini' : %v", err)
	}
}
