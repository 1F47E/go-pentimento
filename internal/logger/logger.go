package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

type LoggerEntry struct {
	logrus.Entry
}

func init() {
	Log.Out = os.Stdout

	if os.Getenv("DEBUG") == "1" {
		Log.Level = logrus.DebugLevel
	} else {
		Log.Level = logrus.InfoLevel
	}

	Log.Formatter = &logrus.TextFormatter{
		ForceColors:   true,
		ForceQuote:    true,
		FullTimestamp: false,
		// DisableLevelTruncation: false,
	}
}
