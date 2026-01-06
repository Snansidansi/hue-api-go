package main

import (
	"github.com/Snansidansi/hue-api-go"
)

func TestBridgeHome(client *hueapi.Client) {
	TestGetBridgeHomes(client)
}

func TestGetBridgeHomes(client *hueapi.Client) {
	hueResp, err := client.BridgeHome.GetBridgeHomes()
	printHueResponse(hueResp, err, "Get bridge homes", true)
}
