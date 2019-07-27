package win

import (
	"github.com/gjh33/SurrealEngine/core/event"
)

// WindowEventsDispatcher is a event.Dispatcher that sends out blocking events (processed immediately) for window related state changes.
type WindowEventsDispatcher struct {
	initializeSubs   []WindowInitializedListener
	createSubs       []WindowCreatedListener
	showSubs         []WindowShownListener
	hideSubs         []WindowHiddenListener
	focusLostSubs    []WindowFocusLostListener
	focusedSubs      []WindowFocusedListener
	iconifySubs      []WindowIconifiedListener
	restoredSubs     []WindowRestoredListener
	closeSubs        []WindowClosedListener
	closeRequestSubs []WindowCloseRequestedListener
	resizeSubs       []WindowResizedListener
	locationSubs     []WindowLocationChangedListener
	fullscreenSubs   []WindowFullscreenListener
	windowedSubs     []WindowWindowedListener
}

// Subscribe implements the event.Dispatcher interface
func (dispatcher *WindowEventsDispatcher) Subscribe(subscriber event.Subscriber) error {
	subscribed := false

	if sub, ok := subscriber.(WindowInitializedListener); ok {
		subscribed = true
		dispatcher.initializeSubs = append(dispatcher.initializeSubs, sub)
	}
	if sub, ok := subscriber.(WindowCreatedListener); ok {
		subscribed = true
		dispatcher.createSubs = append(dispatcher.createSubs, sub)
	}
	if sub, ok := subscriber.(WindowShownListener); ok {
		subscribed = true
		dispatcher.showSubs = append(dispatcher.showSubs, sub)
	}
	if sub, ok := subscriber.(WindowHiddenListener); ok {
		subscribed = true
		dispatcher.hideSubs = append(dispatcher.hideSubs, sub)
	}
	if sub, ok := subscriber.(WindowFocusLostListener); ok {
		subscribed = true
		dispatcher.focusLostSubs = append(dispatcher.focusLostSubs, sub)
	}
	if sub, ok := subscriber.(WindowFocusedListener); ok {
		subscribed = true
		dispatcher.focusedSubs = append(dispatcher.focusedSubs, sub)
	}
	if sub, ok := subscriber.(WindowIconifiedListener); ok {
		subscribed = true
		dispatcher.iconifySubs = append(dispatcher.iconifySubs, sub)
	}
	if sub, ok := subscriber.(WindowRestoredListener); ok {
		subscribed = true
		dispatcher.restoredSubs = append(dispatcher.restoredSubs, sub)
	}
	if sub, ok := subscriber.(WindowClosedListener); ok {
		subscribed = true
		dispatcher.closeSubs = append(dispatcher.closeSubs, sub)
	}
	if sub, ok := subscriber.(WindowCloseRequestedListener); ok {
		subscribed = true
		dispatcher.closeRequestSubs = append(dispatcher.closeRequestSubs, sub)
	}
	if sub, ok := subscriber.(WindowResizedListener); ok {
		subscribed = true
		dispatcher.resizeSubs = append(dispatcher.resizeSubs, sub)
	}
	if sub, ok := subscriber.(WindowLocationChangedListener); ok {
		subscribed = true
		dispatcher.locationSubs = append(dispatcher.locationSubs, sub)
	}
	if sub, ok := subscriber.(WindowFullscreenListener); ok {
		subscribed = true
		dispatcher.fullscreenSubs = append(dispatcher.fullscreenSubs, sub)
	}
	if sub, ok := subscriber.(WindowWindowedListener); ok {
		subscribed = true
		dispatcher.windowedSubs = append(dispatcher.windowedSubs, sub)
	}

	if subscribed {
		return nil
	}

	return &event.UnknownSubscriberError{}
}

// Dispatch implements the Dispatcher interface
func (dispatcher *WindowEventsDispatcher) Dispatch(e event.Event) error {
	switch v := e.(type) {
	case WindowInitializedEvent:
		for _, sub := range dispatcher.initializeSubs {
			sub.OnWindowInitialized(v)
		}
	case WindowCreatedEvent:
		for _, sub := range dispatcher.createSubs {
			sub.OnWindowCreated(v)
		}
	case WindowHiddenEvent:
		for _, sub := range dispatcher.hideSubs {
			sub.OnWindowHidden(v)
		}
	case WindowFocusLostEvent:
		for _, sub := range dispatcher.focusLostSubs {
			sub.OnWindowFocusLost(v)
		}
	case WindowFocusedEvent:
		for _, sub := range dispatcher.focusedSubs {
			sub.OnWindowFocused(v)
		}
	case WindowIconifiedEvent:
		for _, sub := range dispatcher.iconifySubs {
			sub.OnWindowIconified(v)
		}
	case WindowRestoredEvent:
		for _, sub := range dispatcher.restoredSubs {
			sub.OnWindowRestored(v)
		}
	case WindowClosedEvent:
		for _, sub := range dispatcher.closeSubs {
			sub.OnWindowClosed(v)
		}
	case WindowCloseRequestedEvent:
		for _, sub := range dispatcher.closeRequestSubs {
			sub.OnWindowCloseRequested(v)
		}
	case WindowResizedEvent:
		for _, sub := range dispatcher.resizeSubs {
			sub.OnWindowResized(v)
		}
	case WindowLocationChangedEvent:
		for _, sub := range dispatcher.locationSubs {
			sub.OnWindowMoved(v)
		}
	case WindowFullscreenEvent:
		for _, sub := range dispatcher.fullscreenSubs {
			sub.OnWindowFullscreen(v)
		}
	case WindowWindowedEvent:
		for _, sub := range dispatcher.windowedSubs {
			sub.OnWindowWindowed(v)
		}
	default:
		return &event.UnknownEventError{}
	}

	return nil
}

