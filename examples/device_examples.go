package main

import (
	"fmt"

	"github.com/snansidansi/hueapi"
	"github.com/snansidansi/hueapi/builders"
	"github.com/snansidansi/hueapi/models"
)

func TestDevice(client *hueapi.Client) {
	// TestGetAllDevice(client)
	// TestUpdateDevice(client)
}

func TestGetAllDevice(client *hueapi.Client) {
	hueResp, err := client.Device.GetDevices()
	printHueResponse(hueResp, err, "Get all devices", true)
}

func TestUpdateDevice(client *hueapi.Client) {
	hueResp, err := client.Device.GetDevices()
	if err != nil || len(hueResp.Data) == 0 {
		fmt.Println("No devices found or error retrieving.")
		return
	}

	var targetDevice *models.Device

	for i := range hueResp.Data {
		device := &hueResp.Data[i]

		isLight := false
		for _, service := range device.Services {
			if service.Rtype == "light" {
				isLight = true
				break
			}
		}

		if isLight {
			targetDevice = device
			break
		}
	}

	if targetDevice == nil {
		fmt.Println("No device of type 'light' found.")
		return
	}

	fmt.Printf("Updating Device (Identify): %s (ID: %s)\n", *targetDevice.Metadata.Name, targetDevice.ID)

	update := builders.NewDeviceBuilder().
		Identify("identify").
		Build()

	actionResp, err := client.Device.UpdateDevice(targetDevice.ID, update)
	printHueActionResponse(actionResp, err, "Identify Device (Update)", true)
}
