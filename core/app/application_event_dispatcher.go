package app

import "github.com/gjh33/SurrealEngine/core/event"

// ApplicationEventsDispatcher is a event.Dispatcher that sends out blocking events (processed immediately) for application life cycle.
type ApplicationEventsDispatcher struct {
	startupListeners     []ApplicationStartupListener
	initializedListeners []ApplicationInitializedListener
	updateListeners      []ApplicationUpdateListener
	quitListeners        []ApplicationQuitListener
	cleanedUpListeners   []ApplicationCleanedUpListener
}

// Subscribe implements the event.Dispatcher interface
func (dispatcher *ApplicationEventsDispatcher) Subscribe(subscriber event.Subscriber) error {
	subscribed := false

	if startupSubscriber, ok := subscriber.(ApplicationStartupListener); ok {
		subscribed = true
		dispatcher.startupListeners = append(dispatcher.startupListeners, startupSubscriber)
	}

	if initSubscriber, ok := subscriber.(ApplicationInitializedListener); ok {
		subscribed = true
		dispatcher.initializedListeners = append(dispatcher.initializedListeners, initSubscriber)
	}

	if updateSubscriber, ok := subscriber.(ApplicationUpdateListener); ok {
		subscribed = true
		dispatcher.updateListeners = append(dispatcher.updateListeners, updateSubscriber)
	}

	if quitSubscriber, ok := subscriber.(ApplicationQuitListener); ok {
		subscribed = true
		dispatcher.quitListeners = append(dispatcher.quitListeners, quitSubscriber)
	}

	if cleanedUpSubscriber, ok := subscriber.(ApplicationCleanedUpListener); ok {
		subscribed = true
		dispatcher.cleanedUpListeners = append(dispatcher.cleanedUpListeners, cleanedUpSubscriber)
	}

	if subscribed {
		return nil
	}

	return &event.UnknownSubscriberError{}
}

// Dispatch implements the Dispatcher interface
func (dispatcher *ApplicationEventsDispatcher) Dispatch(e event.Event) error {
	if _, ok := e.(ApplicationStartupEvent); ok {
		for _, subscriber := range dispatcher.startupListeners {
			subscriber.OnApplicationStartup()
		}
	} else if _, ok := e.(ApplicationInitializedEvent); ok {
		for _, subscriber := range dispatcher.initializedListeners {
			subscriber.OnApplicationInitialized()
		}
	} else if _, ok := e.(ApplicationUpdateEvent); ok {
		for _, subscriber := range dispatcher.updateListeners {
			subscriber.OnApplicationUpdate()
		}
	} else if _, ok := e.(ApplicationQuitEvent); ok {
		for _, subscriber := range dispatcher.quitListeners {
			subscriber.OnApplicationQuit()
		}
	} else if _, ok := e.(ApplicationCleanedUpEvent); ok {
		for _, subscriber := range dispatcher.cleanedUpListeners {
			subscriber.OnApplicationCleanedUp()
		}
	} else {
		return &event.UnknownEventError{}
	}

	return nil
}

// ApplicationStartupEvent is the event called right after the application is started, before any initialization.
// Things like graphics libraries etc will not yet be initialized
type ApplicationStartupEvent struct{}

// ApplicationStartupListener defines the subscriber interface for the ApplicationStartupEvent
type ApplicationStartupListener interface {
	OnApplicationStartup()
}

// ApplicationInitializedEvent is the event called after initialization. It is called just before the first update loop.
// graphics libraries will be initialized
type ApplicationInitializedEvent struct{}

// ApplicationInitializedListener defines the subscriber interface for the ApplicationInitializedEvent
type ApplicationInitializedListener interface {
	OnApplicationInitialized()
}

// ApplicationUpdateEvent is the event called continuously in the run loop. Delta time is provided in the time constant in app package
type ApplicationUpdateEvent struct{}

// ApplicationUpdateListener defines the subsciber interface for the ApplicationUpdateEvent
type ApplicationUpdateListener interface {
	OnApplicationUpdate()
}

// ApplicationQuitEvent is the event called when quitting the application before cleanup. It is called immediately after exiting the game loop
type ApplicationQuitEvent struct{}

// ApplicationQuitListener defines the subscriber interface for the ApplicationQuitEvent
type ApplicationQuitListener interface {
	OnApplicationQuit()
}

// ApplicationCleanedUpEvent is the event called after the application is cleaned up after quitting. It is called just before ending the main thread.
type ApplicationCleanedUpEvent struct{}

// ApplicationCleanedUpListener defines the subscriber interface for the ApplicationCleanedUpEvent
type ApplicationCleanedUpListener interface {
	OnApplicationCleanedUp()
}
