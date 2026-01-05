package hueapi

import (
	"fmt"
	"net/http"

	"github.com/snansidansi/hueapi/models"
)

type DeviceService struct {
	client *Client
}

func (s *DeviceService) GetDevices() (*models.HueResponse[models.Device], error) {
	urlSuffix := "resource/device"
	return doGetRequest[models.Device](s.client, urlSuffix)
}

func (s *DeviceService) GetDeviceByID(id string) (*models.HueResponse[models.Device], error) {
	urlSuffix := fmt.Sprintf("resource/device/%s", id)
	return doGetRequest[models.Device](s.client, urlSuffix)
}

func (s *DeviceService) UpdateDevice(id string, device models.DevicePut) (*models.HueActionResponse, error) {
	urlSuffix := fmt.Sprintf("resource/device/%s", id)
	return doActionRequest(s.client, http.MethodPut, urlSuffix, &device)
}

func (s *DeviceService) DeleteDevice(id string) (*models.HueActionResponse, error) {
	urlSuffix := fmt.Sprintf("resource/device/%s", id)
	return doActionRequest(s.client, http.MethodDelete, urlSuffix, nil)
}
