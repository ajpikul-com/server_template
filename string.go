package main

import (
	//"github.com/ajpikul-com/simpauth"
	"net/http"
)

func produceString(serveMux *http.ServeMux) {
	serveMux.Handle("strings.ajpikul.com/", stringMain)
}

// StringToHandler is a string but fulfils the http.Handler interface as it has ServeHTTP function.
type StringHandler string

func (sh StringHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(sh))
	if err != nil {
		panic(err)
	}
	w.(http.Flusher).Flush()
}

var stringMain StringHandler = `<html>
	<body>
		I live in RAM as a string.
	</body>
</html>`
