package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Golang Sample, Selamat datang! \nApp Version: v1.2 \nBelajar Jenkins Deploy Golang")
}

func envPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List of Environment Variables:")
	for _, env := range os.Environ() {
		fmt.Fprintf(w, "- "+env+"\n")
	}
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/env", envPage)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func main() {
	handleRequests()
}
