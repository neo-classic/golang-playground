package http

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func (h *TaskHTTP) dueHandler(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	h.log.Info(fmt.Sprintf("handling tasks by due at %s\n", req.URL.Path))

	vars := mux.Vars(req)
	badRequestError := func() {
		http.Error(w, fmt.Sprintf("expect /due/<year>/<month>/<day>, got %v", req.URL.Path), http.StatusBadRequest)
	}

	year, _ := strconv.Atoi(vars["year"])
	month, _ := strconv.Atoi(vars["month"])
	day, _ := strconv.Atoi(vars["day"])
	if month < int(time.January) || month > int(time.December) {
		badRequestError()
		return
	}

	tasks, err := h.service.FetchByDueDate(ctx, year, month, day)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	renderJSON(w, tasks)
}
