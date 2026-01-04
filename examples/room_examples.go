package main

import (
	"fmt"

	"github.com/snansidansi/hueapi"
	"github.com/snansidansi/hueapi/builders"
	"github.com/snansidansi/hueapi/models"
)

func TestRooms(client *hueapi.Client) {
	// TestGetAllRooms(client)
	// TestCreateRoom(client)
	// TestUpdateRoom(client)
	// TestDeleteRoom(client)
}

func TestGetAllRooms(client *hueapi.Client) {
	hueResp, err := client.Rooms.GetAllRooms()
	printHueResponse(hueResp, err, "Get all rooms", true)
}

func TestCreateRoom(client *hueapi.Client) {
	hueRespLights, err := client.Lights.GetAllLights()
	printHueResponse(hueRespLights, err, "Get all lights", false)

	if len(hueRespLights.Data) == 0 {
		fmt.Println("Cannot create room: no lights found.")
		return
	}

	children := []models.ResourceIdentifier{
		{
			Rid:   hueRespLights.Data[0].Owner.RID,
			Rtype: hueRespLights.Data[0].Owner.RType,
		},
	}

	roomBuilder := builders.NewRoomBuilder()
	roomBuilder.WithArchetype("living_room")
	roomBuilder.WithName("Test Room")
	roomBuilder.WithChildren(children)

	room := roomBuilder.Build()

	hueResp, err := client.Rooms.CreateRoom(room)
	printHueActionResponse(hueResp, err, "Create room", true)
}

func TestUpdateRoom(client *hueapi.Client) {
	hueResp, err := client.Rooms.GetAllRooms()
	printHueResponse(hueResp, err, "Get all rooms", false)

	if err != nil {
		return
	}

	var testRoomID string
	for _, room := range hueResp.Data {
		if room.Metadata.Name == "Test Room" {
			testRoomID = room.ID
			break
		}
	}

	if testRoomID == "" {
		fmt.Println("No 'Test Room' found to update.")
		return
	}

	updateBuilder := builders.NewRoomBuilder()
	updateBuilder.WithName("Updated test room")
	roomUpdate := updateBuilder.Build()

	updateResp, err := client.Rooms.UpdateRoom(testRoomID, roomUpdate)
	printHueActionResponse(updateResp, err, "Update room", true)
}

func TestDeleteRoom(client *hueapi.Client) {
	hueResp, err := client.Rooms.GetAllRooms()
	printHueResponse(hueResp, err, "Get all rooms", false)

	if err != nil {
		return
	}

	for _, room := range hueResp.Data {
		if room.Metadata.Name == "Test Room" || room.Metadata.Name == "Updated test room" {
			deleteResp, err := client.Rooms.DeleteRoom(room.ID)
			printHueActionResponse(deleteResp, err, fmt.Sprintf("Delete room '%s'", room.Metadata.Name), true)
		}
	}
}
