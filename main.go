package main

import "net/http"

func healthz(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	writer.WriteHeader(200)
	writer.Write([]byte("OK"))
}

func main() {
	serveMux := http.NewServeMux()
	serveMux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir("."))))
	serveMux.HandleFunc("/healthz", healthz)

	server := http.Server{}
	server.Addr = ":8080"
	server.Handler = serveMux

	server.ListenAndServe()
}
