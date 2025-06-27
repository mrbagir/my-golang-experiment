package main

import (
	"net/http"
)

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	// Ganti dengan URL tujuan
	http.Redirect(w, r, "https://google.com", http.StatusFound)
}

func main() {
	http.HandleFunc("/redirect", redirectHandler)

	// Menjalankan server di port 8080
	http.ListenAndServe(":8080", nil)
}
