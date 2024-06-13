package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("./content/")))
	mux.Handle("/assets/logo.png", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets/"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
