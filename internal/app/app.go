package app

import (
	notifysend "github.com/eterline/gotify-catcher/internal/notification/notify-send"
	"github.com/eterline/gotify-catcher/internal/process"
	"github.com/eterline/gotify-catcher/pkg/logging"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

type NotifyProvider interface {
	Push (title, message string) error
}


type Config struct {
	Host 	string
	Token 	string
	_ 		struct {}
}

type App struct {
	Connection *process.Process
	Alarm 		NotifyProvider
	_ struct {}
}

func Init(config Config) *App {
	log = logging.ReturnEntry().Logger
	ws := process.SetServerWSS(config.Host, config.Token)

	conn, err := process.Connect(ws)
	if err != nil {
		log.Fatalf("Init error: %v", err)
	}

	return &App{
		Connection: conn,
		Alarm: notifysend.Init("Notify Checker", "", 1),
	}
}

func (app *App) Run() {
	log.Info("starting app")

	app.processConn()
}

func (app *App) Stop() {
	log.Info("quitting app")
	app.Connection.Close()
}

func (app *App) processConn() {
	for {
		if _, err := app.Connection.Read(); err != nil {
			log.Errorf("(websocket) connection error: %v", err)
			continue
		}

		msg, err := app.Connection.LastDecode()
		if err != nil {
			log.Errorf("message decode error: %v", err)
			continue
		}

		log.Infof("got notification titled: %s", msg.Title)
		if err := app.Alarm.Push(msg.Title, msg.Message) ;err != nil {
			log.Errorf("failed to push notification: %v", err)
			continue
		}
	}
}



