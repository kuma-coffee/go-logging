package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// FOR LOGRUS
func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

func main() {
	// APEX IMPLEMENTATIONS
	// ctx := log.WithFields(log.Fields{
	// 	"file": "something.png",
	// 	"type": "image/png",
	// 	"user": "tobi",
	// })

	// for range time.Tick(time.Millisecond * 200) {
	// 	ctx.Info("upload")
	// 	ctx.Info("upload complete")
	// 	ctx.Warn("upload retry")
	// 	ctx.WithError(errors.New("unauthorized")).Error("upload failed")
	// 	ctx.Errorf("failed to upload %s", "img.png")
	// }

	// ZEROLOG IMPLEMENTATIONS
	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// log.Debug().
	// 	Str("Scale", "833 cents").
	// 	Float64("Interval", 833.09).
	// 	Msg("Fibonacci is everywhere")

	// log.Debug().
	// 	Str("Name", "Tom").
	// 	Send()

	// Output: {"level":"debug","Scale":"833 cents","Interval":833.09,"time":1562212768,"message":"Fibonacci is everywhere"}
	// Output: {"level":"debug","Name":"Tom","time":1562212768}

	// LOGRUS IMPLEMENTATIONS
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")

	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
}
