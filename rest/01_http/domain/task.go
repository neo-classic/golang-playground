package domain

import (
	"time"

	"github.com/pborman/uuid"
)

type Task struct {
	GUID TaskGUID  `json:"guid"`
	Text string    `json:"text"`
	Tags []string  `json:"tags"`
	Due  time.Time `json:"due"`
}

type TaskGUID string

func (g TaskGUID) String() string {
	return string(g)
}

func NewTaskGUID() TaskGUID {
	return TaskGUID(uuid.New())
}
