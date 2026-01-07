package hueapi

import (
	"fmt"

	"github.com/Snansidansi/hue-api-go/models"
)

type entertainmentService struct {
	client *Client
}

func (s *entertainmentService) GetAllEntertainment() (*models.HueResponse[models.Entertainment], error) {
	urlSuffix := "resource/entertainment"
	return doGetRequest[models.Entertainment](s.client, urlSuffix)
}

func (s *entertainmentService) GetEntertainmentByID(id string) (*models.HueResponse[models.Entertainment], error) {
	urlSuffix := fmt.Sprintf("resource/entertainment/%s", id)
	return doGetRequest[models.Entertainment](s.client, urlSuffix)
}
