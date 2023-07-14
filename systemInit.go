package main

import (
	"flag"

	"github.com/ajpikul-com/ilog"
)

var fullChain = flag.String("fullChain", "", "the directory for the full chain")
var privKey = flag.String("privKey", "", "the directory for the private key")

func init() {
	flag.Parse()
	if *fullChain == "" {
		panic("You must set a directory for the full chain, see -h")
	}
	if *privKey == "" {
		panic("You must set a directory for the private key, see -h")
	}
}

var defaultLogger ilog.LoggerInterface

func init() {
	if defaultLogger == nil {
		defaultLogger = &ilog.ZapWrap{Sugar: false}
		defaultLogger.Init()
		defaultLogger.Info("Logging started")
	}
}
