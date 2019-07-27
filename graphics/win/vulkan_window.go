package win

import (
	"image"

	"github.com/pkg/errors"

	"github.com/vulkan-go/glfw/v3.3/glfw"
)

// VulkanWindow is the openGL Implementation of a window
type VulkanWindow struct {
	BaseWindow
	Handle *glfw.Window

	baseEvent BaseWindowEvent
}

// Initialize implements Window interface
func (window *VulkanWindow) Initialize() error {
	if err := glfw.Init(); err != nil {
		return err
	}
	if !glfw.VulkanSupported() {
		return errors.New("vulkan drivers not found")
	}

	window.baseEvent = BaseWindowEvent{window}

	// Set default settings
	window.Settings.Title = "Surreal Application"
	window.Settings.FullScreen = false
	window.Settings.Resizable = false
	window.Settings.Decorated = true
	window.Settings.CursorLocked = false
	window.Settings.CursorHidden = false

	window.State.Visible = true
	window.State.Focused = true
	window.State.Iconified = false
	window.State.Size = Size{1024, 720}

	// Default to center of screen
	videoMode := glfw.GetPrimaryMonitor().GetVideoMode()
	window.State.Location = Location{
		(videoMode.Width / 2) - (window.State.Size.Width / 2),
		(videoMode.Height / 2) - (window.State.Size.Height / 2),
	}

	window.State.Initialized = true
	_ = window.Dispatch(WindowInitializedEvent{window.baseEvent})
	return nil
}

// Create implements Window interface
func (window *VulkanWindow) Create() error {
	glfw.WindowHint(glfw.ClientAPI, glfw.NoAPI)
	glfw.WindowHint(glfw.Resizable, boolToGLFW(window.Settings.Resizable))
	glfw.WindowHint(glfw.Decorated, boolToGLFW(window.Settings.Decorated))
	glfw.WindowHint(glfw.Focused, boolToGLFW(window.State.Focused)) // I'm open to this being an option in the future
	glfw.WindowHint(glfw.Visible, boolToGLFW(window.State.Visible))
	glfw.WindowHint(glfw.AutoIconify, glfw.False) // If you set this to true, you are what's wrong with this world
	var err error
	if window.Settings.FullScreen {
		window.Handle, err = glfw.CreateWindow(window.State.Size.Width, window.State.Size.Height, window.Settings.Title, glfw.GetPrimaryMonitor(), nil)
	} else {
		window.Handle, err = glfw.CreateWindow(window.State.Size.Width, window.State.Size.Height, window.Settings.Title, nil, nil)
	}
	if err != nil {
		return err
	}
	window.State.Created = true
	if err := window.setCursorMode(); err != nil {
		return err
	}

	// Honor the state the window had before creation
	if window.State.Iconified {
		err = window.Handle.Iconify()
		if err != nil {
			return err
		}
	} else {
		err = window.Handle.Restore()
		if err != nil {
			return err
		}
	}
	window.Handle.SetPos(window.State.Location.X, window.State.Location.Y)
	window.Handle.SetIcon(window.Icons)

	// Update all values to make sure if anything wasn't created correctly, it's reflected in the model
	window.State.Visible = glfwToBool(window.Handle.GetAttrib(glfw.Visible))
	window.State.Focused = glfwToBool(window.Handle.GetAttrib(glfw.Focused))
	window.State.Iconified = glfwToBool(window.Handle.GetAttrib(glfw.Iconified))
	width, height := window.Handle.GetSize()
	window.State.Size = Size{width, height}
	x, y := window.Handle.GetPos()
	window.State.Location = Location{x, y}

	// Register events
	// Sometimes user api calls can fail, or a user changes the state so we need to use these callbacks
	window.Handle.SetFocusCallback(window.focusChangedCallback)
	window.Handle.SetPosCallback(window.locationChangedCallback)
	window.Handle.SetSizeCallback(window.sizeChangedCallback)
	window.Handle.SetIconifyCallback(window.iconifyChangedCallback)

	_ = window.Dispatch(WindowCreatedEvent{window.baseEvent})

	return nil
}

// Show implements Window interface
func (window *VulkanWindow) Show() error {
	if !window.Visible {
		if window.Created {
			window.Handle.Show()
		}
		window.State.Visible = true
		_ = window.Dispatch(WindowShownEvent{window.baseEvent})
	}
	return nil
}

