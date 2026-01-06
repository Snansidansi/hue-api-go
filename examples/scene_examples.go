package main

import (
	"fmt"
	"time"

	"github.com/Snansidansi/hue-api-go"
	"github.com/Snansidansi/hue-api-go/builders"
	"github.com/Snansidansi/hue-api-go/models"
)

func TestScenes(client *hueapi.Client) {
	// TestGetAllScenes(client)
	// TestCreateScene(client, true)
	// TestActivateScene(client)
	// TestDeleteScene(client)
}

func TestGetAllScenes(client *hueapi.Client) {
	hueResp, err := client.Scenes.GetScenes()
	printHueResponse(hueResp, err, "Get all scenes", true)
}

func TestCreateScene(client *hueapi.Client, cleanupScene bool) (zoneID, sceneID string) {
	hueRespLights, err := client.Lights.GetAllLights()
	if err != nil || len(hueRespLights.Data) == 0 {
		fmt.Println("Cannot create scene: need at least 1 light.")
		return
	}

	lightIdentifier := models.ResourceIdentifier{Rid: hueRespLights.Data[0].ID, Rtype: "light"}

	zoneChildren := []models.ResourceIdentifier{lightIdentifier}
	newZone := builders.NewZone("API Test Zone", "other", zoneChildren)

	hueRespZone, err := client.Zones.CreateZone(newZone)
	if err != nil || len(hueRespZone.Data) == 0 {
		printHueActionResponse(hueRespZone, err, "Create Zone", true)
		return
	}

	zoneIdentifier := models.ResourceIdentifier{Rid: hueRespZone.Data[0].Rid, Rtype: hueRespZone.Data[0].Rtype}

	sceneBuilder := builders.NewCreateSceneBuilder("Test Scene", zoneIdentifier)
	sceneBuilder.AddAction(lightIdentifier, func(sab *builders.SceneActionBuilder) {
		sab.SetOnOff(true).Brightness(75)
	})

	sceneToCreate := sceneBuilder.Build()

	printStructFormatted(sceneToCreate)

	hueResp, err := client.Scenes.CreateScene(sceneToCreate)
	printHueActionResponse(hueResp, err, "Create scene", true)

	if cleanupScene {
		removeTestZone(client, zoneIdentifier.Rid)
	}

	return zoneIdentifier.Rid, hueResp.Data[0].Rid
}

func removeTestZone(client *hueapi.Client, zoneID string) {
	fmt.Println("\nPress Enter to delete the test zone")
	fmt.Scanln()

	hueRespDelete, err := client.Zones.DeleteZone(zoneID)
	printHueActionResponse(hueRespDelete, err, "Delete Zone", true)
}

func TestActivateScene(client *hueapi.Client) {
	//Create test scene
	testZoneID, testSceneID := TestCreateScene(client, false)

	recall := builders.NewRecallBuilder().WithAction("active").Build()

	updateResp, err := client.Scenes.ActivateScene(testSceneID, recall)
	printHueActionResponse(updateResp, err, "Activate scene", true)

	removeTestZone(client, testZoneID)
}

func TestDeleteScene(client *hueapi.Client) {
	testZoneId, testSceneID := TestCreateScene(client, false)

	time.Sleep(2 * time.Second)

	deleteResp, err := client.Scenes.DeleteScene(testSceneID)
	printHueActionResponse(deleteResp, err, "Delete scene", true)

	removeTestZone(client, testZoneId)
}
