package dependencyinjection

import (
	"bytes"
	"fmt"
)

func Greet(buffer *bytes.Buffer, name string) {
	// fmt.Printf("Hello, %s", name)
	fmt.Fprintf(buffer, "Hello, %s", name)
}
