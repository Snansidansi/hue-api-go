package hueapi

import (
	"fmt"
	"net/http"

	"github.com/Snansidansi/hue-api-go/models"
)

type zoneService struct {
	client *Client
}

func (s *zoneService) GetZones() (*models.HueResponse[models.Zone], error) {
	urlSuffix := "resource/zone"
	return doGetRequest[models.Zone](s.client, urlSuffix)
}

func (s *zoneService) GetZoneByID(id string) (*models.HueResponse[models.Zone], error) {
	urlSuffix := fmt.Sprintf("resource/zone/%s", id)
	return doGetRequest[models.Zone](s.client, urlSuffix)
}

func (s *zoneService) UpdateZone(id string, zone *models.ZoneEdit) (*models.HueActionResponse, error) {
	urlSuffix := fmt.Sprintf("resource/zone/%s", id)
	return doActionRequest(s.client, http.MethodPut, urlSuffix, zone)
}

func (s *zoneService) CreateZone(zone *models.ZoneEdit) (*models.HueActionResponse, error) {
	urlSuffix := "resource/zone"
	return doActionRequest(s.client, http.MethodPost, urlSuffix, zone)
}

func (s *zoneService) DeleteZone(id string) (*models.HueActionResponse, error) {
	urlSuffix := fmt.Sprintf("resource/zone/%s", id)
	return doActionRequest(s.client, http.MethodDelete, urlSuffix, nil)
}
