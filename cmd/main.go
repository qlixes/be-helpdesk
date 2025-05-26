package main

import "net/http"

func main() {
	mux := http.ServeMux()

	http.ListenAndServe(":8000", mux)
}