// Hide implements Window interface
func (window *VulkanWindow) Hide() error {
	if window.Visible {
		if window.Created {
			window.Handle.Hide()
		}
		window.State.Visible = false
		_ = window.Dispatch(WindowHiddenEvent{window.baseEvent})
	}
	return nil
}

// Focus implements Window interface
// Events and state are handled in callback
func (window *VulkanWindow) Focus() error {
	if !window.Focused {
		if window.Created {
			if err := window.Handle.Focus(); err != nil {
				return err
			}
		} else {
			window.Focused = true
		}
	}
	return nil
}

// Iconify implements Window interface
// Events and state are handled in callback
func (window *VulkanWindow) Iconify() error {
	if !window.Iconified {
		if window.Created {
			if err := window.Handle.Iconify(); err != nil {
				return err
			}
		} else {
			window.Iconified = true
		}
	}
	return nil
}

// Restore implements Window interface
// Events and state are handled in callback
func (window *VulkanWindow) Restore() error {
	if window.Iconified {
		if window.Created {
			if err := window.Handle.Restore(); err != nil {
				return err
			}
		} else {
			window.Iconified = false
		}
	}
	return nil
}

// Close implements Window interface
func (window *VulkanWindow) Close() error {
	if window.Created {
		_ = window.Dispatch(WindowCloseRequestedEvent{window.baseEvent})
		window.Handle.Destroy()
		window.Handle = nil
		window.Created = false
		_ = window.Dispatch(WindowClosedEvent{window.baseEvent})
	}

	return nil
}

// Resize implements Window interface
// Events and state implemented in callback
func (window *VulkanWindow) Resize(size Size) error {
	if window.Created {
		window.Handle.SetSize(size.Width, size.Height)
	} else {
		window.State.Size = size
	}
	return nil
}

// SetIcons implements Window interface
func (window *VulkanWindow) SetIcons(icons []image.Image) error {
	if window.Created {
		window.Handle.SetIcon(icons)
	}
	window.Icons = icons
	return nil
}

// SetTitle implements Window interface
func (window *VulkanWindow) SetTitle(title string) error {
	if window.Created {
		window.Handle.SetTitle(title)
	}
	window.Settings.Title = title
	return nil
}

// SetLocation implements Window interface
// Events and state implemented in callback
func (window *VulkanWindow) SetLocation(location Location) error {
	if window.Created {
		window.Handle.SetPos(location.X, location.Y)
	} else {
		window.State.Location = location
	}
	return nil
}

// SetResizable implements Window interface
// GLFW 3.3 is required to set attributes post creation!
func (window *VulkanWindow) SetResizable(resizable bool) error {
	window.Settings.Resizable = resizable
	if window.Created {
		return errors.New("requires glfw 3.3")
	}
	return nil
}

// SetDecorated implements Window interface
// GLFW 3.3 is required to set attributes post creation!
func (window *VulkanWindow) SetDecorated(decorated bool) error {
	window.Settings.Decorated = decorated
	if window.Created {
		return errors.New("requires glfw 3.3")
	}
	return nil
}

// SetFullscreen implements window interface
func (window *VulkanWindow) SetFullscreen(fullscreen bool) error {
	if window.Created {
		if fullscreen && !window.FullScreen {
			window.Handle.SetMonitor(glfw.GetPrimaryMonitor(), window.State.Location.X, window.State.Location.Y, window.State.Size.Width, window.State.Size.Height, glfw.DontCare)
			_ = window.Dispatch(WindowFullscreenEvent{window.baseEvent})
		}
		if !fullscreen && window.FullScreen {
			window.Handle.SetMonitor(nil, window.State.Location.X, window.State.Location.Y, window.State.Size.Width, window.State.Size.Height, glfw.DontCare)
			_ = window.Dispatch(WindowWindowedEvent{window.baseEvent})
		}
	}
	window.FullScreen = fullscreen
	return nil
}

// SetCursorLocked implements Window interface
func (window *VulkanWindow) SetCursorLocked(locked bool) error {
	window.Settings.CursorLocked = locked
	if window.Created {
		if err := window.setCursorMode(); err != nil {
			return err
		}
	}
	return nil
}

