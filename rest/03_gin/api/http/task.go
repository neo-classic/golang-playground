package http

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neo-classic/golang-playground/rest/03_gin/domain"
)

// createTask godoc
// @Summary Creates user
// @Produce json
// @Success 200 {object} taskReply
// @Router /task [post]
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

// getAllTasks godoc
// @Summary Returns all tasks
// @Produce json
// @Success 200 {array} taskReply
// @Router /task [get]
func (h *TaskHTTP) getAllTasks(c *gin.Context) {
	ctx := context.Background()

	tasks, err := h.service.Fetch(ctx)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, mapDomainsToReply(tasks))
}

// getAllTasks godoc
// @Summary Delete all tasks
// @Router /task [delete]
func (h *TaskHTTP) deleteAllTasks(c *gin.Context) {
	ctx := context.Background()
	err := h.service.DeleteAll(ctx)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

// deleteTask godoc
// @Summary Delete specified task
// @Param guid path string true "The GUID of the task"
// @Router /task/{guid} [delete]
func (h *TaskHTTP) deleteTask(c *gin.Context) {
	ctx := context.Background()

	guid := c.Params.ByName("guid")
	err := h.service.Delete(ctx, domain.TaskGUID(guid))
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	}
}

// getTask godoc
// @Summary Returns a specific task
// @Produce json
// @Param guid path string true "The GUID of the task"
// @Success 200 {object} taskReply
// @Router /task/{guid} [get]
func (h *TaskHTTP) getTask(c *gin.Context) {
	ctx := context.Background()

	guid := c.Params.ByName("guid")
	task, err := h.service.GetByGUID(ctx, domain.TaskGUID(guid))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, mapDomainToReply(task))
}
