package main

import (
	"github.com/imsugeno/gin/internal/handler"
)

func main() {
	r := handler.NewGinEngin()
	r.Run(":3000")
}
