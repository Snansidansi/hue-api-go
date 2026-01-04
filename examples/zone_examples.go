package main

import (
	"fmt"

	"github.com/snansidansi/hueapi"
	"github.com/snansidansi/hueapi/builders"
	"github.com/snansidansi/hueapi/models"
)

func TestZones(client *hueapi.Client) {
	// TestGetAllZones(client)
	// TestCreateZone(client)
	// TestUpdateZone(client)
	// TestDeleteZone(client)
}

func TestGetAllZones(client *hueapi.Client) {
	hueResp, err := client.Zones.GetZones()
	printHueResponse(hueResp, err, "Get all zones", true)
}

func TestCreateZone(client *hueapi.Client) {
	hueRespLights, err := client.Lights.GetAllLights()
	printHueResponse(hueRespLights, err, "Get all lights", false)

	if len(hueRespLights.Data) == 0 {
		fmt.Println("Cannot create zone: no lights found.")
		return
	}

	children := []models.ResourceIdentifier{
		{
			Rid:   hueRespLights.Data[0].ID,
			Rtype: hueRespLights.Data[0].Type,
		},
	}

	newZone := builders.NewZone("Test Zone", "living_room", children)

	hueResp, err := client.Zones.CreateZone(newZone)
	printHueActionResponse(hueResp, err, "Create zone", true)
}

func TestUpdateZone(client *hueapi.Client) {
	hueResp, err := client.Zones.GetZones()
	printHueResponse(hueResp, err, "Get all zones", false)

	if err != nil {
		return
	}

	var testZoneID string
	for _, zone := range hueResp.Data {
		if zone.Metadata.Name == "Test Zone" {
			testZoneID = zone.ID
			break
		}
	}

	if testZoneID == "" {
		fmt.Println("No 'Test Zone' found to update.")
		return
	}

	updateBuilder := builders.NewUpdateZoneBuilder()
	updateBuilder.WithName("Updated test zone")
	zoneUpdate := updateBuilder.Build()

	updateResp, err := client.Zones.UpdateZone(testZoneID, &zoneUpdate)
	printHueActionResponse(updateResp, err, "Update zone", true)
}

func TestDeleteZone(client *hueapi.Client) {
	hueResp, err := client.Zones.GetZones()
	printHueResponse(hueResp, err, "Get all zones", false)

	if err != nil {
		return
	}

	for _, zone := range hueResp.Data {
		if zone.Metadata.Name == "Test Zone" || zone.Metadata.Name == "Updated test zone" {
			deleteResp, err := client.Zones.DeleteZone(zone.ID)
			printHueActionResponse(deleteResp, err, fmt.Sprintf("Delete zone '%s'", zone.Metadata.Name), true)
		}
	}
}
