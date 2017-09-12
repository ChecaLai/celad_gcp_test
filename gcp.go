package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	//web宣告
	router := InitRouter()
	srv := &http.Server{
		Addr:  	 ":80",
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	cancel()
	log.Println("Server exit")
	return
}

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	//router := gin.New()
	router := gin.Default()
	router.GET("/", index)
	return router
}
