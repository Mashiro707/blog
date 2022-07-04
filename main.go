package main

import (
	"blog/config"
	"blog/internal/router"
	"blog/pkg/mysql"
	"blog/pkg/redis"
	"blog/pkg/zap"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	zap.InitLogger()
	var (
		mysqlClient mysql.Mysql
		redisClient redis.Redis
	)
	// init router
	r := router.InitRouter()
	// mysql connect
	mysqlClient.Setup()
	// redis connect
	redisClient.Setup()
	// open port listen
	app := config.ApplicationConfig
	srv := &http.Server{
		Addr:         app.Host + ":" + app.Port,
		Handler:      r,
		ReadTimeout:  time.Duration(app.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(app.WriterTimeout) * time.Second,
	}
	go func() {
		zap.InfoLog("server start")
		if err := srv.ListenAndServe(); err != nil {
			zap.ErrorLog(err)
			panic("open http error")
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
	zap.InfoLog("serve close success")
}
