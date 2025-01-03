package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// This main is following the Dependency Injection chapter
// At first I was importing the module from dependencyinjection module/directory

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "World")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreetHandler)))
}
