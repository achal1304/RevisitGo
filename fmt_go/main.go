package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Fprintln(os.Stdout, "writes formatted output to a writer")
	str := "hello, go world"
	upperStr := strings.ToUpper(str)
	fmt.Println("Upper case:", upperStr)
}
