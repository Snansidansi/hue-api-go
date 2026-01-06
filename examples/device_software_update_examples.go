package main

import (
	"github.com/Snansidansi/hue-api-go"
)

func TestDeviceSoftwareUpdate(client *hueapi.Client) {
	TestGetAllDeviceSoftwareUpdates(client)
}

func TestGetAllDeviceSoftwareUpdates(client *hueapi.Client) {
	hueResp, err := client.DeviceSoftwareUpdate.GetAllDeviceSoftwareUpdates()
	printHueResponse(hueResp, err, "Get all device software updates", true)
}
