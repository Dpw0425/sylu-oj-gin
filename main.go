package main

import (
	_ "sylu-oj-gin/docs"
	"sylu-oj-gin/internal/app"
	"sylu-oj-gin/internal/app/config"
)

// @title sylu-oj
// @version 1.0
// @description This is a backend server for sylu-oj.
// @BasePath /api/v1
func main() {
	r := app.Init()
	r.Run(":" + config.CONFIG.Http.Port)
}
