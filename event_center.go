package wtype

import (
	"reflect"
	"sync"
)

// EventHandler defines a function type for event callbacks.
// The function can accept any number of arguments of any type.
type EventHandler func(data ...any)

// EventCenter manages event registration, emission, and removal.
// It is safe for concurrent use.
type EventCenter struct {
	event map[string][]EventHandler // Map of event keys to their handlers
	lock  sync.RWMutex              // Read/Write lock to ensure concurrency safety
}

// On registers a new handler for the given event key.
// Multiple handlers can be registered under the same key.
func (e *EventCenter) On(key string, handler EventHandler) {
	e.lock.Lock()
	defer e.lock.Unlock()
	e.event[key] = append(e.event[key], handler)
}

// Once registers a handler that will only be executed once for the given event key.
// After the first execution, the handler is automatically removed.
func (e *EventCenter) Once(key string, handler EventHandler) {
	e.lock.Lock()
	defer e.lock.Unlock()

	// Create a wrapper handler that removes itself after execution
	var wrappedHandler EventHandler
	wrappedHandler = func(data ...any) {
		// Execute the original handler
		handler(data...)

		// Remove this handler after execution
		e.Off(key, wrappedHandler)
	}

	e.event[key] = append(e.event[key], wrappedHandler)
}

// Off removes a specific handler for the given event key.
// It compares handlers by their underlying function pointer value.
func (e *EventCenter) Off(key string, handler EventHandler) {
	e.lock.Lock()
	defer e.lock.Unlock()
	if handlers, ok := e.event[key]; ok {
		for i, h := range handlers {
			// Compare the function pointer addresses to identify the handler
			if reflect.ValueOf(h).Pointer() == reflect.ValueOf(handler).Pointer() {
				e.event[key] = append(handlers[:i], handlers[i+1:]...)
				break
			}
		}
	}
}

// Emit triggers all handlers registered for the given event key.
// The data arguments are passed to each handler in order.
// Handlers are copied before invocation to avoid issues if the list is modified during execution.
func (e *EventCenter) Emit(key string, data ...any) {
	e.lock.RLock()
	handlers, ok := e.event[key]
	if !ok {
		e.lock.RUnlock()
		return
	}

	// Make a copy of the handlers slice to avoid race conditions
	handlersCopy := make([]EventHandler, len(handlers))
	copy(handlersCopy, handlers)
	e.lock.RUnlock()

	// Execute each handler with the provided data
	for _, handler := range handlersCopy {
		handler(data...)
	}
}

// NewEventCenter creates and returns a new EventCenter instance.
func NewEventCenter() *EventCenter {
	return &EventCenter{
		event: make(map[string][]EventHandler),
	}
}
