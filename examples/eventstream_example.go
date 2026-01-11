package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/Snansidansi/hue-api-go"
	"github.com/Snansidansi/hue-api-go/models"
	"github.com/Snansidansi/hue-api-go/util"
)

func TestEventStream(client *hueapi.Client) {
	// TestRawEvents(client)
	// TestStructuredEvents(client)
}

func TestRawEvents(client *hueapi.Client) {
	fmt.Println("--> Start Raw Event Listener")
	fmt.Println("Press [ENTER] to stop the listener...")

	rawStream := client.EventStream.GetRawStream(100)
	errorStream := client.EventStream.GetErrorStream(10)

	client.EventStream.Start()
	defer client.EventStream.Stop()

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
				fmt.Println(out.String())
			} else {
				fmt.Println(out.String())
			}
		case err := <-errorStream:
			fmt.Printf("STREAM ERROR: %v\n", err)
		}
	}
}

func printBaseInfo(name string, base models.BaseEventFields) {
	timeStr := base.Timestamp.Format("15:04:05")

	if base.StateChanges {
		fmt.Printf("[%s] %s %s: ID=%s EventID=%s [STATE CHANGE]\n", timeStr, base.EventType, name, base.ID, base.EventID)
	} else {
		// Either Add/Delete OR Config Change -> Reload suggested
		fmt.Printf("[%s] %s %s: ID=%s EventID=%s [RELOAD RESOURCE]\n", timeStr, base.EventType, name, base.ID, base.EventID)
	}
}

func TestStructuredEvents(client *hueapi.Client) {
	fmt.Println("--> Start Structured Event Listener")
	fmt.Println("Press [ENTER] to stop the listener...")

	events := client.EventStream.GetEventStream(100)
	errors := client.EventStream.GetErrorStream(10)

	client.EventStream.Start()
	defer client.EventStream.Stop()

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

			case *models.LightChangeEvent:
				printBaseInfo("LIGHT", e.BaseEventFields)
				if e.StateChanges {
					if e.On != nil {
						fmt.Printf("   -> On: %v\n", *e.On.On)
					}
					if e.Dimming != nil {
						fmt.Printf("   -> Brightness: %.2f%%\n", *e.Dimming.Brightness)
					}
					if e.ColorTemperature != nil && e.ColorTemperature.Mirek != nil {
						fmt.Printf("   -> Mirek: %d\n", *e.ColorTemperature.Mirek)
					}
					if e.Color != nil && e.Color.XY != nil {
						r, g, b := util.XYToRGB(e.Color.XY.X, e.Color.XY.Y, 100)
						fmt.Printf("   -> RGB: %v, %v, %v\n", r, g, b)
					}
				}

			case *models.GroupChangeEvent:
				printBaseInfo("GROUP ("+e.Type+")", e.BaseEventFields)
				if e.StateChanges {
					if e.On != nil {
						fmt.Printf("   -> On: %v\n", *e.On.On)
					}
					if e.Dimming != nil {
						fmt.Printf("   -> Brightness: %.2f%%\n", *e.Dimming.Brightness)
					}
				}

			case *models.ButtonEvent:
				printBaseInfo("BUTTON", e.BaseEventFields)
				if e.Button != "" {
					fmt.Printf("   -> Action: %s\n", e.Button)
				}

			case *models.SceneEvent:
				printBaseInfo("SCENE", e.BaseEventFields)
				if e.StateChanges && e.Status != nil {
					if e.Status.Active != "" {
						fmt.Printf("   -> Active: %s\n", e.Status.Active)
					}
					if e.Status.LastRecall != nil {
						fmt.Printf("   -> Last Recall: %s\n", e.Status.LastRecall.Format("15:04:05"))
					}
				}

			case *models.EntertainmentConfigurationEvent:
				printBaseInfo("ENTERTAINMENT", e.BaseEventFields)
				if e.StateChanges {
					if e.Status != "" {
						fmt.Printf("   -> Status: %s\n", e.Status)
					}
					if e.ActiveStreamer != nil {
						fmt.Printf("   -> Streamer: %s\n", e.ActiveStreamer.Rid)
					}
				}

			default:
				fmt.Printf(" UNKNOWN EVENT TYPE: %T\n", e)
			}

		case err := <-errors:
			fmt.Printf("STREAM ERROR: %v\n", err)
		}
	}
}
