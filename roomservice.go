package hueapi

import (
	"fmt"
	"net/http"

	"github.com/Snansidansi/hue-api-go/models"
)

type roomService struct {
	client *Client
}

func (s *roomService) GetAllRooms() (*models.HueResponse[models.Room], error) {
	urlSuffix := "resource/room"
	return doGetRequest[models.Room](s.client, urlSuffix)
}

func (s *roomService) GetRoomByID(id string) (*models.HueResponse[models.Room], error) {
	urlSuffix := fmt.Sprintf("resource/room/%s", id)
	return doGetRequest[models.Room](s.client, urlSuffix)
}

func (s *roomService) UpdateRoom(id string, roomUpdate *models.RoomEdit) (*models.HueActionResponse, error) {
	urlSuffix := fmt.Sprintf("resource/room/%s", id)
	return doActionRequest(s.client, http.MethodPut, urlSuffix, roomUpdate)
}

func (s *roomService) CreateRoom(room *models.RoomEdit) (*models.HueActionResponse, error) {
	urlSuffix := "resource/room"
	return doActionRequest(s.client, http.MethodPost, urlSuffix, room)
}

func (s *roomService) DeleteRoom(id string) (*models.HueActionResponse, error) {
	urlSuffix := fmt.Sprintf("resource/room/%s", id)
	return doActionRequest(s.client, http.MethodDelete, urlSuffix, nil)
}
