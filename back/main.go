//
//  main.go
//  GinBlog
//
//  Created by Inso on 2018/10/31.
//  Copyright © 2018 Inso. All rights reserved.
//
package main

import (
	"GinBlog/Service/models"
	"GinBlog/Service/pkg/setting"
	"GinBlog/Service/routers"
	"context"
	"fmt"
	"github.com/robfig/cron"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	log.Println("Starting...")
	c := cron.New()
	c.AddFunc("* * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})
	c.AddFunc("* * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})
	c.Start()

	t1 := time.NewTimer(time.Second * 100)

	for {
		select {
			case <-t1.C:
				t1.Reset(time.Second * 100)
		}
	}

	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	//go chan
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<- quit

	log.Println("ShutDown Server ...")

	ctx , cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
	log.Println("Server exiting")
}
