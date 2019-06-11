
package log

import "github.com/sirupsen/logrus"

func setupLogging() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)
	log.SetFormatter(&log.TextFormatter{
		ForceColors: true,
	})

}