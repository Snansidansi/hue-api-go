package hueapi

import (
	"fmt"

	"github.com/Snansidansi/hue-api-go/models"
)

type ButtonService struct {
	client *Client
}

func (s *ButtonService) GetAllButtons() (*models.HueResponse[models.Button], error) {
	urlSuffix := "resource/button"
	return doGetRequest[models.Button](s.client, urlSuffix)
}

func (s *ButtonService) GetButtonByID(id string) (*models.HueResponse[models.Button], error) {
	urlSuffix := fmt.Sprintf("resource/button/%s", id)
	return doGetRequest[models.Button](s.client, urlSuffix)
}
