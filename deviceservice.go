package hueapi

import (
	"fmt"
	"net/http"

	"github.com/Snansidansi/hue-api-go/models"
)

type deviceService struct {
	client *Client
}

func (s *deviceService) GetDevices() (*models.HueResponse[models.Device], error) {
	urlSuffix := "resource/device"
	return doGetRequest[models.Device](s.client, urlSuffix)
}

func (s *deviceService) GetDeviceByID(id string) (*models.HueResponse[models.Device], error) {
	urlSuffix := fmt.Sprintf("resource/device/%s", id)
	return doGetRequest[models.Device](s.client, urlSuffix)
}

func (s *deviceService) UpdateDevice(id string, device models.DevicePut) (*models.HueActionResponse, error) {
	urlSuffix := fmt.Sprintf("resource/device/%s", id)
	return doActionRequest(s.client, http.MethodPut, urlSuffix, &device)
}

func (s *deviceService) DeleteDevice(id string) (*models.HueActionResponse, error) {
	urlSuffix := fmt.Sprintf("resource/device/%s", id)
	return doActionRequest(s.client, http.MethodDelete, urlSuffix, nil)
}
