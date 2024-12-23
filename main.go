package main

import (
	"fmt"
	"net/http"
	"web-server/handlers"
)

func main(){

	http.HandleFunc("/hello",handlers.HelloHandler)

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/",fs)

	fmt.Println("Starting server on : 8080")
	err:=http.ListenAndServe(":8080" ,nil)
	if err != nil{
		fmt.Println("error starting server:", err)
	}

}