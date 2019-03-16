package win

import (
	"image/color"

	"github.com/gjh33/SurrealEngine/core/event"
)

// Window represents a window on your native operating system. This API is platform independent.
type Window interface {
	// Implemented in BaseWindow
	ActiveSettings() WindowSettings // Retrieve the window's active settings
	event.Dispatcher                // Don't forget to dispatch the events in your implementation

	// Implemented by platform
	Initialize() error               // Any initialization needed before calling any other methods
	Modify(WindowSettings) error     // Updates the window settings and updates the window itself if open
	DefaultSettings() WindowSettings // Returns the default settings of the current platform
	AvailableResolutions() []Size    // Returns valid resolutions for current system
	IsOpen() bool                    // Whether or not the window is open
	Launch() error                   // Opens the window and initializes it for use
	Close() error                    // Closes the window and cleans up necessary resources
	Draw() error                     // Flushes the graphics api calls to the screen
}

// BaseWindow represents the basis for a window struct in Surreal.
type BaseWindow struct {
	WindowSettings
	WindowState
	WindowEventsDispatcher
}

// ActiveSettings implements the Window interface. Returns the window's live settings.
func (window *BaseWindow) ActiveSettings() WindowSettings {
	return window.WindowSettings
}

// WindowSettings represent the set of display settings a user can change
type WindowSettings struct {
	Title        string
	Resolution   Size
	ClearColor   color.Color
	DisplayMode  DisplayMode
	VSync        bool
	AntiAliasing AntiAliasingMode
	MSAA         MSAAMode
	CursorLocked bool
	CursorHidden bool
}

// WindowState represents the set of variables that represent the window state
// These are queryable and set via function calls individually if possible
type WindowState struct {
	ScreenPosition Location
	Iconified      bool
	Focused        bool
}

// DisplayMode is an enum for different display modes
type DisplayMode int

// Declaring DisplayMode enum values
const (
	Windowed DisplayMode = iota
	BorderlessWindow
	Fullscreen
)

// AntiAliasingMode is the different AA modes supported
type AntiAliasingMode int

// Declaring AntiAliasingMode enum values
const (
	None AntiAliasingMode = iota
	MSAA
)

// MSAAMode are the different modes for Multi Sample Anti Aliasing. It determines the number of samples taken.
type MSAAMode int

// Declaring MSAAMode enum values
const (
	Sample2X MSAAMode = iota
	Sample4X
	Sample8X
	Sample16X
	Sample32X
)

// Size represents the dimensions of a window
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
