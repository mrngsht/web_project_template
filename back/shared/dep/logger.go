package dep

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	l := logrus.New()

	// see here: https://github.com/sirupsen/logrus/blob/0c8c93fe4d2fb9013b83ae5f3151608f69f562ca/logrus.go#L25
	if level, err := logrus.ParseLevel(GetEnv("LOG_LEVEL", "debug")); err != nil {
		panic(err)
	} else {
		l.SetLevel(level)
	}

	l.SetFormatter(&logrus.TextFormatter{})
	l.SetOutput(os.Stdout)
	l.SetReportCaller(true)
	Log = l
}
