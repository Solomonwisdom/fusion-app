package mqhub

import (
	core "github.com/lxs137/fusion-app/pkg/types"
)

// SubscriberInterface defines how to receive Event
type SubscriberInterface interface {
	Init(config interface{}) error
	SubscribeTo(source *core.EventSource) (<-chan interface{}, error)
	ListResourceEvents(query interface{}) ([]interface{}, error)
	ListAppEvents(query interface{}) ([]interface{}, error)
}
