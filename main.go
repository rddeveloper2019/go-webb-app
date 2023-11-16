package main

import (
	"fmt"
	"net/http"
)

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	user := "Alex"
	responseMsg := fmt.Sprintf("<h1>Hello %s! Lenslocked</h1>", user)
	w.Write([]byte(responseMsg))

}
func main() {
	http.HandleFunc("/", HandleFunc)
	fmt.Println("Server started at http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