// BaseWindowEvent holds the base parameters for a window event
type BaseWindowEvent struct {
	Window Window
}

// WindowInitializedEvent is called when the window is initialized
type WindowInitializedEvent struct {
	BaseWindowEvent
}

// WindowInitializedListener defines the subscriber interface for WindowInitializedEvent
type WindowInitializedListener interface {
	OnWindowInitialized(e WindowInitializedEvent)
}

// WindowCreatedEvent is called when the window is created
type WindowCreatedEvent struct {
	BaseWindowEvent
}

// WindowCreatedListener defines the subscriber interface for WindowCreatedEvent
type WindowCreatedListener interface {
	OnWindowCreated(e WindowCreatedEvent)
}

// WindowShownEvent is called when a window becomes visible to the user
type WindowShownEvent struct {
	BaseWindowEvent
}

// WindowShownListener defines the subscriber interface for WindowShownEvent
type WindowShownListener interface {
	OnWindowShown(e WindowShownEvent)
}

// WindowHiddenEvent is called when a window becomes hidden from the user
type WindowHiddenEvent struct {
	BaseWindowEvent
}

// WindowHiddenListener defines the subscriber interface for WindowHiddenEvent
type WindowHiddenListener interface {
	OnWindowHidden(e WindowHiddenEvent)
}

// WindowFocusLostEvent is the event called when a window is unfocused
type WindowFocusLostEvent struct {
	BaseWindowEvent
}

// WindowFocusLostListener defines the subscriber interface for WindowFocusLostEvent
type WindowFocusLostListener interface {
	OnWindowFocusLost(e WindowFocusLostEvent)
}

// WindowFocusedEvent is the event called when a window regains focus
type WindowFocusedEvent struct {
	BaseWindowEvent
}

// WindowFocusedListener defines the subscriber interface for WindowFocusedEvent
type WindowFocusedListener interface {
	OnWindowFocused(e WindowFocusedEvent)
}

// WindowIconifiedEvent is the event called when a window is minimized/iconified
type WindowIconifiedEvent struct {
	BaseWindowEvent
}

// WindowIconifiedListener defines the subscriber interface for WindowIconifiedEvent
type WindowIconifiedListener interface {
	OnWindowIconified(e WindowIconifiedEvent)
}

// WindowRestoredEvent is the event called when a window is un iconified
type WindowRestoredEvent struct {
	BaseWindowEvent
}

// WindowRestoredListener defines the subscriber interface for WindowRestoredEvent
type WindowRestoredListener interface {
	OnWindowRestored(e WindowRestoredEvent)
}

// WindowClosedEvent is the event called after the window has closed
type WindowClosedEvent struct {
	BaseWindowEvent
}

// WindowClosedListener defines the subscriber interface for WindowClosedEvent
type WindowClosedListener interface {
	OnWindowClosed(e WindowClosedEvent)
}

// WindowCloseRequestedEvent is the event called right after a close request has been called, before the window closes
type WindowCloseRequestedEvent struct {
	BaseWindowEvent
}

// WindowCloseRequestedListener defines the subscriber interface for WindowCloseRequestedEvent
type WindowCloseRequestedListener interface {
	OnWindowCloseRequested(e WindowCloseRequestedEvent)
}

// WindowResizedEvent is called when the window is resized
type WindowResizedEvent struct {
	BaseWindowEvent
	OldSize Size
	NewSize Size
}

// WindowResizedListener defines the subscriber interface for WindowResizedEvent
type WindowResizedListener interface {
	OnWindowResized(e WindowResizedEvent)
}

// WindowLocationChangedEvent is the event called when a window is moved
type WindowLocationChangedEvent struct {
	BaseWindowEvent
	OldLocation Location
	NewLocation Location
}

// WindowLocationChangedListener defines the subscriber interface for WindowLocationChangedEvent
type WindowLocationChangedListener interface {
	OnWindowMoved(e WindowLocationChangedEvent)
}

// WindowFullscreenEvent is called when the window enters fullscreen mode
type WindowFullscreenEvent struct {
	BaseWindowEvent
}

// WindowFullscreenListener defines the subscribe interface for WindowFullscreenEvent
type WindowFullscreenListener interface {
	OnWindowFullscreen(e WindowFullscreenEvent)
}

// WindowWindowedEvent is called when the window is taken out of fullscreen mode
type WindowWindowedEvent struct {
	BaseWindowEvent
}

// WindowWindowedListener defines the subscriber interface for WindowWindowedEvent
type WindowWindowedListener interface {
	OnWindowWindowed(e WindowWindowedEvent)
}
