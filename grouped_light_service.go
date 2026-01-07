package hueapi

import (
	"fmt"
	"net/http"

	"github.com/Snansidansi/hue-api-go/models"
)

type groupedLightService struct {
	client *Client
}

func (s *groupedLightService) GetAllGroupedLights() (*models.HueResponse[models.GroupedLight], error) {
	urlSuffix := "resource/grouped_light"
	return doGetRequest[models.GroupedLight](s.client, urlSuffix)
}

func (s *groupedLightService) GetGroupedLightByID(id string) (*models.HueResponse[models.GroupedLight], error) {
	urlSuffix := fmt.Sprintf("resource/grouped_light/%s", id)
	return doGetRequest[models.GroupedLight](s.client, urlSuffix)
}

func (s *groupedLightService) UpdateGroupedLight(id string, update models.GroupedLightUpdate) (*models.HueActionResponse, error) {
	urlSuffix := fmt.Sprintf("resource/grouped_light/%s", id)
	return doActionRequest(s.client, http.MethodPut, urlSuffix, &update)
}
