package win

import (
	"image"

	"github.com/gjh33/SurrealEngine/core/event"
)

// Window represents a window on your native operating system. This API is platform independent.
type Window interface {
	event.Dispatcher // Don't forget to dispatch the events in your implementation

	// Actions
	// - Most actions can be set before creation to set initial state
	// - After changing size or location, verify after that the window was able to set the correct value
	Initialize() error                   // Called before opening the window
	Create() error                       // Creates a window that will try to honor the existing window state
	Show() error                         // Makes the window visible to the user
	Hide() error                         // Makes the window invisible to the user
	Focus() error                        // Force focus on the window
	Iconify() error                      // Minimizes the window
	Restore() error                      // Unminimizes the window
	Close() error                        // Destroys the window
	Resize(size Size) error              // Resizes the window
	SetTitle(title string) error         // Sets the window title
	SetIcons(icons []image.Image) error  // Set window icon to best matched image. To return to default pass nil
	SetLocation(location Location) error // Sets the position of the upper left corner of the window content
	SetFullscreen(fullscreen bool) error // Sets the window to fullscreen mode
	// NOTE: Below requires glfw 3.3 and since I'm not interested in forking the go-glfw right now, we'll just disable them
	// but really these should be relatively unused features anyways. So for now all they do it set initial state for
	// window creation
	SetResizable(resizable bool) error // Sets whether or not the window can be resized
	SetDecorated(decorated bool) error // Sets whether or not the window is decorated or just content
	SetCursorLocked(locked bool) error // Sets whether the cursor is locked to the center of window or not
	SetCursorHidden(hidden bool) error // Sets whether the cursor is visible over the window

	// Information Queries
	// NOTE: While many of these could be implemented on base window, rather than commit to an implementation
	// I chose to leave it to the specific implementation of the Window interface
	IsInitialized() bool // Check if the window has been initialized
	IsCreated() bool     // Check if the window has been created yet. A Destroyed window is no longer created
	IsVisible() bool     // Whether a window is being shown or hidden
	IsFocused() bool     // Whether or not the window is
	IsIconified() bool   // Whether a window is minimized or not
	Size() Size          // Window's current size
	Title() string       // Window's current title
	Location() Location  // Window's current location
	Fullscreen() bool    // Whether or not the window is fullscreen or in windowed mode
	Resizable() bool     // Whether or not the window can be resized
	Decorated() bool     // Whether or not the window is decorated
	CursorLocked() bool  // Whether or not the cursor is locked to the center of the window
	CursorHidden() bool  // Whether or not the cursor is hidden while over
	ShouldClose() bool   // Whether or not for any reason the window wants to close. I.e. pressing close button
}

// BaseWindow represents the basis for a window struct in Surreal.
type BaseWindow struct {
	Settings
	State
	WindowEventsDispatcher
}

// ActiveSettings implements the Window interface. Returns the window's live settings.
func (window *BaseWindow) ActiveSettings() Settings {
	return window.Settings
}

// Settings represent the set of display settings that only change when requested via API
type Settings struct {
	Title        string
	Icons        []image.Image
	FullScreen   bool
	Resizable    bool
	Decorated    bool
	CursorLocked bool
	CursorHidden bool
}

// State represents a set of variables that may change without explicit API Calls (i.e. via user interaction)
type State struct {
	Initialized bool
	Created     bool
	Visible     bool
	Focused     bool
	Iconified   bool
	Size        Size
	Location    Location
}

// Size represents the dimensions of a window in pixels
type Size struct {
	Width  int
	Height int
}

// PixelCount returns the number of pixels for a given size
func (size Size) PixelCount() int {
	return size.Width * size.Height
}

// Location represents the screen position of the upper left corner of the window
type Location struct {
	X int
	Y int
}
