package main

import (
	"context"
	"fmt"

	"github.com/neo-classic/golang-playground/rest/01_http/adapters/logger"
	taskRepository "github.com/neo-classic/golang-playground/rest/01_http/adapters/repository/task"
	http2 "github.com/neo-classic/golang-playground/rest/01_http/api/http"
	cfg "github.com/neo-classic/golang-playground/rest/01_http/config"
	"github.com/neo-classic/golang-playground/rest/01_http/internal/task"
	"gopkg.in/go-playground/validator.v9"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log := logger.Logger{}

	config, err := cfg.GetConfigFromFile("config.yml")
	if err != nil {
		log.Error(ctx, fmt.Sprintf("couldn't load config: %s", err))
		return
	}
	log.Info(ctx, fmt.Sprintf("[APP] loaded with config: %+v", config))

	taskRepo, err := taskRepository.NewTaskRepository(config.Mongo.Host, config.Mongo.Database, config.Mongo.User, config.Mongo.Password, "task")
	if err != nil {
		log.Error(ctx, fmt.Sprintf("couldn't get task repository: %s", err))
		return
	}
	log.Info(ctx, "[APP] task repo loaded")

	taskService := task.NewTaskService(taskRepo)
	log.Info(ctx, "[APP] task service loaded")

	v := validator.New()
	http2.NewTaskHTTP(ctx, taskService, v, config, log)
}
