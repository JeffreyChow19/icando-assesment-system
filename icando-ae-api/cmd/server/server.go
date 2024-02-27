package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"icando/app"
	"os"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	app.StartApp()
}
