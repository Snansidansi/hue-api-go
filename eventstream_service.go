package hueapi

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Snansidansi/hue-api-go/models"
)

type EventService struct {
	client *Client

	ctx    context.Context
	cancel context.CancelFunc

	rawChan    chan []byte
	eventsChan chan any
	errorChan  chan error
}

func NewEventService(client *Client) *EventService {
	ctx, cancel := context.WithCancel(context.Background())

	return &EventService{
		client: client,
		ctx:    ctx,
		cancel: cancel,
	}
}

func (es *EventService) GetRawStream(chanBufSize uint) <-chan []byte {
	if es.rawChan == nil {
		es.rawChan = make(chan []byte, chanBufSize)
	}
	return es.rawChan
}

// possible events are LightChangeEvent, GroupChangeEvent, ButtonEvent, SceneEvent
func (es *EventService) GetEventStream(chanBufSize uint) <-chan any {
	if es.eventsChan == nil {
		es.eventsChan = make(chan any, chanBufSize)
	}
	return es.eventsChan
}

func (es *EventService) GetErrorStream(chanBufSize uint) <-chan error {
	if es.errorChan == nil {
		es.errorChan = make(chan error, chanBufSize)
	}
	return es.errorChan
}

func (es *EventService) Start() {
	go es.listenLoop()
}

func (es *EventService) Stop() {
	es.cancel()
}

func (es *EventService) listenLoop() {
	url := fmt.Sprintf("https://%s/eventstream/clip/v2", es.client.Bridge.IPAdress)

	for {
		if es.ctx.Err() != nil {
			return
		}

		req, err := http.NewRequestWithContext(es.ctx, "GET", url, nil)
		if err != nil && es.errorChan != nil {
			es.errorChan <- fmt.Errorf("Error creating request: %w", err)
		}
		req.Header.Set("Accept", "text/event-stream")

		resp, err := es.client.HTTPClient.Do(req)
		if err != nil {
			if es.ctx.Err() != nil {
				return
			}
			if es.errorChan != nil {
				es.errorChan <- fmt.Errorf("Eventstream connect error. Retrying in 2s. Error: %w", err)
			}
			select {
			case <-time.After(2 * time.Second):
				continue
			case <-es.ctx.Done():
				return
			}
		}

		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Text()
			if after, ok := strings.CutPrefix(line, "data: "); ok {
				es.dispatch([]byte(after))
			}
		}

		resp.Body.Close()
		if es.ctx.Err() == nil {
			time.Sleep(1 * time.Second)
		}
	}
}

func (es *EventService) dispatch(data []byte) {
	if es.rawChan != nil {
		select {
		case es.rawChan <- data:
		default:
		}
	}

	if es.eventsChan != nil {
		es.processStructured(data)
	}
}

type streamMessage struct {
	Type         string       `json:"type"` // "update", "add", "delete"
	CreationTime time.Time    `json:"creationtime"`
	Data         []streamData `json:"data"`
}

type streamData struct {
	ID   string `json:"id"`
	Type string `json:"type"`

	On               *models.On               `json:"on,omitempty"`
	Dimming          *models.Dimming          `json:"dimming,omitempty"`
	Color            *models.Color            `json:"color,omitempty"`
	ColorTemperature *models.ColorTemperature `json:"color_temperature,omitempty"`
	Dynamics         *models.Dynamics         `json:"dynamics,omitempty"`

	Status json.RawMessage `json:"status,omitempty"`

	ActiveStreamer *models.ResourceIdentifier `json:"active_streamer,omitempty"`

	Button *struct {
		LastEvent string `json:"last_event"`
	} `json:"button,omitempty"`

	Metadata *struct {
		Name      string `json:"name"`
		Archetype string `json:"archetype"`
	} `json:"metadata,omitempty"`

	Owner    *models.ResourceIdentifier  `json:"owner,omitempty"`
	Children []models.ResourceIdentifier `json:"children,omitempty"`
	Group    *models.ResourceIdentifier  `json:"group,omitempty"`
	Actions  []models.SceneAction        `json:"actions,omitempty"`
	Speed    *float64                    `json:"speed,omitempty"`
}

