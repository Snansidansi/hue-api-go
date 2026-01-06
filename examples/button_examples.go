package main

import (
	"fmt"

	"github.com/snansidansi/hueapi"
)

func TestButtons(client *hueapi.Client) {
	// TestGetAllButtons(client)
	// TestGetButtonByID(client)
}

func TestGetAllButtons(client *hueapi.Client) {
	hueResp, err := client.Button.GetAllButtons()
	printHueResponse(hueResp, err, "Get all buttons", true)
}

func TestGetButtonByID(client *hueapi.Client) {
	hueResp, err := client.Button.GetAllButtons()
	printHueResponse(hueResp, err, "Get all buttons (to pick one for ID test)", false)

	if err != nil || len(hueResp.Data) == 0 {
		fmt.Println("No buttons found to test GetButtonByID.")
		return
	}

	targetID := hueResp.Data[0].ID
	buttonResp, err := client.Button.GetButtonByID(targetID)
	printHueResponse(buttonResp, err, fmt.Sprintf("Get button by ID: %s", targetID), true)
}
