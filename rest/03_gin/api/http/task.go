package http

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neo-classic/golang-playground/rest/03_gin/domain"
)

func (h *TaskHTTP) createTask(c *gin.Context) {
	ctx := context.Background()
	input := createTaskRequest{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	if err := h.validate.Struct(input); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	t := mapCreateRequestToDomain(&input)
	task, err := h.service.Create(ctx, *t)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, mapDomainToReply(task))
}

func (h *TaskHTTP) getAllTasks(c *gin.Context) {
	ctx := context.Background()

	tasks, err := h.service.Fetch(ctx)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, mapDomainsToReply(tasks))
}

func (h *TaskHTTP) deleteAllTasks(c *gin.Context) {
	ctx := context.Background()
	err := h.service.DeleteAll(ctx)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func (h *TaskHTTP) deleteTask(c *gin.Context) {
	ctx := context.Background()

	guid := c.Params.ByName("guid")
	err := h.service.Delete(ctx, domain.TaskGUID(guid))
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	}
}

func (h *TaskHTTP) getTask(c *gin.Context) {
	ctx := context.Background()

	guid := c.Params.ByName("guid")
	task, err := h.service.GetByGUID(ctx, domain.TaskGUID(guid))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, mapDomainToReply(task))
}
