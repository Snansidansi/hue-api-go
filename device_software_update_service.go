package hueapi

import (
	"fmt"

	"github.com/Snansidansi/hue-api-go/models"
)

type DeviceSoftwareUpdateService struct {
	client *Client
}

func (s *DeviceSoftwareUpdateService) GetAllDeviceSoftwareUpdates() (*models.HueResponse[models.DeviceSoftwareUpdate], error) {
	urlSuffix := "resource/device_software_update"
	return doGetRequest[models.DeviceSoftwareUpdate](s.client, urlSuffix)
}

func (s *DeviceSoftwareUpdateService) GetDeviceSoftwareUpdateByID(id string) (*models.HueResponse[models.DeviceSoftwareUpdate], error) {
	urlSuffix := fmt.Sprintf("resource/device_software_update/%s", id)
	return doGetRequest[models.DeviceSoftwareUpdate](s.client, urlSuffix)
}
