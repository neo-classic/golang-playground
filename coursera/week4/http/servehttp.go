package main

import (
	"fmt"
	"net/http"
)

type Handler struct {
	Name string
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "NAME:", h.Name, "\nURL:", r.URL.String())
}

func main() {
	testHandler := &Handler{Name: "test"}
	http.Handle("/test/", testHandler)

	rootHandler := &Handler{Name: "root"}
	http.Handle("/", rootHandler)

	adminHandler := &Handler{Name: "admin"}
	http.Handle("/admin/", adminHandler)

	fmt.Println("starting server at :8088")
	http.ListenAndServe(":8088", nil)
}
