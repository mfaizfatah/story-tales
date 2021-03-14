package logger

import (
	"context"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mfaizfatah/story-tales/app/adapter"
	"github.com/sirupsen/logrus"
)

// EndRecord for initialize context first time
func EndRecord(ctx context.Context, response string, statuscode int) {
	var level string

	v, ok := ctx.Value(logKey).(*Data)
	if ok {
		t := time.Since(v.TimeStart)

		if statuscode >= 200 && statuscode < 400 {
			level = "INFO"
		} else if statuscode >= 400 && statuscode < 500 {
			level = "WARN"
		} else {
			level = "ERROR"
		}

		v.StatusCode = statuscode
		v.Response = response
		v.ExecTime = t.Seconds()

		Output(v, level)
	}
}

// Output for output to terminal
func Output(out *Data, level string) {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	if level == "ERROR" {
		logrus.WithField("data", out).Error("apps")
	} else if level == "INFO" {
		logrus.WithField("data", out).Info("apps")
	} else if level == "WARN" {
		logrus.WithField("data", out).Warn("apps")
	}

	go adapter.MongoDatabase().Collection("logs").InsertOne(context.Background(), out)
}
