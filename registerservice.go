package hueapi

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Snansidansi/hue-api-go/models"
)

type RegisterService struct {
	client *Client
}

type registrationResponse struct {
	Success *struct {
		Username  string `json:"username"`
		ClientKey string `json:"clientkey"`
	} `json:"success"`
	Error *struct {
		Type        int    `json:"type"`
		Description string `json:"description"`
	} `json:"error"`
}

// Uses http.DefaultCLient when httpClient is nil.
// The username is the normal http api key.
// The device type is app_name#instance_name
func RegisterBridge(
	client *http.Client,
	bridge *models.Bridge,
	deviceName string,
	generateClientKey bool,
) (username, clientKey string, err error) {
	if client == nil {
		client = &http.Client{}
	}

	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	client.Transport = customTransport

	// Request
	data := struct {
		DeviceType        string `json:"devicetype"`
		GenerateClientKey bool   `json:"generateclientkey"`
	}{
		DeviceType:        deviceName,
		GenerateClientKey: generateClientKey,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", "", err
	}

	url := fmt.Sprintf("https://%s/api", bridge.IPAdress)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", "", err
	}

	req.Header.Set("Content-Type", "application/json")

	// Response
	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	var respData []registrationResponse
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return "", "", err
	}

	if len(respData) == 0 {
		return "", "", fmt.Errorf("empty response from bridge")
	}

	if respData[0].Error != nil {
		hueError := models.HueError{
			Type:        respData[0].Error.Type,
			Description: respData[0].Error.Description,
		}

		return "", "", &hueError
	}

	if respData[0].Success != nil {
		return respData[0].Success.Username, respData[0].Success.ClientKey, nil
	}

	return "", "", fmt.Errorf("Unexpected error: responsedata fits non of the given cases")
}