// SetCursorHidden implements Window interface
func (window *VulkanWindow) SetCursorHidden(hidden bool) error {
	window.Settings.CursorHidden = hidden
	if window.Created {
		if err := window.setCursorMode(); err != nil {
			return err
		}
	}
	return nil
}

// IsInitialized implements Window interface
func (window *VulkanWindow) IsInitialized() bool {
	return window.Initialized
}

// IsCreated implements Window interface
func (window *VulkanWindow) IsCreated() bool {
	return window.Created
}

// IsVisible implements Window interface
func (window *VulkanWindow) IsVisible() bool {
	return window.Visible
}

// IsFocused implements Window interface
func (window *VulkanWindow) IsFocused() bool {
	return window.Focused
}

// IsIconified implements Window interface
func (window *VulkanWindow) IsIconified() bool {
	return window.Iconified
}

// Size implements Window interface
func (window *VulkanWindow) Size() Size {
	return window.State.Size
}

// Title implements Window interface
func (window *VulkanWindow) Title() string {
	return window.Settings.Title
}

// Location implements Window interface
func (window *VulkanWindow) Location() Location {
	return window.State.Location
}

// Resizable implements Window interface
func (window *VulkanWindow) Resizable() bool {
	return window.Settings.Resizable
}

// Decorated implements Window interface
func (window *VulkanWindow) Decorated() bool {
	return window.Settings.Decorated
}

// CursorLocked implements Window interface
func (window *VulkanWindow) CursorLocked() bool {
	return window.Settings.CursorLocked
}

// CursorHidden implements Window interface
func (window *VulkanWindow) CursorHidden() bool {
	return window.Settings.CursorHidden
}

// Fullscreen implements window interface
func (window *VulkanWindow) Fullscreen() bool {
	return window.FullScreen
}

// ShouldClose implements window interface
func (window *VulkanWindow) ShouldClose() bool {
	if !window.Created {
		return false
	}
	return window.Handle.ShouldClose()
}

// VerifyInitialized returns an error if the window has not been initialized
func (window *VulkanWindow) VerifyInitialized() error {
	if !window.Initialized {
		return errors.New("window must be initialized")
	}
	return nil
}

// VerifyCreated returns an error if the window has not been created
func (window *VulkanWindow) VerifyCreated() error {
	if !window.Created {
		return errors.New("window must be created")
	}
	return nil
}

func (window *VulkanWindow) setCursorMode() error {
	if err := window.VerifyCreated(); err != nil {
		return err
	}
	if window.Settings.CursorLocked {
		window.Handle.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
		window.Settings.CursorHidden = true // For now i have no choice. I need custom cursor to lock and show
	} else if window.Settings.CursorHidden {
		window.Handle.SetInputMode(glfw.CursorMode, glfw.CursorHidden)
	}
	return nil
}

func (window *VulkanWindow) focusChangedCallback(handle *glfw.Window, focused bool) {
	window.Focused = focused
	if !focused {
		_ = window.Dispatch(WindowFocusLostEvent{window.baseEvent})
	} else {
		_ = window.Dispatch(WindowFocusedEvent{window.baseEvent})
	}
}

func (window *VulkanWindow) locationChangedCallback(handle *glfw.Window, x int, y int) {
	newLocation := Location{x, y}
	oldLocation := window.State.Location
	window.State.Location = newLocation
	_ = window.Dispatch(WindowLocationChangedEvent{
		window.baseEvent,
		oldLocation,
		newLocation,
	})
}

func (window *VulkanWindow) sizeChangedCallback(handle *glfw.Window, width int, height int) {
	newSize := Size{width, height}
	oldSize := window.State.Size
	window.State.Size = newSize
	_ = window.Dispatch(WindowResizedEvent{
		window.baseEvent,
		oldSize,
		newSize,
	})
}

func (window *VulkanWindow) iconifyChangedCallback(handle *glfw.Window, iconified bool) {
	window.Iconified = iconified
	if iconified {
		_ = window.Dispatch(WindowIconifiedEvent{window.baseEvent})
	} else {
		_ = window.Dispatch(WindowRestoredEvent{window.baseEvent})
	}
}

func boolToGLFW(value bool) int {
	if value {
		return glfw.True
	}
	return glfw.False
}

func glfwToBool(value int) bool {
	if value == glfw.True {
		return true
	}
	return false
}
