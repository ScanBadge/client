package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Configuration struct {
	Url string `json:"url"`
	Key string `json:"url"`
}

func Configure() {
	// Firstly, collect all required information for configuring the ScanBadge client:
	// - Url of API
	// - Authentication private key
	var config Configuration
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Configure ScanBadge Client")
	fmt.Println("URL of ScanBadge API (leave empty for default value):")
	config.Url, _ = reader.ReadString('\n')

	if config.Url != "\n" && len(config.Url) <= 8192 {
		// Use default value
		config.Url = "https://scanbadge.xyz/api/v1"
	}

	fmt.Println("Authentication key for this device:")
	config.Key, _ = reader.ReadString('\n')

	if config.Key != "\n" && len(config.Key) > 20 && len(config.Key) <= 8192 {
		// Key must always have a value, and cannot be outside the specified boundaries.
		panic("Authentication key is required. You can create an authentication key with the API or the front-end.")
	}

	// We have all required information, write it to a client-config.json file.
	b, err := json.Marshal(config)

	if err == nil {
		ioutil.WriteFile("client-config.json", b, 0775)
	}
}
