package hueapi

import (
	"fmt"

	"github.com/snansidansi/hueapi/models"
)

type DevicePowerService struct {
	client *Client
}

func (s *DevicePowerService) GetAllDevicePower() (*models.HueResponse[models.DevicePower], error) {
	urlSuffix := "resource/device_power"
	return doGetRequest[models.DevicePower](s.client, urlSuffix)
}

func (s *DevicePowerService) GetDevicePowerByID(id string) (*models.HueResponse[models.DevicePower], error) {
	urlSuffix := fmt.Sprintf("resource/device_power/%s", id)
	return doGetRequest[models.DevicePower](s.client, urlSuffix)
}
