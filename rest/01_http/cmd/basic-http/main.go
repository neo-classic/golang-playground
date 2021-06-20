package main

import (
	"log"

	taskRepository "github.com/neo-classic/golang-playground/rest/01_http/adapters/repository/task"
	http2 "github.com/neo-classic/golang-playground/rest/01_http/api/http"
	cfg "github.com/neo-classic/golang-playground/rest/01_http/config"
	"github.com/neo-classic/golang-playground/rest/01_http/internal/task"
	"gopkg.in/go-playground/validator.v9"
)

func main() {
	config, err := cfg.GetConfigFromFile("config.yml")
	if err != nil {
		log.Fatalf("couldn't load config: %s", err)
		return
	}
	log.Printf("[APP] loaded with config: %+v", config)

	taskRepo, err := taskRepository.NewTaskRepository(config.Mongo.Host, config.Mongo.Database, config.Mongo.User, config.Mongo.Password, "task")
	if err != nil {
		log.Fatalf("couldn't get task repository: %s", err)
		return
	}
	log.Printf("[APP] task repo loaded")

	taskService := task.NewTaskService(taskRepo)
	log.Printf("[APP] task service loaded")

	v := validator.New()
	http2.NewTaskHTTP(taskService, v, config)
}
