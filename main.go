package main

import (
	"net/http"
	"time"

	"github.com/ajpikul-com/sutils"
	"github.com/ajpikul-com/uwho"
	"github.com/ajpikul-com/uwho/googlelogin"
	"github.com/ajpikul-com/uwho/usersessioncookie"
)

func produceMain(serveMux *http.ServeMux) {
	productionFileServer := http.FileServer(http.Dir("/var/www/ajpikul.com"))
	serveMux.Handle("ajpikul.com/", productionFileServer)
	serveMux.Handle("www.ajpikul.com/", productionFileServer)
}

func produceStage(serveMux *http.ServeMux) {
	// ** SET UP UWHO
	cookieSessions := usersessioncookie.New(7*24*time.Hour, globalConfig.PrivateKey)
	googleIdent := googlelogin.New(globalConfig.GoogleClientID)

	loginScreen := googleIdent.DefaultLoginPortal("/login")
	stageFileServer := uwho.New(http.FileServer(http.Dir("/var/www/stage.ajpikul.com")),
		&googlelogin.DefaultLoginResult{},
		loginScreen,
		&googlelogin.RedirectHome{},
		"/login",
		"/logout",
		&stageFactory{},
	)
	stageFileServer.AddIdentifier(googleIdent)
	stageFileServer.AttachSessionManager(cookieSessions)
	// ** CREATE HOOKS FOR UWHO
	// State can't access session directly, session not configured to refresh (should be), so we must put it in a hook.
	refreshSession := func(stateCoord uwho.ReqByCoord, w http.ResponseWriter, r *http.Request) error {
		if !stateCoord.(*stageState).failedAuth {
			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			w.Header().Set("Pragma", "no-cache")
			w.Header().Set("Expires", "0")
		}
		return nil
	}
	stageFileServer.SetHooks(&stageFileServer.Hooks.AboutToLoad, []uwho.Hook{&refreshSession})

	// FINALLY ATTACH MUX
	serveMux.Handle("stage.ajpikul.com/", &stageFileServer)
}

func main() {
	var err error
	serveMux := http.NewServeMux()

	produceMain(serveMux)
	produceStage(serveMux)
	produceSysBoss(serveMux)
	//produceDev(serveMux)
	//produceWiki(serveMux)

	go func() {
		// Redirect this
		redirect := sutils.RedirectSchemeHandler("https", http.StatusMovedPermanently)
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
		Handler: serveMux, // you could wrap servemux and then call it's ServeHTTP to always serve a certain header
	}

	err = serverHTTPS.ListenAndServeTLS(globalConfig.FullChain, globalConfig.PrivateKey)
	if err != nil {
		panic(err)
	}
}
