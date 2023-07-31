package main

import (
	"github.com/ajpikul-com/ilog"
	"github.com/ajpikul-com/uwho"
	"github.com/ajpikul-com/uwho/googlelogin"
	"github.com/ajpikul-com/uwho/usersessioncookie"
)

var defaultLogger ilog.LoggerInterface

func loggerInit() {
	if defaultLogger == nil {
		//defaultLogger = &ilog.ZapWrap{Sugar: false}
		defaultLogger = &ilog.SimpleLogger{}
		defaultLogger.Init()
		defaultLogger.Info("Logging started")
	}
	uwho.SetDefaultLogger(defaultLogger)
	googlelogin.SetDefaultLogger(defaultLogger)
	usersessioncookie.SetDefaultLogger(defaultLogger)
}
