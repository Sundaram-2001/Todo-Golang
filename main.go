package main

import (
	"api/main_api"
	"net/http"
)

func main() {

	// http.HandleFunc("/hello", main_api.Hello)
	// http.HandleFunc("/about", main_api.About)
	http.HandleFunc("/newtask", main_api.CreateTask)
	http.HandleFunc("/showtask", main_api.ShowTasks)
	port := "8090"
	http.ListenAndServe(":"+port, nil)
	// fmt.Printf("Server listening on %s", port)
}
