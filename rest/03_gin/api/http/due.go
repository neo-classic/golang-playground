package http

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *TaskHTTP) dueHandler(c *gin.Context) {
	ctx := context.Background()

	badRequestError := func() {
		c.String(http.StatusBadRequest, "expect /due/<year>/<month>/<day>, got %v", c.FullPath())
	}

	year, err := strconv.Atoi(c.Params.ByName("year"))
	if err != nil {
		badRequestError()
		return
	}
	month, err := strconv.Atoi(c.Params.ByName("month"))
	if err != nil {
		badRequestError()
		return
	}
	day, err := strconv.Atoi(c.Params.ByName("day"))
	if err != nil {
		badRequestError()
		return
	}
	if month < int(time.January) || month > int(time.December) {
		badRequestError()
		return
	}

	tasks, err := h.service.FetchByDueDate(ctx, year, month, day)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, tasks)
}
