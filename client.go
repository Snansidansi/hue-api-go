package hueapi

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snansidansi/hueapi/models"
)

type headerTransport struct {
	base    http.RoundTripper
	apiKey  string
	logging bool
}

func (t *headerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	newReq := req.Clone(req.Context())
	newReq.Header.Add("hue-application-key", t.apiKey)

	if t.logging {
		fmt.Printf("\n[HUE] %s %s", newReq.Method, newReq.URL.String())
	}

	return t.base.RoundTrip(newReq)
}

type Client struct {
	Bridge     models.Bridge
	HTTPClient *http.Client

	Lights                     *LightService
	Rooms                      *RoomService
	Zones                      *ZoneService
	Scenes                     *SceneService
	Register                   *RegisterService
	DevicePower                *DevicePowerService
	DeviceSoftwareUpdate       *DeviceSoftwareUpdateService
	EntertainmentConfiguration *EntertainmentConfigurationService
	Entertainment              *EntertainmentService
	BridgeHome                 *BridgeHomeService
	Device                     *DeviceService
}

// Uses http.DefaultCLient when httpClient is nil.
func NewClient(bridge models.Bridge, apiKey string, httpClient *http.Client, logging bool) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	httpClient.Transport = &headerTransport{
		base:    customTransport,
		apiKey:  apiKey,
		logging: logging,
	}

	c := &Client{
		Bridge:     bridge,
		HTTPClient: httpClient,
	}

	c.Lights = &LightService{client: c}
	c.Rooms = &RoomService{client: c}
	c.Zones = &ZoneService{client: c}
	c.Scenes = &SceneService{client: c}
	c.Register = &RegisterService{client: c}
	c.DevicePower = &DevicePowerService{client: c}
	c.DeviceSoftwareUpdate = &DeviceSoftwareUpdateService{client: c}
	c.EntertainmentConfiguration = &EntertainmentConfigurationService{client: c}
	c.Entertainment = &EntertainmentService{client: c}
	c.BridgeHome = &BridgeHomeService{client: c}
	c.Device = &DeviceService{client: c}

	return c
}

// Uses http.DefaultCLient when httpClient is nil.
func DiscoverBridges(httpClient *http.Client) ([]models.Bridge, error) {
	const DiscoveryURL = "https://discovery.meethue.com/"

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	resp, err := httpClient.Get(DiscoveryURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var foundBridges []models.Bridge
	if err := json.NewDecoder(resp.Body).Decode(&foundBridges); err != nil {
		return nil, err
	}

	return foundBridges, nil
}

func (c *Client) CreateURL(urlSuffix string) string {
	return fmt.Sprintf("https://%s/clip/v2/%s", c.Bridge.IPAdress, urlSuffix)
}

func doGetRequest[T any](client *Client, urlSuffix string) (*models.HueResponse[T], error) {
	url := client.CreateURL(urlSuffix)

	resp, err := client.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var hueResp models.HueResponse[T]
	hueResp.StatusCode = resp.StatusCode

	if err := json.NewDecoder(resp.Body).Decode(&hueResp); err != nil {
		return &hueResp, fmt.Errorf("decoding failed (status %d): %w", resp.StatusCode, err)
	}

	return &hueResp, nil
}

func doActionRequest(client *Client, method, urlSuffix string, body any) (*models.HueActionResponse, error) {
	url := client.CreateURL(urlSuffix)

	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var hueResp models.HueActionResponse
	hueResp.StatusCode = resp.StatusCode

	if err := json.NewDecoder(resp.Body).Decode(&hueResp); err != nil {
		return &hueResp, fmt.Errorf("decoding failed (status %d): %w", resp.StatusCode, err)
	}

	return &hueResp, nil
}
