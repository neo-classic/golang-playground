package http

import (
	"time"

	"github.com/neo-classic/golang-playground/rest/01_http/domain"
)

type createTaskRequest struct {
	GUID string    `json:"guid"`
	Text string    `json:"text" validate:"required"`
	Tags []string  `json:"tags"`
	Due  time.Time `json:"due"`
}

type taskReply struct {
	GUID string    `json:"guid"`
	Text string    `json:"text"`
	Tags []string  `json:"tags"`
	Due  time.Time `json:"due"`
}

func mapCreateRequestToDomain(in *createTaskRequest) *domain.Task {
	return &domain.Task{
		GUID: domain.TaskGUID(in.GUID),
		Text: in.Text,
		Tags: in.Tags,
		Due:  in.Due,
	}
}

func mapDomainToReply(in *domain.Task) *taskReply {
	return &taskReply{
		GUID: string(in.GUID),
		Text: in.Text,
		Tags: in.Tags,
		Due:  in.Due,
	}
}

func mapDomainsToReply(in []*domain.Task) []*taskReply {
	var tasks []*taskReply
	for _, t := range in {
		tasks = append(tasks, mapDomainToReply(t))
	}
	return tasks
}
