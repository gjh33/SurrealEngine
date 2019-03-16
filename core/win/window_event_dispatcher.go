package win

import "github.com/gjh33/SurrealEngine/core/event"

// WindowEventsDispatcher is a event.Dispatcher that sends out blocking events (processed immediately) for window related state changes.
type WindowEventsDispatcher struct {
	closeSubs           []WindowClosedListener
	closeRequestSubs    []WindowCloseRequestListener
	iconifySubs         []WindowIconifiedListener
	restoredSubs        []WindowRestoredListener
	focusLostSubs       []WindowFocusLostListener
	focusedSubs         []WindowFocusedListener
	movedSubs           []WindowMovedListener
	settingsChangedSubs []WindowSettingsChangedListener
	cursorEnterSubs     []WindowCursorEnteredListener
	cursorExitSubs      []WindowCursorExitListener
}

// Subscribe implements the event.Dispatcher interface
func (dispatcher *WindowEventsDispatcher) Subscribe(subscriber event.Subscriber) error {
	subscribed := false

	if sub, ok := subscriber.(WindowClosedListener); ok {
		subscribed = true
		dispatcher.closeSubs = append(dispatcher.closeSubs, sub)
	}
	if sub, ok := subscriber.(WindowCloseRequestListener); ok {
		subscribed = true
		dispatcher.closeRequestSubs = append(dispatcher.closeRequestSubs, sub)
	}
	if sub, ok := subscriber.(WindowIconifiedListener); ok {
		subscribed = true
		dispatcher.iconifySubs = append(dispatcher.iconifySubs, sub)
	}
	if sub, ok := subscriber.(WindowRestoredListener); ok {
		subscribed = true
		dispatcher.restoredSubs = append(dispatcher.restoredSubs, sub)
	}
	if sub, ok := subscriber.(WindowFocusLostListener); ok {
		subscribed = true
		dispatcher.focusLostSubs = append(dispatcher.focusLostSubs, sub)
	}
	if sub, ok := subscriber.(WindowFocusedListener); ok {
		subscribed = true
		dispatcher.focusedSubs = append(dispatcher.focusedSubs, sub)
	}
	if sub, ok := subscriber.(WindowMovedListener); ok {
		subscribed = true
		dispatcher.movedSubs = append(dispatcher.movedSubs, sub)
	}
	if sub, ok := subscriber.(WindowSettingsChangedListener); ok {
		subscribed = true
		dispatcher.settingsChangedSubs = append(dispatcher.settingsChangedSubs, sub)
	}
	if sub, ok := subscriber.(WindowCursorEnteredListener); ok {
		subscribed = true
		dispatcher.cursorEnterSubs = append(dispatcher.cursorEnterSubs, sub)
	}
	if sub, ok := subscriber.(WindowCursorExitListener); ok {
		subscribed = true
		dispatcher.cursorExitSubs = append(dispatcher.cursorExitSubs, sub)
	}

	if subscribed {
		return nil
	}

	return &event.UnknownSubscriberError{}
}

// Dispatch implements the Dispatcher interface
func (dispatcher *WindowEventsDispatcher) Dispatch(e event.Event) error {
	switch v := e.(type) {
	case WindowClosedEvent:
		for _, sub := range dispatcher.closeSubs {
			sub.OnWindowClosed(v)
		}
	case WindowCloseRequestEvent:
		for _, sub := range dispatcher.closeRequestSubs {
			sub.OnWindowCloseRequested(v)
		}
	case WindowIconifiedEvent:
		for _, sub := range dispatcher.iconifySubs {
			sub.OnWindowIconified(v)
		}
	case WindowRestoredEvent:
		for _, sub := range dispatcher.restoredSubs {
			sub.OnWindowRestored(v)
		}
	case WindowFocusLostEvent:
		for _, sub := range dispatcher.focusLostSubs {
			sub.OnWindowFocusLost(v)
		}
	case WindowFocusedEvent:
		for _, sub := range dispatcher.focusedSubs {
			sub.OnWindowFocused(v)
		}
	case WindowMovedEvent:
		for _, sub := range dispatcher.movedSubs {
			sub.OnWindowMoved(v)
		}
	case WindowSettingsChangedEvent:
		for _, sub := range dispatcher.settingsChangedSubs {
			sub.OnWindowSettingsChanged(v)
		}
	case WindowCursorEnteredEvent:
		for _, sub := range dispatcher.cursorEnterSubs {
			sub.OnWindowCursorEnter(v)
		}
	case WindowCursorExitEvent:
		for _, sub := range dispatcher.cursorExitSubs {
			sub.OnWindowCursorExit(v)
		}
	}

	return nil
}

// BaseWindowEvent holds the base parameters for a window event
type BaseWindowEvent struct {
	Window Window
}

// WindowClosedEvent is the event called after the window has closed
type WindowClosedEvent struct {
	BaseWindowEvent
}

// WindowClosedListener defines the subscriber interface for WindowClosedEvent
type WindowClosedListener interface {
	OnWindowClosed(e WindowClosedEvent)
}

// WindowCloseRequestEvent is the event called right after a close request has been called, before the window closes
type WindowCloseRequestEvent struct {
	BaseWindowEvent
}

// WindowCloseRequestListener defines the subscriber interface for WindowCloseRequestEvent
type WindowCloseRequestListener interface {
	OnWindowCloseRequested(e WindowCloseRequestEvent)
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

// WindowMovedEvent is the event called when a window is moved
type WindowMovedEvent struct {
	BaseWindowEvent
	OldLocation Location
	NewLocation Location
}

// WindowMovedListener defines the subscriber interface for WindowMovedEvent
type WindowMovedListener interface {
	OnWindowMoved(e WindowMovedEvent)
}

// WindowSettingsChangedEvent is the event called when a window's resolution is changed
type WindowSettingsChangedEvent struct {
	BaseWindowEvent
	OldWindowSettings WindowSettings
	NewWindowSettings WindowSettings
}

// WindowSettingsChangedListener defines the subscriber interface for WindowSettingsChangedEvent
type WindowSettingsChangedListener interface {
	OnWindowSettingsChanged(e WindowSettingsChangedEvent)
}

// WindowCursorEnteredEvent is the event called when a system cursor enters the window region
type WindowCursorEnteredEvent struct {
	BaseWindowEvent
	MouseX int
	MouseY int
}

// WindowCursorEnteredListener defines the subscriber interface for WindowCursorEnteredEvent
type WindowCursorEnteredListener interface {
	OnWindowCursorEnter(e WindowCursorEnteredEvent)
}

// WindowCursorExitEvent is the event called when a system cursor exits the window region
type WindowCursorExitEvent struct {
	BaseWindowEvent
	MouseX int
	MouseY int
}

// WindowCursorExitListener defines the subscriber interface for WindowCursorExitEvent
type WindowCursorExitListener interface {
	OnWindowCursorExit(e WindowCursorExitEvent)
}
