package main

import (
	"IOBound/logging"
	"IOBound/pkg/config"
	"IOBound/pkg/task"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.GetConfig()
	logger, err := logging.NewLogger("logs", logrus.TraceLevel)
	if err != nil {
		logger.Fatal(err)
	}

	tasks := make(map[string]*task.Task)

	logger.Info("Create Service and Handler")
	taskService := task.NewService(tasks)
	taskHandler := task.NewHandler(taskService, logger)
	logger.Info("Create complited")

	
	r := mux.NewRouter()
	logger.Info("Registered handlers")
	taskHandler.Register(r)
	start(r, cfg, logger)
}

func start(router *mux.Router, cfg *config.Config, logger *logging.Logger) {
	var listener net.Listener
	var listenError error

	logger.Info("Init listner")
	listener, listenError = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
	if listenError != nil {
		logger.Fatal(listenError)
	}

	server := &http.Server{
		Handler: router,
		WriteTimeout: 15* time.Second,
		ReadTimeout: 15 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}	
