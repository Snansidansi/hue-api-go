package main

import (
	"fmt"
	"github.com/snansidansi/hueapi"
	"github.com/snansidansi/hueapi/builders"
)

func TestGroupedLight(client *hueapi.Client) {
	// TestGetAllGroupedLights(client)
	// TestUpdateGroupedLight(client)
}

func TestGetAllGroupedLights(client *hueapi.Client) {
	hueResp, err := client.GroupedLight.GetAllGroupedLights()
	printHueResponse(hueResp, err, "Get all grouped lights", true)
}

func TestUpdateGroupedLight(client *hueapi.Client) {
	hueResp, err := client.GroupedLight.GetAllGroupedLights()
	if err != nil || len(hueResp.Data) == 0 {
		fmt.Println("No Grouped Lights found for testing.")
		return
	}

	targetID := hueResp.Data[0].ID
	fmt.Printf("Test Update a Grouped Light ID: %s\n", targetID)

	builder := builders.NewGroupedLightBuilder()

	update := builder.On(true).
		SignalingRGB("on_off_color", 2000, 23, 23, 23).
		Build()
	//  Brightness(50).
	// Dynamics(2000).

	actionResp, err := client.GroupedLight.UpdateGroupedLight("5ef1bd6e-a560-4414-9339-ce4cc93cc997", update)
	printHueActionResponse(actionResp, err, "Update Grouped Light", true)
}
