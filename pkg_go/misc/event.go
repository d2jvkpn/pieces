package misc

import (
	"time"
)

//// Event
type Event struct {
	Name  string      `json:"name" bson:"name"`                       // event name
	Code  int32       `json:"code" bson:"code"`                       // event code
	At    time.Time   `json:"at" bson:"at"`                           // envet occured at
	Data  interface{} `json:"data,omitempty" bson:"data,omitempty"`   // attach data
	Error Error       `json:"error,omitempty" bson:"error,omitempty"` // event error information
}

type EventOption func(event *Event)

func WithError(err error) EventOption {
	return func(event *Event) { event.Error = NewError(err) }
}

func WithData(data interface{}) EventOption {
	return func(event *Event) { event.Data = data }
}

func NewEvent(name string, code int32, opts ...EventOption) (event *Event) {
	event = &Event{Name: name, Code: code, At: time.Now()}

	for i := range opts {
		opts[i](event)
	}
	return
}

//// Error
type Error struct{ x error }

func NewError(err error) (e Error) { return Error{x: err} }

func (e Error) Error() string {
	if e.x == nil {
		return "<nil>"
	}
	return e.x.Error()
}

func (e Error) MarshalText() ([]byte, error) {
	if e.x == nil {
		return []byte("null"), nil
	}
	return []byte(e.x.Error()), nil
}
