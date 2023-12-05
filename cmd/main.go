package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"gitlab.com/ast3am77/test-go/api/handlers"
	"gitlab.com/ast3am77/test-go/internal/config"
	"gitlab.com/ast3am77/test-go/internal/db"
	"gitlab.com/ast3am77/test-go/internal/mailSender"
	"gitlab.com/ast3am77/test-go/internal/service"
	"gitlab.com/ast3am77/test-go/pkg/logging"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cfg := config.GetConfig("config/config.yml")
	log := logging.GetLogger(cfg.LogLevel)
	testDB, err := db.NewClient(ctx, cfg, log)
	if err != nil {
		log.FatalMsg("", err)
	}
	defer testDB.Close(ctx)

	emailDealer, err := mailSender.NewSender(cfg, log)
	if err != nil {
		log.FatalMsg("", err)
	}

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	servise := service.NewService(testDB, log, emailDealer)
	handler := handlers.NewHandler(servise, log)
	handler.RegisterHandlers(router)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.FatalMsg("listen: %s\n", err)
		}
	}()

	log.InfoMsg("service is running")
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
