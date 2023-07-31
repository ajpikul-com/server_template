package main

import (
	"encoding/json"
	"flag"
	"os"
)

type config struct {
	SysBoss        string
	FullChain      string
	PrivateKey     string
	GoogleClientID string
}

// b, err := json.Marshal(instance of config)
var globalConfig config

func initConfig() {
	configFile, err := os.Open(`WHERE IS YOUR CONFIG FILE`)
	if err != nil {
		panic(err.Error())
	}
	defer configFile.Close()
	configDecoder := json.NewDecoder(configFile)
	if err != nil {
		panic(err.Error())
	}
	//err := json.Unmarshal(bytes, &config)
	err = configDecoder.Decode(&globalConfig)
	if err != nil {
		panic(err.Error())
	}
}

var fullChainFlag = flag.String("fullChain", "", "the directory for the full chain")
var privKeyFlag = flag.String("privKey", "", "the directory for the private key")

func init() {
	loggerInit()
	initConfig()
	flag.Parse()
	if *fullChainFlag != "" {
		globalConfig.FullChain = *fullChainFlag
	}
	if *privKeyFlag != "" {
		globalConfig.PrivateKey = *privKeyFlag
	}
	if globalConfig.SysBoss == "" || // TODO this sucks
		globalConfig.FullChain == "" ||
		globalConfig.PrivateKey == "" ||
		globalConfig.GoogleClientID == "" {
		panic("Config not fully filled out!")
	}
}
