package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func init() {
	// disable timestamps in the log format -- that's handled by syslog
	log.SetFlags(0)
}

func withLogs(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func getAddr() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return fmt.Sprintf(":%s", port)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("assets")))
	addr := getAddr()
	log.Printf("Listening on %s", addr)
	if err := http.ListenAndServe(addr, withLogs(http.DefaultServeMux)); err != nil {
		log.Fatal(err)
	}
}
