package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"github.com/talkanbaev-artur/auca-numericals-template/src/backend/config"
	"github.com/talkanbaev-artur/auca-numericals-template/src/backend/logic"
	"github.com/talkanbaev-artur/auca-numericals-template/src/backend/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	lg := logrus.New()
	config.Init()

	//REPLACE BY REAL SERVICE
	s := logic.NewAPIService()

	r := mux.NewRouter()
	server.MakeMuxRoutes(s, r, lg)
	go func() {
		c := cors.New(cors.Options{AllowedOrigins: []string{"*"}, AllowedMethods: []string{"POST", "GET", "OPTIONS"}})
		srv := http.Server{
			Handler:      c.Handler(r),
			ReadTimeout:  time.Minute,
			WriteTimeout: time.Minute,
			Addr:         fmt.Sprintf("0.0.0.0:%d", config.Config().Port),
		}
		lg.WithField("port", config.Config().Port).Info("The web server is up")
		lg.Error(srv.ListenAndServe().Error())
		cancel()
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		lg.Info("Catched shutdown signal - initiating graceful shutdown")
		cancel()
	}()

	lg.Info("Initialised numericals application")
	<-ctx.Done()
	lg.Info("Shutting down...")
}
