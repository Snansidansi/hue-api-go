package hueapi

import (
	"fmt"
	"net/http"

	"github.com/Snansidansi/hue-api-go/models"
)

type SceneService struct {
	client *Client
}

func (s *SceneService) GetScenes() (*models.HueResponse[models.Scene], error) {
	urlSuffix := "resource/scene"
	return doGetRequest[models.Scene](s.client, urlSuffix)
}

func (s *SceneService) GetScene(id string) (*models.HueResponse[models.Scene], error) {
	urlSuffix := fmt.Sprintf("resource/scene/%s", id)
	return doGetRequest[models.Scene](s.client, urlSuffix)
}

func (s *SceneService) CreateScene(scene *models.SceneNew) (*models.HueActionResponse, error) {
	urlSuffix := "resource/scene"
	return doActionRequest(s.client, http.MethodPost, urlSuffix, scene)
}

func (s *SceneService) UpdateScene(id string, scene *models.SceneUpdate) (*models.HueActionResponse, error) {
	urlSuffix := fmt.Sprintf("resource/scene/%s", id)
	return doActionRequest(s.client, http.MethodPut, urlSuffix, scene)
}

func (s *SceneService) DeleteScene(id string) (*models.HueActionResponse, error) {
	urlSuffix := fmt.Sprintf("resource/scene/%s", id)
	return doActionRequest(s.client, http.MethodDelete, urlSuffix, nil)
}

func (s *SceneService) ActivateScene(id string, recall models.SceneRecallAction) (*models.HueActionResponse, error) {
	body := models.SceneUpdate{
		Recall: &recall,
	}

	urlSuffix := fmt.Sprintf("resource/scene/%s", id)
	return doActionRequest(s.client, http.MethodPut, urlSuffix, &body)
}

