package main

import (
	"hayai/config"

	"hayai/render"
	"hayai/waves"
	"hayai/wolfx"
)

func main() {
	config.Init()
	waves.Init()
	render.Init()
	wolfx.Listen()
}
