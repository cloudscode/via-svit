package main

import (
	"embed"
	"github.com/cloudcodes/via-svit/backend/database"
	"github.com/cloudcodes/via-svit/backend/router"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// 初始化数据库连接
	database.InitDB()

	// Create an instance of the app structure
	app := NewApp()
	// 启动 Gin 服务器
	go func() {
		r := router.SetupRouter()
		// 监听端口，可根据需要调整
		r.Run(":8080")
	}()
	// Create application with options
	err := wails.Run(&options.App{
		Title:  "via-svit",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
