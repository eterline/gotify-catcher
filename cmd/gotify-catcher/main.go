package main

import (
	"flag"

	"github.com/eterline/gotify-catcher/internal/app"
	"github.com/eterline/gotify-catcher/pkg/logging"
)

var (
	logPath, gotifyURL, gotifyToken string
)

func init () {
	flag.StringVar(&logPath, "log", "", "Setting log file path.")
	flag.StringVar(&gotifyToken, "token", "", "Setting gotify client token.")
	flag.StringVar(&gotifyURL, "gotify", "example-gotify.com", "Setting gotify ws path.")
}

func main () {
	flag.Parse()

	logging.InitLogger(logPath, "gotify-checker.log")
	validGotify()

	conf := app.Config{
		Host: gotifyURL,
		Token: gotifyToken,
	}
	app := app.Init(conf)

	app.Run()
}

func validGotify() {
	log := logging.ReturnEntry()
	if gotifyToken == "" || gotifyURL == "ws://example-gotify.com" {
		log.Fatal("Gotify settings not configured")
	}
}