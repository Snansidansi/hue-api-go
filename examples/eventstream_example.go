package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/snansidansi/hueapi"
	"github.com/snansidansi/hueapi/models"
	"github.com/snansidansi/hueapi/util"
)

func TestEventStream(client *hueapi.Client) {
	TestRawEvents(client)
	// TestStructuredEvents(client)
}

func TestRawEvents(client *hueapi.Client) {
	fmt.Println("--> Start Raw Event Listener")
	fmt.Println("Press [ENTER] to stop the listener...")

	es := hueapi.NewEventService(client)

	// Get Channels (BEFORE calling Start)
	rawStream := es.GetRawStream(100)
	errorStream := es.GetErrorStream(10)

	es.Start()
	defer es.Stop()

	stopSignal := make(chan struct{})
	go func() {
		fmt.Scanln()
		close(stopSignal)
	}()

	for {
		select {
		case <-stopSignal:
			fmt.Println("Stopped by user.")
			return

		case data := <-rawStream:
			fmt.Printf("RAW DATA (%d bytes):", len(data))

			var out bytes.Buffer
			err := json.Indent(&out, data, "", "  ")

			if err != nil {
				fmt.Println("Error while formatting response: ", err)
				return
			}

			fmt.Println(out.String())

		case err := <-errorStream:
			fmt.Printf("STREAM ERROR: %v\n", err)
		}
	}
}

func TestStructuredEvents(client *hueapi.Client) {
	fmt.Println("--> Start Structured Event Listener")
	fmt.Println("Press [ENTER] to stop the listener...")

	es := hueapi.NewEventService(client)

	// Get Channels (BEFORE calling Start)
	events := es.GetEventStream(100)
	errors := es.GetErrorStream(10)

	es.Start()
	defer es.Stop()

	stopSignal := make(chan struct{})
	go func() {
		fmt.Scanln()
		close(stopSignal)
	}()

	for {
		select {
		case <-stopSignal:
			fmt.Println("Stopped by user.")
			return

		case rawEvent := <-events:
			switch e := rawEvent.(type) {

			case models.LightChangeEvent:
				fmt.Printf("LIGHT EVENT: ID=%s\n", e.ID)
				if e.On != nil {
					fmt.Printf("   -> On: %v\n", *e.On.On)
				}
				if e.Dimming != nil {
					fmt.Printf("   -> Brightness: %.2f%%\n", *e.Dimming.Brightness)
				}
				if e.Color != nil && e.Color.XY != nil {
					r, g, b := util.XYToRGB(e.Color.XY.X, e.Color.XY.Y, 100)
					fmt.Printf("   -> Color RGB: %v, %v, %v\n", r, g, b)
				}

			case models.GroupChangeEvent:
				fmt.Printf("GROUP EVENT: ID=%s\n", e.ID)
				if e.On != nil {
					fmt.Printf("   -> On: %v\n", *e.On.On)
					fmt.Printf("   -> Type: %v\n", e.Type)
				}

			case models.ButtonEvent:
				fmt.Printf("BUTTON EVENT: ID=%s | Action: %s\n", e.ID, e.EventType)

			default:
				fmt.Printf(" UNKNOWN EVENT TYPE: %T\n", e)
			}

		case err := <-errors:
			fmt.Printf("STREAM ERROR: %v\n", err)
		}
	}
}
