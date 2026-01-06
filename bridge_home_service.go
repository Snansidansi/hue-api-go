package hueapi

import (
	"fmt"

	"github.com/Snansidansi/hue-api-go/models"
)

type BridgeHomeService struct {
	client *Client
}

func (s *BridgeHomeService) GetBridgeHomes() (*models.HueResponse[models.BridgeHome], error) {
	urlSuffix := "resource/bridge_home"
	return doGetRequest[models.BridgeHome](s.client, urlSuffix)
}

func (s *BridgeHomeService) GetBridgeHome(id string) (*models.HueResponse[models.BridgeHome], error) {
	urlSuffix := fmt.Sprintf("resource/bridge_home/%s", id)
	return doGetRequest[models.BridgeHome](s.client, urlSuffix)
}
