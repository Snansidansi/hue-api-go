# Unofficial Phillips Hue API

This is a Go client for the Phillips Hue API v2. It provides a convenient way to interact with your Hue Bridge and control your lights and other devices.

For a complete overview of all features, please have a look at the `examples` directory.

## Getting Started

To use the API, you first need to create a client:

```go
import "github.com/Snansidansi/hue-api-go"
import "github.com/Snansidansi/hue-api-go/models"

// ...

bridge := models.Bridge{
    Id:       "your-bridge-id",
    IPAdress: "your-bridge-ip",
}
apiKey := "your-hue-application-key"

client := hueapi.NewClient(bridge, apiKey, nil, true)
```

Have a look at examples/bridge_examples.go to see how you can discover bridges in your local network and register your client.

---

The `client` object provides access to various services to interact with your Hue devices, such as `Lights`, `Rooms`, `Zones`, and `Scenes`.

For more detailed information about the API, including all available endpoints and HTTP status codes, please refer to the [official Hue API v2 documentation](https://developers.meethue.com/develop/hue-api-v2/api-reference/).

## Live API (Eventstream)

The Live API connects to the Hue Bridge's event stream (SSE) to receive real-time updates without polling. It parses incoming JSON data into typed Go structs and intelligently filters between state changes and configuration changes.

### Key Features

- **Real-time Updates:** Receives instant feedback for lights, buttons, groups, and scenes.
- **Smart Change Detection:** Every event includes a `StateChanges` boolean flag.
  - `true`: Pure state change (e.g., On/Off, Brightness, Color). You can update your UI directly.
  - `false`: Structural or configuration change (e.g., Name change, Room edit, Scene modification) or `Add`/`Delete` event. It is recommended to re-fetch the resource via the REST API.
- **Raw Access:** Provides access to the raw JSON byte stream for debugging or custom parsing.

### Available Events

The event models are defined in `models/eventstream.go`. The service logic resides in `eventstream_service.go`.

| Event Type                        | Description                                                                 |
| :-------------------------------- | :-------------------------------------------------------------------------- |
| `LightChangeEvent`                | Updates for On/Off, Dimming, Color, Color Temperature, and Dynamics.        |
| `GroupChangeEvent`                | Updates for Rooms, Zones, and Grouped Lights (State only).                  |
| `SceneEvent`                      | Updates for Scene activations (`last_recall`) or structure changes.         |
| `ButtonEvent`                     | Real-time button interactions (initial_press, repeat, short_release, etc.). |
| `EntertainmentConfigurationEvent` | Status updates for entertainment areas (active/inactive) and streamer info. |

### Usage

**Structured Event Stream (Recommended):**

```go
es := hueapi.NewEventService(client)
es.Start()

// Receive events from the channel
for event := range es.GetEventStream(100) {
    switch e := event.(type) {
    case models.LightChangeEvent:
        if e.StateChanges {
            // Update UI/State directly
            fmt.Printf("Light %s changed brightness\n", e.ID)
        } else {
            // Config changed or added/deleted -> Reload from REST API
            fmt.Printf("Light %s config changed, reloading...\n", e.ID)
        }
    }
}
```
