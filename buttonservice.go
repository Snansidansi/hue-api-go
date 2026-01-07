package hueapi

import (
	"fmt"

	"github.com/Snansidansi/hue-api-go/models"
)

type buttonService struct {
	client *Client
}

func (s *buttonService) GetAllButtons() (*models.HueResponse[models.Button], error) {
	urlSuffix := "resource/button"
	return doGetRequest[models.Button](s.client, urlSuffix)
}

func (s *buttonService) GetButtonByID(id string) (*models.HueResponse[models.Button], error) {
	urlSuffix := fmt.Sprintf("resource/button/%s", id)
	return doGetRequest[models.Button](s.client, urlSuffix)
}
