package event

// Subscriber represents an object that can subscribe to a dispatcher.
type Subscriber interface{}

// Event represents an object that holds data specific to a type of event.
type Event interface{}

// Dispatcher is the interface for an event dispatcher
type Dispatcher interface {
	// Should use type switch to check it implements and interface with a supported callback function, then remember it.
	Subscribe(Subscriber) error
	// Should use type switch to check event data type and call the correct subscribers
	Dispatch(Event) error
}

// UnknownSubscriberError should be thrown if the interface of the subscriber is not supported by the dispatcher
type UnknownSubscriberError struct{}

// Error implements the error interface
func (err *UnknownSubscriberError) Error() string {
	return "Unknown Subscriber: An unsupported interface tried to subscribe to this dispatcher. Check the dispatcher implementation to see valid listener interfaces"
}

// UnknownEventError should be thrown if the struct trying to be dispatched is not recognized as a valid event type
type UnknownEventError struct{}

func (err *UnknownEventError) Error() string {
	return "Unkown Event: An unsupported type tried to be fired as event data to this dispatcher. Check the dispatcher implementation to see valid event data struct"
}
