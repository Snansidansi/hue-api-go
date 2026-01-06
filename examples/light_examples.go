package main

import (
	"fmt"

	"github.com/Snansidansi/hue-api-go"
)

func TestLights(client *hueapi.Client) {
	TestGetAllLights(client)
	TestGetLightByID(client)
	TestSetLight(client)
}

func TestGetAllLights(client *hueapi.Client) {
	hueResp, err := client.Lights.GetAllLights()
	printHueResponse(hueResp, err, "Get all lights", true)
}

func TestGetLightByID(client *hueapi.Client) {
	hueResp, err := client.Lights.GetAllLights()
	printHueResponse(hueResp, err, "Get all lights for id", false)

	if len(hueResp.Data) == 0 {
		fmt.Println("No lights found.")
		return
	}

	light, err := client.Lights.GetLightByID(hueResp.Data[0].ID)
	printHueResponse(light, err, fmt.Sprintf("Get light by ID: %s", hueResp.Data[0].ID), true)
}

func TestSetLight(client *hueapi.Client) {
	hueResp, err := client.Lights.GetAllLights()
	printHueResponse(hueResp, err, "Get all lights for set", false)

	if len(hueResp.Data) == 0 {
		fmt.Println("No lights found.")
		return
	}

	// lightEnabled := !light.On.On
	// brightness := rand.Intn(100) + 1
	// lightBuilder := builders.NewLightBuilder()
	// lightBuilder.SetOnOff(lightEnabled)
	// lightBuilder.Brightness(float64(brightness))
	// update := lightBuilder.Build()
	//
	// hueResp, err := client.Lights.SetLightState(light.ID, update)

	onResp, err := client.Lights.On(hueResp.Data[0].ID)
	printHueActionResponse(onResp, err, "Turn light on", true)

	identifyResp, err := client.Lights.Identify(hueResp.Data[0].ID, 800)
	printHueActionResponse(identifyResp, err, "Identify on", true)
}
