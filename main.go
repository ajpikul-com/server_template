package main

import (
	"net/http"

	"github.com/ajpikul-com/server_utils"
	"github.com/ajpikul-com/simpauth"
)

// produceMain demonstrates creating a production path (example.com) and a staging path
// (stage.example.com) pointed to reasonable directories. It's a static server. It uses
// simpauth to protect the stage and to protect a directory w/in the production path.
func produceMain(serveMux *http.ServeMux) {

	var err error
	productionFileServer := new(simpauth.Bouncer)
	err = productionFileServer.Init(http.FileServer(http.Dir("/var/www/example.com")))
	if err != nil {
		panic(err)
	}
	serveMux.Handle("example.com/", productionFileServer)
	serveMux.Handle("www.example.com/", productionFileServer)

	stageFileServer := new(simpauth.Bouncer)
	err = stageFileServer.Init(http.FileServer(http.Dir("/var/www/stage.example.com")))
	if err != nil {
		panic(err)
	}

	serveMux.Handle("stage.example.com/", stageFileServer)

}

func main() {
	var err error
	serveMux := http.NewServeMux()

	// Attach paths to serveMux
	produceMain(serveMux)
	produceString(serveMux)

	defaultLogger.Info("Running serverMux")

	// Run serveMux
	go func() {
		// All http are changed to https using one of my utilities.
		redirect := server_utils.RedirectSchemeHandler("https", http.StatusMovedPermanently)
		serverHTTP := &http.Server{
			Addr:    ":http",
			Handler: redirect,
		}

		err := serverHTTP.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	serverHTTPS := &http.Server{
		Addr:    ":https",
		Handler: serveMux,
	}

	err = serverHTTPS.ListenAndServeTLS(*fullChain, *privKey)
	if err != nil {
		panic(err)
	}
}
