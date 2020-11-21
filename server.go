package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/home":
		fmt.Fprintf(w, "welcome to home\n")
	case "/profile":
		fmt.Fprintf(w, "This is your profile\n")
	default:
		fmt.Fprintf(w, "Page not found\n")

	}

}

func main() {

	http.HandleFunc("/", handleFunc)
	http.ListenAndServe(":8090", nil)
}
