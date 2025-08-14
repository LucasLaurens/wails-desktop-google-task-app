package main

import (
	"calandar-desktop-task/internal/errors"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:         "calandar-desktop-task",
		Width:         600,
		Height:        60,
		DisableResize: true,
		Frameless:     true,
		AlwaysOnTop:   false,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		Windows: &windows.Options{
			WebviewIsTransparent:              true,
			WindowIsTranslucent:               true,
			DisableWindowIcon:                 true,
			IsZoomControlEnabled:              false,
			ZoomFactor:                        0,
			DisablePinchZoom:                  false,
			DisableFramelessWindowDecorations: false,
		},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
		},
	})

	errors.Fatal(
		"The calandar desktop task return an error and cannot be started correctly: %v",
		errors.FatalError{
			Err:  err,
			Args: []interface{}{},
		},
	)
}
