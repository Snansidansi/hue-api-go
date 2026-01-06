package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Snansidansi/hue-api-go"
	"github.com/Snansidansi/hue-api-go/models"
	"github.com/joho/godotenv"
)

func main() {
	println("Hue API test.")
	println()

	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error while loading .env file: %v", err)
		os.Exit(1)
	}

	bridge := models.Bridge{
		Id:       os.Getenv("HUE_BRIDGE_ID"),
		IPAdress: os.Getenv("HUE_BRIDGE_IP"),
	}

	apiKey := os.Getenv("HUE_BRIDGE_USERNAME")
	client := hueapi.NewClient(bridge, apiKey, nil, true)

	// TestBridge(client)
	// TestLights(client)
	// TestRooms(client)
	// TestZones(client)
	// TestScenes(client)
	// TestDevicePower(client)
	// TestDeviceSoftwareUpdate(client)
	// TestEntertainmentConfiguration(client)
	// TestEntertainment(client)
	// TestBridgeHome(client)
	// TestDevice(client)
	// TestGroupedLight(client)
	// TestEventStream(client)
	TestButtons(client)
}

func printStructFormatted(data any) {
	bytes, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(bytes))
}

func printHueResponse[T any](
	resp *models.HueResponse[T],
	err error,
	title string,
	printData bool,
) {
	fmt.Println()
	fmt.Println("--> " + title)

	if err != nil {
		fmt.Printf("Error getting rooms: %v\n", err)
		return
	}

	if len(resp.Errors) > 0 {
		fmt.Print("Hue error: ")
		printStructFormatted(resp.Errors)
	}

	if printData {
		printStructFormatted(resp.Data)
	}
}

func printHueActionResponse(
	resp *models.HueActionResponse,
	err error,
	title string,
	printData bool,
) {
	fmt.Println()
	fmt.Println("--> " + title)

	if err != nil {
		fmt.Printf("Error getting rooms: %v\n", err)
		return
	}

	if len(resp.Errors) > 0 {
		fmt.Print("Hue error: ")
		printStructFormatted(resp.Errors)
	}

	if printData {
		printStructFormatted(resp.Data)
	}
}
