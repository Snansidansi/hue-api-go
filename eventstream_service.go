package hueapi

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/snansidansi/hueapi/models"
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

// Get the raw/eventstream before running Start() or the events till you get the channels are lost
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

			// Hue Bridge sends data in this format: "data: [...]"
			if after, ok := strings.CutPrefix(line, "data: "); ok {
				payload := after
				jsonBytes := []byte(payload)

				es.dispatch(jsonBytes)
			}
		}

		resp.Body.Close()

		// Bridge closed the connection or Scanner closed
		if es.ctx.Err() == nil {
			// Add a brief delay to prevent overwhelming the bridge
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
	Type string       `json:"type"` // "update", "add"
	Data []streamData `json:"data"`
}

type streamData struct {
	ID      string          `json:"id"`
	Type    string          `json:"type"` // "light", "grouped_light", "button"
	On      *models.On      `json:"on,omitempty"`
	Dimming *models.Dimming `json:"dimming,omitempty"`
	Color   *models.Color   `json:"color,omitempty"`
	Button  *struct {
		LastEvent string `json:"last_event"`
	} `json:"button,omitempty"`
}

// Can return the following events: LightChangeEvent, GroupChangeEvent, ButtonEvent
func (es *EventService) processStructured(data []byte) {
	var msgs []streamMessage
	if err := json.Unmarshal(data, &msgs); err != nil {
		if es.errorChan != nil {
			es.errorChan <- fmt.Errorf("Error while processing structured data: %w", err)
		}
	}

	for _, msg := range msgs {
		if msg.Type != "update" && msg.Type != "add" {
			continue
		}

		for _, item := range msg.Data {
			var event any

			switch item.Type {
			case "light":
				event = models.LightChangeEvent{
					ID:      item.ID,
					On:      item.On,
					Dimming: item.Dimming,
					Color:   item.Color,
				}
			case "grouped_light", "zone", "room":
				event = models.GroupChangeEvent{
					ID:      item.ID,
					On:      item.On,
					Dimming: item.Dimming,
					Type:    item.Type,
				}
			case "button":
				if item.Button != nil {
					event = models.ButtonEvent{
						ID:        item.ID,
						EventType: item.Button.LastEvent,
					}
				}
			}

			if event != nil {
				select {
				case es.eventsChan <- event:
				default:
				}
			}
		}
	}
}
