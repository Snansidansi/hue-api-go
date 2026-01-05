package hueapi

import (
	"fmt"

	"github.com/snansidansi/hueapi/models"
)

type EntertainmentService struct {
	client *Client
}

func (s *EntertainmentService) GetAllEntertainment() (*models.HueResponse[models.Entertainment], error) {
	urlSuffix := "resource/entertainment"
	return doGetRequest[models.Entertainment](s.client, urlSuffix)
}

func (s *EntertainmentService) GetEntertainmentByID(id string) (*models.HueResponse[models.Entertainment], error) {
	urlSuffix := fmt.Sprintf("resource/entertainment/%s", id)
	return doGetRequest[models.Entertainment](s.client, urlSuffix)
}
