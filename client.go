package main

import (
	"flag"
	"fmt"
)

var isDebug, isConfigure bool

func init() {
	flag.BoolVar(&isDebug, "d", false, "If this flag is applied, output debug information.")
	flag.BoolVar(&isConfigure, "c", false, "If this flag is applied, open configuration mode.")
	flag.Parse()
}

func main() {
	fmt.Println("[ScanBadge] Starting ScanBadge client...")

	if isDebug {
		fmt.Println("[ScanBadge] Running in debug mode...")
		// TODO: implement a logger (either via API or local - or both).
	}

	if isConfigure {
		configure()
	}
}
