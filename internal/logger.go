package internal

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	Log.Out = os.Stdout

	Log.SetLevel(logrus.InfoLevel)
}
