package domain

import (
	"time"

	"github.com/pborman/uuid"
)

type Task struct {
	GUID TaskGUID
	Text string
	Tags []string
	Due  time.Time
}

type TaskGUID string

func (g TaskGUID) String() string {
	return string(g)
}

func NewTaskGUID() TaskGUID {
	return TaskGUID(uuid.New())
}
