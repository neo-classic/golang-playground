package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *TaskHTTP) tagHandler(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	h.log.Info(fmt.Sprintf("handling tasks by tag at %s\n", req.URL.Path))

	tag := mux.Vars(req)["tag"]

	tasks, err := h.service.FetchByTag(ctx, tag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	renderJSON(w, mapDomainsToReply(tasks))
}
