package win

import (
	"image/color"
	"math"
	"runtime"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

// OpenGLWindow is the openGL Implementation of a window
type OpenGLWindow struct {
	BaseWindow

	Open   bool
	Handle *glfw.Window

	initialized bool
}

// IsOpen implements Window interface
func (window *OpenGLWindow) IsOpen() bool {
	return window.Open
}

// Launch implements Window interface
func (window *OpenGLWindow) Launch() error {
	var err error

	// Query for information
	primaryMonitor := glfw.GetPrimaryMonitor()
	primaryVideoMode := primaryMonitor.GetVideoMode()

	// Set window hints
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	glfw.WindowHint(glfw.RedBits, primaryVideoMode.RedBits)
	glfw.WindowHint(glfw.GreenBits, primaryVideoMode.GreenBits)
	glfw.WindowHint(glfw.BlueBits, primaryVideoMode.BlueBits)
	glfw.WindowHint(glfw.RefreshRate, primaryVideoMode.RefreshRate)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	switch window.AntiAliasing {
	case None:
		glfw.WindowHint(glfw.Samples, 1)
	case MSAA:
		switch window.MSAA {
		case Sample2X:
			glfw.WindowHint(glfw.Samples, 2)
		case Sample4X:
			glfw.WindowHint(glfw.Samples, 4)
		case Sample8X:
			glfw.WindowHint(glfw.Samples, 8)
		case Sample16X:
			glfw.WindowHint(glfw.Samples, 16)
		case Sample32X:
			glfw.WindowHint(glfw.Samples, 32)
		}
	}

	if !window.VSync {
		glfw.WindowHint(glfw.DoubleBuffer, glfw.False)
	}

	// Create the window
	switch window.DisplayMode {
	case Windowed:
		glfw.WindowHint(glfw.Resizable, glfw.False)
		window.Handle, err = glfw.CreateWindow(window.Resolution.Width, window.Resolution.Height, window.Title, nil, nil)
	case BorderlessWindow:
		glfw.WindowHint(glfw.Resizable, glfw.False)
		glfw.WindowHint(glfw.Decorated, glfw.False)
		window.Handle, err = glfw.CreateWindow(primaryVideoMode.Width, primaryVideoMode.Height, window.Title, nil, nil)
	case Fullscreen:
		window.Handle, err = glfw.CreateWindow(primaryVideoMode.Width, primaryVideoMode.Height, window.Title, primaryMonitor, nil)
	}

	if err != nil {
		return err
	}

	mode := glfw.CursorNormal
	if window.CursorLocked {
		mode = glfw.CursorDisabled
	} else {
		mode = glfw.CursorHidden
	}
	window.Handle.SetInputMode(glfw.CursorMode, mode)
	window.Handle.Focus()
	x, y := window.Handle.GetPos()
	window.ScreenPosition = Location{x, y}
	window.Iconified = false
	window.Focused = true

	// Register callbacks
	window.Handle.SetCloseCallback(window.windowCloseEventCallback)
	window.Handle.SetCursorPosCallback(window.windowCursorMoveEventCallback)
	window.Handle.SetCursorEnterCallback(window.windowCursorEnterExitEventCallback)
	window.Handle.SetIconifyCallback(window.windowIconifiedCallback)
	window.Handle.SetFocusCallback(window.windowFocusedCallback)
	window.Handle.SetPosCallback(window.windowPositionCallback)

	// Set context to be current
	window.Handle.MakeContextCurrent()

	// Handle vsync
	if window.VSync {
		glfw.SwapInterval(1)
	} else {
		glfw.SwapInterval(0)
	}

	// Setup and Bind openGL context
	if err = gl.Init(); err != nil {
		return err
	}

	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.CULL_FACE)
	gl.Enable(gl.MULTISAMPLE)
	gl.FrontFace(gl.CCW)
	r, g, b, a := window.ClearColor.RGBA()
	gl.ClearColor(float32(r)/255, float32(g)/255, float32(b)/255, float32(a))

	window.Open = true
	return nil
}

