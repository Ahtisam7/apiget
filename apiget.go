package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handleRequests() {
	http.HandleFunc("/", homePage)
	//	log.Println(mux.Vars())
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	s := r.RequestURI
	f, err := os.OpenFile("text.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(s[1:] + "\n"); err != nil {
		log.Println(err)
	}
	log.Println(s[1:])
	fmt.Println("Endpoint Hit: homePage")
}
func main() {

	handleRequests()
}
