package main

import (
	"gin/internal/handler"
)

func main() {
	r := handler.NewGinEngin()
	r.Run(":3000")
}
