package hueapi

import (
	"fmt"
	"net/http"

	"github.com/snansidansi/hueapi/models"
)

type GroupedLightService struct {
	client *Client
}

func (s *GroupedLightService) GetAllGroupedLights() (*models.HueResponse[models.GroupedLight], error) {
	urlSuffix := "resource/grouped_light"
	return doGetRequest[models.GroupedLight](s.client, urlSuffix)
}

func (s *GroupedLightService) GetGroupedLightByID(id string) (*models.HueResponse[models.GroupedLight], error) {
	urlSuffix := fmt.Sprintf("resource/grouped_light/%s", id)
	return doGetRequest[models.GroupedLight](s.client, urlSuffix)
}

func (s *GroupedLightService) UpdateGroupedLight(id string, update models.GroupedLightUpdate) (*models.HueActionResponse, error) {
	urlSuffix := fmt.Sprintf("resource/grouped_light/%s", id)
	return doActionRequest(s.client, http.MethodPut, urlSuffix, &update)
}
