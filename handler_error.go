package main

import "net/http"

func handleErr(w http.ResponseWriter, r *http.Request) {

	responseWithError(w, 400, "Failed test endpoint")
}
