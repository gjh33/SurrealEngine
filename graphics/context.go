package graphics

import (
	"image/color"

	"github.com/gjh33/SurrealEngine/graphics/win"
)

// Context is the main platform agnostic API for graphics implementations
type Context interface {
	// Actions
	Initialize() error  // Initializes context and binds it to a window
	Window() win.Window // Returns the window this context is bound to

	IsInitialized() bool
}

// BaseContext implements some base state and settings functionality for a graphics context
type BaseContext struct {
	State
	Settings
}

// Settings are the platform agnostic graphics settings supported by the Surreal Engine
type Settings struct {
	DisplayResolution Resolution       // The resolution at which to render the content
	DisplayMode       DisplayMode      // How the content should be displayed in the window
	VSync             VSyncMode        // VSync mode
	AntiAliasing      AntiAliasingMode // How AA is performed
	MSAASamples       MSAAMode         // Only needs to be set if AA is MSAA mode
	TargetResolution  Resolution       // The resolution we render to
	OutputResolution  Resolution       // The resolution we output to the monitor
	ClearColor        color.Color      // TODO: Move to Camera
}

// State represents the changing state of a graphics context. These can change at any time
type State struct {
	Initialized bool
	Window      win.Window
}

// Resolution represents a resolution in pixels
type Resolution struct {
	Width  int
	Height int
}

// PixelCount returns the number of pixels for a given size
func (res Resolution) PixelCount() int {
	return res.Width * res.Height
}

// DisplayMode is an enum for different display modes
type DisplayMode int

// Declaring DisplayMode enum values
const (
	Windowed DisplayMode = iota
	BorderlessWindow
	Fullscreen
)

// VSyncMode is an enum for different vsync modes
type VSyncMode int

// Declaring VSyncMode enum values
const (
	NoSync VSyncMode = iota
	DoubleBuffered
	TripleBuffered
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
