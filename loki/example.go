// https://github.com/evalphobia/logrus_fluent

package main

import (
	"context"
	"errors"

	"github.com/evalphobia/logrus_fluent"
	"github.com/sirupsen/logrus"
)

func main() {
	hook, err := logrus_fluent.NewWithConfig(logrus_fluent.Config{
		Host: "localhost",
		Port: 24224,
	})
	if err != nil {
		panic(err)
	}

	// set custom fire level
	hook.SetLevels([]logrus.Level{
		logrus.PanicLevel,
		logrus.ErrorLevel,
	})

	// set static tag
	hook.SetTag("original.tag")

	// ignore field
	hook.AddIgnore("context")

	// filter func
	hook.AddFilter("error", logrus_fluent.FilterError)

	logrus.AddHook(hook)

	// write log
	logging(context.Background())
}

func logging(ctx context.Context) {
	logrus.WithFields(logrus.Fields{
		"app":     "dev",
		"value":   "some content...",
		"error":   errors.New("unknown error"), // this field will be applied filter function in the hook.
		"context": ctx,                         // this field will be ignored in the hook.
	}).Error("error message")
}
