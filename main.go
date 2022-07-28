package main

import (
	"capi/app"
	"capi/logger"
	"log"
)

func main() {
	log.Println("starting application...")
	logger.Log.Info("starting application....")
	app.Start()
}
