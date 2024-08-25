package main

import (
	"hayai/config"
	"hayai/wolfx"
)

func main() {
	config.Init()
	wolfx.Listen()
}
