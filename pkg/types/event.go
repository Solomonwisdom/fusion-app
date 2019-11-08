package types

type EventSourceType string

const (
	InternalEventHub       EventSourceType = "eventHub"
	ExternalWorkflowEngine EventSourceType = "workflowEngine"
	ExternalService        EventSourceType = "service"
)

type EventSource struct {
	Component EventSourceType
	Host      string
}

type Event struct {
	ID     string
	Source EventSource
}
