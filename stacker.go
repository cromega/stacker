package stacker

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
)

func init() {
	setSigHandler()
	setHttpHandler()
}

func setSigHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		panic("Interrupt signal received.")
	}()
}

func setHttpHandler() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		trace := make([]byte, 1024*1024)
		len := runtime.Stack(trace, true)
		fmt.Fprintf(w, string(trace[:len]))
	})

	port := os.Getenv("STACKER_PORT")
	if port == "" {
		port = "6000"
	}
	go func() {
		http.ListenAndServe("localhost:"+port, nil)
	}()
}
