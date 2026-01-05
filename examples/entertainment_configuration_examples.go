package main

import (
	"fmt"

	"github.com/snansidansi/hueapi"
)

func TestEntertainmentConfiguration(client *hueapi.Client) {
	// TestGetAllEntertainmentConfigurations(client)
	// TestGetEntertainmentConfigurationByID(client)
	// TestStopEntertainmentConfiguration(client)
}

func TestGetAllEntertainmentConfigurations(client *hueapi.Client) {
	hueResp, err := client.EntertainmentConfiguration.GetAllEntertainmentConfigurations()
	printHueResponse(hueResp, err, "Get all entertainment configurations", true)
}

func TestGetEntertainmentConfigurationByID(client *hueapi.Client) {
	hueResp, err := client.EntertainmentConfiguration.GetAllEntertainmentConfigurations()
	if err != nil || len(hueResp.Data) == 0 {
		fmt.Println("No entertainment configurations found, can't test GetEntertainmentConfigurationByID")
		return
	}

	id := hueResp.Data[0].ID
	hueResp, err = client.EntertainmentConfiguration.GetEntertainmentConfigurationByID(id)
	printHueResponse(hueResp, err, fmt.Sprintf("Get entertainment configuration by ID: %s", id), true)
}

func TestStopEntertainmentConfiguration(client *hueapi.Client) {
	hueResp, err := client.EntertainmentConfiguration.GetAllEntertainmentConfigurations()
	if err != nil || len(hueResp.Data) == 0 {
		fmt.Println("No entertainment configurations found, can't test Start/Stop")
		return
	}
	id := hueResp.Data[0].ID

	// Stop
	actionResp, err := client.EntertainmentConfiguration.StopEntertainmentConfiguration(id)
	printHueActionResponse(actionResp, err, fmt.Sprintf("Stopping entertainment configuration %s", id), true)
}
