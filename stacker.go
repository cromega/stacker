package stacker

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"
)

const packageName = "github.com/cromega/stacker"

func init() {
	setSigHandler()
	setHttpHandler()
}

func setSigHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Printf("\nInterrupt signal received:\n\n %s", getTrace())
		os.Exit(1)
	}()
}

func setHttpHandler() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, getTrace())
	})

	port := os.Getenv("STACKER_PORT")
	if port == "" {
		port = "6000"
	}
	go func() {
		http.ListenAndServe("localhost:"+port, nil)
	}()
}

func getTrace() string {
	trace := make([]byte, 102400)
	size := runtime.Stack(trace, true)
	if size > len(trace) {
		trace = make([]byte, size)
		runtime.Stack(trace, true)
	}

	return traceWithoutOwnGoroutines(string(trace[:size]))
}

func traceWithoutOwnGoroutines(trace string) (filteredTrace string) {
	goroutines := strings.Split(trace, "\n\n")
	for _, goroutine := range goroutines {
		if strings.Contains(goroutine, packageName) {
			continue
		}

		filteredTrace += goroutine + "\n\n"
	}

	return
}
