package logger

import (
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
)

func Init() {
	log.SetFormatter(&log.JSONFormatter{})
	debugLogFile, err := os.OpenFile("debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		log.Error(err)
	}
	errorLogFile, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		log.Error(err)
	}
	if os.Getenv("RUN_MODE") == "release" {
		log.SetOutput(ioutil.Discard)
		log.AddHook(&writer.Hook{
			Writer: errorLogFile,
			LogLevels: []log.Level{
				log.PanicLevel,
				log.FatalLevel,
				log.ErrorLevel,
			},
		})
		log.AddHook(&writer.Hook{
			Writer: debugLogFile,
			LogLevels: []log.Level{
				log.PanicLevel,
				log.FatalLevel,
				log.ErrorLevel,
				log.WarnLevel,
				log.InfoLevel,
				log.DebugLevel,
			},
		})
	}
}
