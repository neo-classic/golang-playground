package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (h *TaskHTTP) tagHandler(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	h.log.Info(ctx, fmt.Sprintf("handling tasks by tag at %s\n", req.URL.Path))

	if req.Method != http.MethodGet {
		http.Error(w, fmt.Sprintf("expect method GET /tag/<tag>, got %v", req.Method), http.StatusMethodNotAllowed)
		return
	}

	path := strings.Trim(req.URL.Path, "/")
	pathParts := strings.Split(path, "/")
	if len(pathParts) < 2 {
		http.Error(w, "expect /tag/<tag> path", http.StatusBadRequest)
		return
	}
	tag := pathParts[1]

	tasks, err := h.service.FetchByTag(ctx, tag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(mapDomainsToReply(tasks))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
