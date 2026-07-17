package main

import (
	"fmt"
	"net/http"
)

type Task struct{
	ID int
	Title string
	Status string
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Hello ChatGpt!!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server Running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