func (es *EventService) processStructured(data []byte) {
	var msgs []streamMessage
	if err := json.Unmarshal(data, &msgs); err != nil {
		if es.errorChan != nil {
			es.errorChan <- fmt.Errorf("Error while processing structured data: %w", err)
		}
		return
	}

	for _, msg := range msgs {
		eventType := msg.Type
		timestamp := msg.CreationTime

		for _, item := range msg.Data {
			stateChanges := es.determineStateChange(eventType, item)

			base := models.BaseEventFields{
				EventType:    eventType,
				ID:           item.ID,
				Timestamp:    timestamp,
				StateChanges: stateChanges,
			}

			event := es.createEventModel(base, item)

			if event != nil {
				select {
				case es.eventsChan <- event:
				default:
				}
			}
		}
	}
}

func (es *EventService) determineStateChange(eventType string, item streamData) bool {
	hasState := false
	hasConfig := false

	switch item.Type {
	case "light":
		if item.On != nil || item.Dimming != nil || item.Color != nil || item.ColorTemperature != nil || item.Dynamics != nil {
			hasState = true
		}

		if item.Metadata != nil {
			hasConfig = true
		}

	case "grouped_light":
		if item.On != nil || item.Dimming != nil {
			hasState = true
		}
		// The Hue API sends 'owner' with state updates (context), so it's not a config change.
		if item.Metadata != nil {
			hasConfig = true
		}

	case "zone", "room":
		if item.On != nil || item.Dimming != nil {
			hasState = true
		}

		// If children change or name changes
		if item.Metadata != nil || len(item.Children) > 0 {
			hasConfig = true
		}

	case "button":
		if item.Button != nil {
			hasState = true
		}

	case "scene":
		if len(item.Status) > 0 && string(item.Status) != "null" {
			hasState = true
		}

		if item.Metadata != nil || item.Group != nil || len(item.Actions) > 0 || item.Speed != nil {
			hasConfig = true
		}

	case "entertainment_configuration":
		if len(item.Status) > 0 && string(item.Status) != "null" {
			hasState = true
		}

		if item.ActiveStreamer != nil {
			hasState = true
		}
		if item.Metadata != nil {
			hasConfig = true
		}
	}

	if eventType == models.EventTypeUpdate && hasState && !hasConfig {
		return true
	}

	return false
}

func (es *EventService) createEventModel(base models.BaseEventFields, item streamData) any {
	switch item.Type {
	case "light":
		return models.LightChangeEvent{
			BaseEventFields:  base,
			On:               item.On,
			Dimming:          item.Dimming,
			Color:            item.Color,
			ColorTemperature: item.ColorTemperature,
			Dynamics:         item.Dynamics,
		}

	case "grouped_light", "zone", "room":
		return models.GroupChangeEvent{
			BaseEventFields: base,
			Type:            item.Type,
			On:              item.On,
			Dimming:         item.Dimming,
		}

	case "button":
		btnAction := ""
		if item.Button != nil {
			btnAction = item.Button.LastEvent
		}
		return models.ButtonEvent{
			BaseEventFields: base,
			Button:          btnAction,
		}

	case "scene":
		var sceneStatus *models.SceneStatusEvent
		if len(item.Status) > 0 && string(item.Status) != "null" {
			err := json.Unmarshal(item.Status, &sceneStatus)
			if err != nil && es.errorChan != nil {
				es.errorChan <- err
			}
		}
		return models.SceneEvent{
			BaseEventFields: base,
			Status:          sceneStatus,
		}

	case "entertainment_configuration":
		var statusString string
		if len(item.Status) > 0 && string(item.Status) != "null" {
			err := json.Unmarshal(item.Status, &statusString)
			if err != nil && es.errorChan != nil {
				es.errorChan <- err
			}
		}
		return models.EntertainmentConfigurationEvent{
			BaseEventFields: base,
			Status:          statusString,
			ActiveStreamer:  item.ActiveStreamer,
		}
	}

	return nil
}
