//
//  main.go
//  GinBlog
//
//  Created by Inso on 2018/10/31.
//  Copyright © 2018 Inso. All rights reserved.
//
package main

import (
	"GinBlog/Service/pkg/setting"
	"GinBlog/Service/routers"
	"fmt"
	"net/http"
)

func main() {
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
