package hueapi

import (
	"fmt"
	"net/http"

	"github.com/snansidansi/hueapi/models"
)

type RoomService struct {
	client *Client
}

func (s *RoomService) GetAllRooms() (*models.HueResponse[models.Room], error) {
	urlSuffix := "resource/room"
	return doGetRequest[models.Room](s.client, urlSuffix)
}

func (s *RoomService) GetRoomByID(id string) (*models.HueResponse[models.Room], error) {
	urlSuffix := fmt.Sprintf("resource/room/%s", id)
	return doGetRequest[models.Room](s.client, urlSuffix)
}

func (s *RoomService) UpdateRoom(id string, roomUpdate *models.RoomEdit) (*models.HueActionResponse, error) {
	urlSuffix := fmt.Sprintf("resource/room/%s", id)
	return doActionRequest(s.client, http.MethodPut, urlSuffix, roomUpdate)
}

func (s *RoomService) CreateRoom(room *models.RoomEdit) (*models.HueActionResponse, error) {
	urlSuffix := "resource/room"
	return doActionRequest(s.client, http.MethodPost, urlSuffix, room)
}

func (s *RoomService) DeleteRoom(id string) (*models.HueActionResponse, error) {
	urlSuffix := fmt.Sprintf("resource/room/%s", id)
	return doActionRequest(s.client, http.MethodDelete, urlSuffix, nil)
}
