package main

import (
	"os"
	"time"

	"example.com/hello/mocking"
)

// This main is following the Dependency Injection chapter
// At first I was importing the module from dependencyinjection module/directory
// func Greet(writer io.Writer, name string) {
// 	fmt.Fprintf(writer, "Hello, %s", name)
// }

// func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
// 	Greet(w, "World")
// }

// func main() {
// 	// Greet(os.Stdout, "Andrew")
// 	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreetHandler)))
// }

// Next stuff is for the Mocking chapter...
// Using the importing approach to uhm,... it feels more orderly I guess
func main() {
	sleeper := &mocking.ConfigurableSleeper{1 * time.Second, time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
}
