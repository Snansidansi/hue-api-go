package main

import "github.com/snansidansi/hueapi"

func TestDevicePower(client *hueapi.Client) {
	TestGetAllDevicePower(client)
}

func TestGetAllDevicePower(client *hueapi.Client) {
	hueResp, err := client.DevicePower.GetAllDevicePower()
	printHueResponse(hueResp, err, "Get device power", true)
}
