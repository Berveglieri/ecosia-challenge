package server

import (
	"ecosia/pkg/router"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"time"
)

type serverConfig struct {
	port         int
	routes       httprouter.Router
	idleTimeout  string
	readTimeout  string
	writeTimeout string
}

type ApplicationServer interface {
	RunServer() error
}

func NewApplicationServer(port int, idleTimeout string, readTimeout string, writeTimeout string) *serverConfig {
	cfg := &serverConfig{
		port:         port,
		idleTimeout:  idleTimeout,
		readTimeout:  readTimeout,
		writeTimeout: writeTimeout,
	}

	return cfg
}

func (cfg *serverConfig) RunServer() error {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	idle, err := time.ParseDuration(cfg.idleTimeout)
	if err != nil {
		logger.Println(err)
	}

	read, err := time.ParseDuration(cfg.readTimeout)
	if err != nil {
		logger.Println(err)
	}

	write, err := time.ParseDuration(cfg.writeTimeout)
	if err != nil {
		logger.Println(err)
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      router.Routes(),
		IdleTimeout:  idle,
		ReadTimeout:  read,
		WriteTimeout: write,
	}

	logger.Println("Starting server on port", cfg.port)

	err = srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
