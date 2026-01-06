package hueapi

import (
	"fmt"
	"net/http"

	"github.com/Snansidansi/hue-api-go/models"
)

type ZoneService struct {
	client *Client
}

func (s *ZoneService) GetZones() (*models.HueResponse[models.Zone], error) {
	urlSuffix := "resource/zone"
	return doGetRequest[models.Zone](s.client, urlSuffix)
}

func (s *ZoneService) GetZone(id string) (*models.HueResponse[models.Zone], error) {
	urlSuffix := fmt.Sprintf("resource/zone/%s", id)
	return doGetRequest[models.Zone](s.client, urlSuffix)
}

func (s *ZoneService) UpdateZone(id string, zone *models.ZoneEdit) (*models.HueActionResponse, error) {
	urlSuffix := fmt.Sprintf("resource/zone/%s", id)
	return doActionRequest(s.client, http.MethodPut, urlSuffix, zone)
}

func (s *ZoneService) CreateZone(zone *models.ZoneEdit) (*models.HueActionResponse, error) {
	urlSuffix := "resource/zone"
	return doActionRequest(s.client, http.MethodPost, urlSuffix, zone)
}

func (s *ZoneService) DeleteZone(id string) (*models.HueActionResponse, error) {
	urlSuffix := fmt.Sprintf("resource/zone/%s", id)
	return doActionRequest(s.client, http.MethodDelete, urlSuffix, nil)
}
