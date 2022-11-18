package main

import (
	_ "b0go/apps/app1"
	_ "b0go/apps/docs"

	"b0go/core/engine"
	_ "b0go/core/gateway"
)

func main() {
	engine.Run("config.ini")
	select {}
}
