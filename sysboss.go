package main

import (
	"net/http"

	"github.com/ajpikul-com/sutils"
)

func produceSysBoss(serveMux *http.ServeMux) {
	SysBossServer, err := sutils.NewSingleHostReverseProxy(globalConfig.SysBoss)
	if err != nil {
		panic(err.Error())
	}
	serveMux.Handle("sysboss.ajpikul.com/", SysBossServer)
}
