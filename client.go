package main

import (
	"flag"
)

func main() {
	var debug = flag.Bool("d", false, "If this flag is applied, output debug information.")
	var configuration = flag.Bool("c", false, "If this flag is applied, open configuration mode.")
	flag.Parse()

	if *debug {
		// Register a logger.
	}

	if *configuration {
		Configure()
	}
}
