package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mang022/cafe/conf"
	"github.com/mang022/cafe/db"
)

func main() {
	conf.SetupConfig()
	log.Println("컨피그 파일을 설정하였습니다.")

	db.SetupDB()
	log.Println("DB를 설정하였습니다.")

	router := setupRouter()
	log.Println("Router를 설정하였습니다.")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	log.Println("서버를 실행합니다.")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	db.CloseDB()
	log.Println("DB 연결을 종료합니다.")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		panic(err)
	}
	log.Println("서버를 종료합니다")
}
