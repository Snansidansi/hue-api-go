# Inoffical Phillips Hue API

## Live API (Eventstream)

The Live API connects to the Hue Bridge's event stream (SSE) to receive real-time updates without polling. It parses incoming JSON data into typed Go structs and intelligently filters between state changes and configuration changes.

### Key Features
* **Real-time Updates:** Receives instant feedback for lights, buttons, groups, and scenes.
* **Smart Change Detection:** Every event includes a `StateChanges` boolean flag.
    * `true`: Pure state change (e.g., On/Off, Brightness, Color). You can update your UI directly.
    * `false`: Structural or configuration change (e.g., Name change, Room edit, Scene modification) or `Add`/`Delete` event. It is recommended to re-fetch the resource via the REST API.
* **Raw Access:** Provides access to the raw JSON byte stream for debugging or custom parsing.

### Available Events
The event models are defined in `models/eventstream.go`. The service logic resides in `eventstream_service.go`.

| Event Type | Description |
| :--- | :--- |
| `LightChangeEvent` | Updates for On/Off, Dimming, Color, Color Temperature, and Dynamics. |
| `GroupChangeEvent` | Updates for Rooms, Zones, and Grouped Lights (State only). |
| `SceneEvent` | Updates for Scene activations (`last_recall`) or structure changes. |
| `ButtonEvent` | Real-time button interactions (initial_press, repeat, short_release, etc.). |
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
