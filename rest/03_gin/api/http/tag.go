package http

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *TaskHTTP) tagHandler(c *gin.Context) {
	ctx := context.Background()

	tag := c.Params.ByName("tag")

	tasks, err := h.service.FetchByTag(ctx, tag)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, mapDomainsToReply(tasks))
}
