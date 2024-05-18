package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/name":
		fmt.Fprintln(w, "Name: vijay")
	case "/rollnumber":
		fmt.Fprintln(w, "Roll Number: 12345")
	case "/favoritecolor":
		fmt.Fprintln(w, "Favorite Color: Blue")
	default:
		http.NotFound(w, r)
	}
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Printf("Server running (port=8080)\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
