package main

import (
	"encoding/json"
	"fmt"
	"github.com/mewbak/gopass"
	"github.com/scanbadge/client/utility"
	"io/ioutil"
)

type configuration struct {
	URL        string `json:"url"`
	DeviceName string `json:"device_name"`
	DeviceKey  string `json:"device_key"`
}

func configure() {
	var config configuration
	fmt.Println("Configure ScanBadge Client")

	fmt.Println("ScanBadge API URL:")
	fmt.Scanln(&config.URL)
	if config.URL == "" {
		config.URL = "https://scanbadge.xyz/api/v1"
		fmt.Println("Using default API url: " + config.URL)
	}

	fmt.Println("Device name:")
	fmt.Scanln(&config.DeviceName)
	if config.DeviceName == "" {
		fmt.Println("Device name is required.")
		return
	}

	key := ""
	if key, err := gopass.GetPass("Device key:\n"); err != nil {
		if key == "" || len(key) < 20 || len(key) > 8192 {
			fmt.Println("Device key is required. It must be 20 characters in length and cannot exceed 8192 characters")
			return
		}

		fmt.Println(err)
		return
	}

	config.DeviceKey = utility.HashPassword(key)

	// We have all required information, write it to a client-config.json file.
	b, err := json.Marshal(config)

	if err != nil {
		fmt.Println(err)
		return
	}

	ioutil.WriteFile("client-config.json", b, 0775)
}
