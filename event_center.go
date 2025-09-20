package wtype

import (
	"reflect"
	"sync"
)

type EventHandler func(data ...any)

type EventCenter struct {
	event map[string][]EventHandler
	lock  sync.RWMutex
}

func (e *EventCenter) On(key string, handler EventHandler) {
	e.lock.Lock()
	defer e.lock.Unlock()
	e.event[key] = append(e.event[key], handler)
}

func (e *EventCenter) Off(key string, handler EventHandler) {
	e.lock.Lock()
	defer e.lock.Unlock()
	if handlers, ok := e.event[key]; ok {
		for i, h := range handlers {
			if reflect.ValueOf(h).Pointer() == reflect.ValueOf(handler).Pointer() {
				e.event[key] = append(handlers[:i], handlers[i+1:]...)
				break
			}
		}
	}
}

func (e *EventCenter) Emit(key string, data ...any) {
	e.lock.RLock()
	handlers, ok := e.event[key]
	if !ok {
		e.lock.RUnlock()
		return
	}
	handlersCopy := make([]EventHandler, len(handlers))
	copy(handlersCopy, handlers)
	e.lock.RUnlock()

	for _, handler := range handlersCopy {
		handler(data...)
	}
}

func NewEventCenter() *EventCenter {
	return &EventCenter{
		event: make(map[string][]EventHandler),
	}
}
