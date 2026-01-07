package hueapi

import (
	"fmt"
	"net/http"

	"github.com/Snansidansi/hue-api-go/models"
)

type entertainmentConfigurationService struct {
	client *Client
}

func (s *entertainmentConfigurationService) GetAllEntertainmentConfigurations() (*models.HueResponse[models.EntertainmentConfiguration], error) {
	urlSuffix := "resource/entertainment_configuration"
	return doGetRequest[models.EntertainmentConfiguration](s.client, urlSuffix)
}

func (s *entertainmentConfigurationService) GetEntertainmentConfigurationByID(id string) (*models.HueResponse[models.EntertainmentConfiguration], error) {
	urlSuffix := fmt.Sprintf("resource/entertainment_configuration/%s", id)
	return doGetRequest[models.EntertainmentConfiguration](s.client, urlSuffix)
}

func (s *entertainmentConfigurationService) CreateEntertainmentConfiguration(config models.EntertainmentConfigurationNew) (*models.HueActionResponse, error) {
	urlSuffix := "resource/entertainment_configuration"
	return doActionRequest(s.client, http.MethodPost, urlSuffix, &config)
}

func (s *entertainmentConfigurationService) UpdateEntertainmentConfiguration(id string, config models.EntertainmentConfigurationUpdate) (*models.HueActionResponse, error) {
	urlSuffix := fmt.Sprintf("resource/entertainment_configuration/%s", id)
	return doActionRequest(s.client, http.MethodPut, urlSuffix, &config)
}

func (s *entertainmentConfigurationService) DeleteEntertainmentConfiguration(id string) (*models.HueActionResponse, error) {
	urlSuffix := fmt.Sprintf("resource/entertainment_configuration/%s", id)
	return doActionRequest(s.client, http.MethodDelete, urlSuffix, nil)
}
func (s *entertainmentConfigurationService) StopEntertainmentConfiguration(id string) (*models.HueActionResponse, error) {
	actionValue := "stop"
	update := models.EntertainmentConfigurationUpdate{
		Action: &actionValue,
	}
	return s.UpdateEntertainmentConfiguration(id, update)
}