// Modify implements Window interface
func (window *OpenGLWindow) Modify(settings WindowSettings) error {
	e := WindowSettingsChangedEvent{}
	e.Window = window
	e.OldWindowSettings = window.WindowSettings
	e.NewWindowSettings = settings

	if !window.IsOpen() {
		window.WindowSettings = settings
		window.Dispatch(e)
		return nil
	}

	requireRelaunch := false

	if window.DisplayMode != settings.DisplayMode {
		requireRelaunch = true
	}

	// Anything with MSAA requires relaunch
	if window.AntiAliasing == MSAA {
		if window.AntiAliasing != settings.AntiAliasing {
			requireRelaunch = true
		} else if window.MSAA != settings.MSAA {
			requireRelaunch = true
		}
	} else if settings.AntiAliasing == MSAA {
		requireRelaunch = true
	}

	if requireRelaunch {
		err := window.Close()
		if err != nil {
			return err
		}
		window.WindowSettings = settings
		err = window.Launch()
		if err != nil {
			return err
		}
	} else {
		window.Handle.MakeContextCurrent()
		if window.Title != settings.Title {
			window.Title = settings.Title
			window.Handle.SetTitle(window.Title)
		}
		if window.Resolution != settings.Resolution {
			if window.DisplayMode == Windowed {
				window.Handle.SetSize(settings.Resolution.Width, settings.Resolution.Height)
			}
			window.Resolution.Width = settings.Resolution.Width
			window.Resolution.Height = settings.Resolution.Height
		}
		if window.ClearColor != settings.ClearColor {
			window.ClearColor = settings.ClearColor
			r, g, b, a := window.ClearColor.RGBA()
			gl.ClearColor(float32(r)/255, float32(g)/255, float32(b)/255, float32(a))
		}
		if window.AntiAliasing != settings.AntiAliasing {
			// Handle any future AA modes here. Right now this
			// will never fire without a relaunch
		}
		if window.VSync != settings.VSync {
			window.VSync = settings.VSync
			if window.VSync {
				glfw.SwapInterval(1)
			} else {
				glfw.SwapInterval(0)
			}
		}

		if window.CursorHidden != settings.CursorHidden || window.CursorLocked != settings.CursorLocked {
			mode := glfw.CursorNormal
			if settings.CursorLocked {
				mode = glfw.CursorDisabled
			} else {
				mode = glfw.CursorHidden
			}
			window.Handle.SetInputMode(glfw.CursorMode, mode)
		}
	}

	window.Dispatch(e)
	return nil
}

// DefaultSettings implements Window interface
func (window *OpenGLWindow) DefaultSettings() WindowSettings {

	// Query for information
	primaryMonitor := glfw.GetPrimaryMonitor()
	primaryVideoMode := primaryMonitor.GetVideoMode()

	settings := WindowSettings{
		Title:        "",
		Resolution:   Size{Width: primaryVideoMode.Width, Height: primaryVideoMode.Height},
		ClearColor:   color.Black,
		DisplayMode:  Fullscreen,
		VSync:        true,
		AntiAliasing: MSAA,
		MSAA:         Sample4X,
		CursorLocked: true,
		CursorHidden: true,
	}

	return settings
}

// AvailableResolutions implements Window interface
func (window *OpenGLWindow) AvailableResolutions() []Size {
	videoModes := window.Handle.GetMonitor().GetVideoModes()
	var resolutions []Size
	for _, mode := range videoModes {
		size := Size{Width: mode.Width, Height: mode.Height}
		contains := false
		for _, res := range resolutions {
			if res == size {
				contains = true
			}
		}

		if !contains {
			resolutions = append(resolutions, size)
		}
	}

	return resolutions
}

// Initialize implements the Window interface
// Initializing a window locks the OS thread it was intialized from
func (window *OpenGLWindow) Initialize() error {
	if !window.initialized {
		// Lock OS Thread (should be refactored to handle multiple windows)
		runtime.LockOSThread()

		err := glfw.Init()
		if err != nil {
			return err
		}
		window.initialized = true
	}
	return nil
}

// Close implements Window interface
func (window *OpenGLWindow) Close() error {
	e := WindowCloseRequestEvent{}
	e.Window = window
	window.Dispatch(e)
	window.Handle.Destroy()
	window.Open = false
	return nil
}

// Draw implements Window interface
func (window *OpenGLWindow) Draw() error {
	window.Handle.SwapBuffers()
	return nil
}

func (window *OpenGLWindow) windowCloseEventCallback(handle *glfw.Window) {
	window.Close()
}

func (window *OpenGLWindow) windowCursorMoveEventCallback(handle *glfw.Window, xpos float64, ypos float64) {
}

func (window *OpenGLWindow) windowCursorEnterExitEventCallback(handle *glfw.Window, entered bool) {
	if entered {
		e := WindowCursorEnteredEvent{}
		e.Window = window
		x, y := handle.GetCursorPos()
		e.MouseX = int(math.Floor(x))
		e.MouseY = int(math.Floor(y))
		window.Dispatch(e)
	} else {
		e := WindowCursorExitEvent{}
		e.Window = window
		x, y := handle.GetCursorPos()
		e.MouseX = int(math.Floor(x))
		e.MouseY = int(math.Floor(y))
		window.Dispatch(e)
	}
}

func (window *OpenGLWindow) windowIconifiedCallback(handle *glfw.Window, iconified bool) {
	window.Iconified = iconified
	if iconified {
		e := WindowIconifiedEvent{}
		e.Window = window
		window.Dispatch(e)
	} else {
		e := WindowRestoredEvent{}
		e.Window = window
		window.Dispatch(e)
	}
}

func (window *OpenGLWindow) windowFocusedCallback(handle *glfw.Window, focused bool) {
	window.Focused = focused
	if focused {
		e := WindowFocusedEvent{}
		e.Window = window
		window.Dispatch(e)
	} else {
		e := WindowFocusLostEvent{}
		e.Window = window
		window.Dispatch(e)
	}
}

func (window *OpenGLWindow) windowPositionCallback(handle *glfw.Window, x int, y int) {
	e := WindowMovedEvent{}
	e.Window = window
	e.OldLocation = window.ScreenPosition
	e.NewLocation.X = x
	e.NewLocation.Y = y
	window.ScreenPosition.X = x
	window.ScreenPosition.Y = y
	window.Dispatch(e)
}
