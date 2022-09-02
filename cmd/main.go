package main

import (
	"os"

	"github.com/Kolyan4ik99/blog-app/internal/app"
	"github.com/Kolyan4ik99/blog-app/internal/logger"
)

func init() {
	logger.InitLogger(os.Stdout)
}

func main() {
	app.Run()
}
