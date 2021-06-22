package http

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/neo-classic/golang-playground/rest/03_gin/config"
	"github.com/neo-classic/golang-playground/rest/03_gin/domain"
	"gopkg.in/go-playground/validator.v9"
)

type TaskService interface {
	Create(ctx context.Context, task domain.Task) (*domain.Task, error)
	Fetch(ctx context.Context) ([]*domain.Task, error)
	FetchByTag(ctx context.Context, tag string) ([]*domain.Task, error)
	FetchByDueDate(ctx context.Context, y, m, d int) ([]*domain.Task, error)
	GetByGUID(ctx context.Context, guid domain.TaskGUID) (*domain.Task, error)
	Delete(ctx context.Context, guid domain.TaskGUID) error
	DeleteAll(ctx context.Context) error
}

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Error(msg string)
}

type TaskHTTP struct {
	service  TaskService
	validate *validator.Validate
	cfg      *config.Config
	log      Logger
}

func NewTaskHTTP(s TaskService, v *validator.Validate, cfg *config.Config, log Logger) {
	h := &TaskHTTP{
		service:  s,
		validate: v,
		cfg:      cfg,
		log:      log,
	}

	router := gin.Default()

	router.POST("/task/", h.createTask)
	router.GET("/task/", h.getAllTasks)
	router.DELETE("/task/", h.deleteAllTasks)
	router.GET("/task/:guid", h.getTask)
	router.DELETE("/task/:guid", h.deleteTask)

	router.GET("/tag/:tag", h.tagHandler)
	router.GET("/due/:year/:month/:day", h.dueHandler)

	router.Run(fmt.Sprintf("localhost:%d", cfg.Server.Port))
}
