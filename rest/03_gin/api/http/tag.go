package http

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// tagHandler godoc
// @Summary Returns tasks by provided tag
// @Produce json
// @Success 200 {array} taskReply
// @Param tag path string true "The tag belongs to tasks"
// @Router /tag/{tag} [get]
func (h *TaskHTTP) tagHandler(c *gin.Context) {
	ctx := context.Background()

	tag := c.Params.ByName("tag")

	tasks, err := h.service.FetchByTag(ctx, tag)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, mapDomainsToReply(tasks))
}
