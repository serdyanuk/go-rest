package main

import (
	"github.com/serdyanuk/go-rest/config"
	"github.com/serdyanuk/go-rest/internal/app"
)

func main() {
	cfg := config.Get()
	a := app.New(cfg)
	a.Run()
}
