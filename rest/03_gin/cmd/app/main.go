package main

import (
	"fmt"

	"github.com/neo-classic/golang-playground/rest/03_gin/adapters/logger"
	taskRepository "github.com/neo-classic/golang-playground/rest/03_gin/adapters/repository/task"
	http2 "github.com/neo-classic/golang-playground/rest/03_gin/api/http"
	cfg "github.com/neo-classic/golang-playground/rest/03_gin/config"
	"github.com/neo-classic/golang-playground/rest/03_gin/internal/task"
	"gopkg.in/go-playground/validator.v9"
)

func main() {
	log := logger.NewLogger(logger.Dev)

	config, err := cfg.GetConfigFromFile("config.yml")
	if err != nil {
		log.Error(fmt.Sprintf("couldn't load config: %s", err))
		return
	}
	log.Info(fmt.Sprintf("[APP] loaded with config: %+v", config))

	taskRepo, err := taskRepository.NewTaskRepository(config.Mongo.Host, config.Mongo.Database, config.Mongo.User, config.Mongo.Password, "task")
	if err != nil {
		log.Error(fmt.Sprintf("couldn't get task repository: %s", err))
		return
	}
	log.Info("[APP] task repo loaded")

	taskService := task.NewTaskService(taskRepo)
	log.Info("[APP] task service loaded")

	v := validator.New()
	http2.NewTaskHTTP(taskService, v, config, log)
}
