package examples

import (
	"github.com/gjh33/SurrealEngine/core/app"
)

// CreateABasicApplication is the entry point for the Creating An Application example
func CreateABasicApplication() {
	application := app.New("Basic Application", "1.0.0")
	application.Start()
}
