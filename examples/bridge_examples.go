package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Snansidansi/hue-api-go"
	"github.com/Snansidansi/hue-api-go/models"
)

func TestBridge(client *hueapi.Client) {
	// discoverBridges()
	// registerInBridge(&client.Bridge)
}

func discoverBridges() {
	bridges, _ := hueapi.DiscoverBridges(http.DefaultClient)
	fmt.Println("--> Discover bridges")
	printStructFormatted(bridges)
}

func registerInBridge(bridge *models.Bridge) {
	fmt.Println()
	fmt.Println("--> Register in bridge")
	devicename := "HueApi-testing"
	generateClientKey := true
	username, clientkey, err := hueapi.RegisterBridge(http.DefaultClient, bridge, devicename, generateClientKey)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		var hueErr *models.HueError
		if errors.As(err, &hueErr) {
			if hueErr.Type == models.ErrLinkButtonNotPressed {
				fmt.Println("Link button not pressed")
			}
		}
		return
	}

	fmt.Println("Registration successful:")
	fmt.Printf("  Username: %s\n", username)
	fmt.Printf("  ClientKey: %s\n", clientkey)
}
