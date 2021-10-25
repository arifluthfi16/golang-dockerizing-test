package main

import (
	"fmt"
	"net/http"
)

func main() {
	request()
}

func homePage(response http.ResponseWriter, r *http.Request) {
	fmt.Printf("Welcome!")
	fmt.Println("Endpoint Hit: Homepage")
}

func aboutPage(response http.ResponseWriter, r *http.Request) {
	who := "X"

	fmt.Fprintf(response, "A little bit about...")
	fmt.Println("Endpoint Hit : ", who)
}

func request() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", aboutPage)
	fmt.Println("SERVER IS STARTING ON PORT : 4016")
	fmt.Println("http://localhost:4016/")
	fmt.Println()
	http.ListenAndServe(":4016", nil)
}