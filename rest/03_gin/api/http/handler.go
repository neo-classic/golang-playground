package http

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/neo-classic/golang-playground/rest/03_gin/config"
	_ "github.com/neo-classic/golang-playground/rest/03_gin/docs/task"
	"github.com/neo-classic/golang-playground/rest/03_gin/domain"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

// @title Task API
// @version 1.0
// @description Gin playground
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func NewTaskHTTP(s TaskService, v *validator.Validate, cfg *config.Config, log Logger) {
	h := &TaskHTTP{
		service:  s,
		validate: v,
		cfg:      cfg,
		log:      log,
	}

	router := gin.Default()

	task := router.Group("/task")
	{
		task.POST("", h.createTask)
		task.GET("", h.getAllTasks)
		task.DELETE("", h.deleteAllTasks)
		task.GET(":guid", h.getTask)
		task.DELETE(":guid", h.deleteTask)
	}

	router.GET("/tag/:tag", h.tagHandler)
	router.GET("/due/:year/:month/:day", h.dueHandler)

	// Swagger
	url := ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", cfg.Server.Port))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run(fmt.Sprintf("localhost:%d", cfg.Server.Port))
}
