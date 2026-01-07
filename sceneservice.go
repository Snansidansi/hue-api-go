package hueapi

import (
	"fmt"
	"net/http"

	"github.com/Snansidansi/hue-api-go/models"
)

type sceneService struct {
	client *Client
}

func (s *sceneService) GetScenes() (*models.HueResponse[models.Scene], error) {
	urlSuffix := "resource/scene"
	return doGetRequest[models.Scene](s.client, urlSuffix)
}

func (s *sceneService) GetScene(id string) (*models.HueResponse[models.Scene], error) {
	urlSuffix := fmt.Sprintf("resource/scene/%s", id)
	return doGetRequest[models.Scene](s.client, urlSuffix)
}

func (s *sceneService) CreateScene(scene *models.SceneNew) (*models.HueActionResponse, error) {
	urlSuffix := "resource/scene"
	return doActionRequest(s.client, http.MethodPost, urlSuffix, scene)
}

func (s *sceneService) UpdateScene(id string, scene *models.SceneUpdate) (*models.HueActionResponse, error) {
	urlSuffix := fmt.Sprintf("resource/scene/%s", id)
	return doActionRequest(s.client, http.MethodPut, urlSuffix, scene)
}

func (s *sceneService) DeleteScene(id string) (*models.HueActionResponse, error) {
	urlSuffix := fmt.Sprintf("resource/scene/%s", id)
	return doActionRequest(s.client, http.MethodDelete, urlSuffix, nil)
}

func (s *sceneService) ActivateScene(id string, recall models.SceneRecallAction) (*models.HueActionResponse, error) {
	body := models.SceneUpdate{
		Recall: &recall,
	}

	urlSuffix := fmt.Sprintf("resource/scene/%s", id)
	return doActionRequest(s.client, http.MethodPut, urlSuffix, &body)
}
