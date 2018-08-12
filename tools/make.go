package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	out = flag.String("out", "", "File to write to")
)

func main() {
	flag.Parse()

	output := os.Stdout
	if *out != "" {
		output, err := os.Open(*out)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer output.Close()
	}

	Out(output, 20, 10, 5)
}
