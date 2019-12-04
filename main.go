package main

import (
	"golang.org/x/net/http2"
	"log"
	"net/http"
	"ocr-service/src/config"
	"ocr-service/src/constant"
)

func main() {
	mux := config.MakeHttpHandler()
	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	constant.Dbinit()
	http2.ConfigureServer(s, &http2.Server{})
	log.Fatal(s.ListenAndServe())
}
