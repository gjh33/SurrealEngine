package app

import (
	"fmt"
	"image/color"

	"github.com/gjh33/SurrealEngine/core/win"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

// Application represents top most information about a Surreal Application
type Application struct {
	Name                        string          // The name of the application
	Version                     SemanticVersion // The version of the application
	ApplicationEventsDispatcher                 // Application is an event dispatcher

	windows []win.Window
}

// New is the default constructor for an Application
func New(name string, version string) (obj *Application) {
	obj = new(Application)
	obj.Name = name
	var err error
	obj.Version, err = ParseVersion(version)
	if err != nil {
		obj.Version = SemanticVersion{MajorRelease: 0, MinorRelease: 0, Patch: 0}
	}
	return
}

// Start is the main entry point for Surreal Applications
func (application *Application) Start() {
	test := &listenerTester{}
	application.Subscribe(test)
	application.Dispatch(ApplicationStartupEvent{})
	window := new(win.OpenGLWindow)
	err := window.Initialize()
	if err != nil {
		panic("Failed to initialize window!\n" + err.Error())
	}
	window.WindowSettings = window.DefaultSettings()
	window.DisplayMode = win.Windowed
	window.Resolution = win.Size{Width: 800, Height: 600}
	window.ClearColor = color.RGBA{R: 128, G: 0, B: 128, A: 1}
	window.CursorLocked = false
	err = window.Launch()

	if err != nil {
		panic("Failed to launch window!" + err.Error())
	}
	window.Subscribe(test)

	application.Dispatch(ApplicationInitializedEvent{})
	for window.IsOpen() {
		application.Dispatch(ApplicationUpdateEvent{})
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		window.Draw()
		glfw.PollEvents()
	}
	application.Dispatch(ApplicationQuitEvent{})
	application.Dispatch(ApplicationCleanedUpEvent{})
}

type listenerTester struct {
}

func (lt *listenerTester) OnWindowClosed(e win.WindowClosedEvent) {
	fmt.Println("Window Closed")
}

func (lt *listenerTester) OnWindowCloseRequest(e win.WindowCloseRequestEvent) {
	fmt.Println("Window Close Request Received")
}

func (lt *listenerTester) OnWindowIconified(e win.WindowIconifiedEvent) {
	fmt.Println("Window Iconified")
}

func (lt *listenerTester) OnWindowRestored(e win.WindowRestoredEvent) {
	fmt.Println("Window Restored")
}

func (lt *listenerTester) OnWindowFocusLost(e win.WindowFocusLostEvent) {
	fmt.Println("Window Focus Lost")
}

func (lt *listenerTester) OnWindowFocused(e win.WindowFocusedEvent) {
	fmt.Println("Window Focused")
}

func (lt *listenerTester) OnWindowMoved(e win.WindowMovedEvent) {
	fmt.Printf("Window Moved from (%v, %v) to (%v, %v)\n", e.OldLocation.X, e.OldLocation.Y, e.NewLocation.X, e.NewLocation.Y)
}

func (lt *listenerTester) OnWindowCursorEnter(e win.WindowCursorEnteredEvent) {
	fmt.Printf("Cursor entered window at (%v, %v)\n", e.MouseX, e.MouseY)
}

func (lt *listenerTester) OnWindowCursorExit(e win.WindowCursorExitEvent) {
	fmt.Printf("Cursor exited window at (%v, %v)\n", e.MouseX, e.MouseY)
}

func (lt *listenerTester) OnApplicationStartup() {
	fmt.Println("Application Started")
}

func (lt *listenerTester) OnApplicationInitialized() {
	fmt.Println("Application Initialized")
}

func (lt *listenerTester) OnApplicationQuit() {
	fmt.Println("Application Quit")
}

func (lt *listenerTester) OnApplicationCleanedUp() {
	fmt.Println("Application Cleaned Up")
}
