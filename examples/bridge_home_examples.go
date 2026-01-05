package main

import (
	"github.com/snansidansi/hueapi"
)

func TestBridgeHome(client *hueapi.Client) {
	TestGetBridgeHomes(client)
}

func TestGetBridgeHomes(client *hueapi.Client) {
	hueResp, err := client.BridgeHome.GetBridgeHomes()
	printHueResponse(hueResp, err, "Get bridge homes", true)
}
