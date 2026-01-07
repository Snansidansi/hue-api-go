package hueapi

import (
	"fmt"

	"github.com/Snansidansi/hue-api-go/models"
)

type devicePowerService struct {
	client *Client
}

func (s *devicePowerService) GetAllDevicePower() (*models.HueResponse[models.DevicePower], error) {
	urlSuffix := "resource/device_power"
	return doGetRequest[models.DevicePower](s.client, urlSuffix)
}

func (s *devicePowerService) GetDevicePowerByID(id string) (*models.HueResponse[models.DevicePower], error) {
	urlSuffix := fmt.Sprintf("resource/device_power/%s", id)
	return doGetRequest[models.DevicePower](s.client, urlSuffix)
}
