package main

import (
	"golang.org/x/net/http2"
	"net/http"
	"ocr-service/config"
	"ocr-service/roule"
)

func main() {
	config.InitLoger()
	config.Dbinit()

	mux := roule.MakeHttpHandler()
	http2.VerboseLogs = true
	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	http2.ConfigureServer(s, &http2.Server{})
	config.Logger.Fatal(s.ListenAndServe())

}
