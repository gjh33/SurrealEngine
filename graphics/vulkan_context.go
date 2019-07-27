package graphics

import (
	"github.com/gjh33/SurrealEngine/graphics/win"
)

// VulkanContext implements the vulkan graphics context for all OS supported
type VulkanContext struct {
	BaseContext
}

// Initialize implements the Context interface
func (vkcxt *VulkanContext) Initialize() error {
	// TODO: Actually init vulkan context
	vkwin := &win.VulkanWindow{}
	if err := vkwin.Initialize(); err != nil {
		return err
	}
	vkcxt.State.Window = vkwin
	return nil
}

// Window implements the Context interface
func (vkcxt *VulkanContext) Window() win.Window {
	return vkcxt.State.Window
}

// IsInitialized implements the Context interface
func (vkcxt *VulkanContext) IsInitialized() bool {
	return vkcxt.Initialized
}
