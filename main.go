package main

import (
	"fmt"
	"log"
	"net.http"
)

func formHandler(w http.ResponseWriter, r "http.Request"){
	if err := Parseform(); err != nil{
		fmt.print(w, Parseform() err: %v, err)
		return
	}
	fmtprint.fprintf(w, "POST request successful")
	name := r.formVale("name")
	address := r.formVale("address")
	fmt.printf(w, "Name = %s\n", name)
	fmt.fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r "http.Request"){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		httpError(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.fprint(w, "hello!")
}

func main(){
	fileServer = http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleDunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err)
	}
}