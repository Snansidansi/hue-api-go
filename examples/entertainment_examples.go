package main

import (
	"fmt"

	"github.com/snansidansi/hueapi"
)

func TestEntertainment(client *hueapi.Client) {
	TestGetAllEntertainment(client)
}

func TestGetAllEntertainment(client *hueapi.Client) {
	hueResp, err := client.Entertainment.GetAllEntertainment()

	printHueResponse(hueResp, err, "Get all entertainment capabilities", true)
	if err == nil && len(hueResp.Data) > 0 {
		first := hueResp.Data[0]
		fmt.Printf("Example info for device %s:\n", first.Owner.Rid)
		fmt.Printf("  - Can render: %v\n", first.Renderer)
		fmt.Printf("  - Can be proxy: %v\n", first.Proxy)
	}
}
