package app

import (
	"fmt"

	"github.com/gjh33/SurrealEngine/graphics/win"

	gfx "github.com/gjh33/SurrealEngine/graphics"

	"github.com/vulkan-go/glfw/v3.3/glfw"
)

// Application represents top most information about a Surreal Application
type Application struct {
	Name     string          // The name of the application
	Version  SemanticVersion // The version of the application
	Contexts []gfx.Context   // Contexts being rendered to

	ApplicationEventsDispatcher // Application is an event dispatcher
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
	// TODO: Remove this code as it is test
	test := &listenerTester{}
	_ = application.Subscribe(test)
	_ = application.Dispatch(ApplicationStartupEvent{})
	context := &gfx.VulkanContext{}
	if err := context.Initialize(); err != nil {
		panic(err.Error())
	}
	window := context.Window()
	_ = window.SetResizable(true)
	_ = window.SetDecorated(true)
	if err := window.Create(); err != nil {
		panic(err.Error())
	}
	_ = window.Subscribe(test)
	// End of test

	_ = application.Dispatch(ApplicationInitializedEvent{})
	for !window.ShouldClose() {
		_ = application.Dispatch(ApplicationUpdateEvent{})

		// TODO: remove all below into main pipeline
		glfw.PollEvents()
	}
	if err := window.Close(); err != nil {
		panic(err.Error())
	}
	_ = application.Dispatch(ApplicationQuitEvent{})
	_ = application.Dispatch(ApplicationCleanedUpEvent{})
}

type listenerTester struct {
}

func (lt *listenerTester) OnWindowInitialized(e win.WindowInitializedEvent) {
	fmt.Println("Window Successfully Initialized!")
}

func (lt *listenerTester) OnWindowCreated(e win.WindowCreatedEvent) {
	fmt.Println("Window Successfully Created!")
}

func (lt *listenerTester) OnWindowShown(e win.WindowShownEvent) {
	fmt.Println("Window Shown")
}

func (lt *listenerTester) OnWindowHidden(e win.WindowHiddenEvent) {
	fmt.Println("Window Hidden")
}

func (lt *listenerTester) OnWindowFocusLost(e win.WindowFocusLostEvent) {
	fmt.Println("Window Focus Lost")
}

func (lt *listenerTester) OnWindowFocused(e win.WindowFocusedEvent) {
	fmt.Println("Window Focused")
}

func (lt *listenerTester) OnWindowClosed(e win.WindowClosedEvent) {
	fmt.Println("Window Closed")
}

func (lt *listenerTester) OnWindowCloseRequested(e win.WindowCloseRequestedEvent) {
	fmt.Println("Window Close Request Received")
}

func (lt *listenerTester) OnWindowIconified(e win.WindowIconifiedEvent) {
	fmt.Println("Window Iconified")
}

func (lt *listenerTester) OnWindowRestored(e win.WindowRestoredEvent) {
	fmt.Println("Window Restored")
}

func (lt *listenerTester) OnWindowResized(e win.WindowResizedEvent) {
	fmt.Printf("Window resized from (%v, %v) to (%v, %v)\n", e.OldSize.Width, e.OldSize.Height, e.NewSize.Width, e.NewSize.Height)
}

func (lt *listenerTester) OnWindowMoved(e win.WindowLocationChangedEvent) {
	fmt.Printf("Window Moved from (%v, %v) to (%v, %v)\n", e.OldLocation.X, e.OldLocation.Y, e.NewLocation.X, e.NewLocation.Y)
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
