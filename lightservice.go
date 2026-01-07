package hueapi

import (
	"fmt"
	"net/http"

	"github.com/Snansidansi/hue-api-go/models"
	"github.com/Snansidansi/hue-api-go/util"
)

type lightService struct {
	client *Client
}

func (l *lightService) GetAllLights() (*models.HueResponse[models.Light], error) {
	urlSuffix := "resource/light"
	return doGetRequest[models.Light](l.client, urlSuffix)
}

func (l *lightService) GetLightByID(id string) (*models.HueResponse[models.Light], error) {
	urlSuffix := fmt.Sprintf("resource/light/%s", id)
	return doGetRequest[models.Light](l.client, urlSuffix)
}

func (s *lightService) SetLightState(id string, update *models.LightUpdate) (*models.HueActionResponse, error) {
	urlSuffix := fmt.Sprintf("resource/light/%s", id)
	return doActionRequest(s.client, http.MethodPut, urlSuffix, update)
}

func (s *lightService) On(id string) (*models.HueActionResponse, error) {
	return s.SetOnOff(id, true)
}

func (s *lightService) Off(id string) (*models.HueActionResponse, error) {
	return s.SetOnOff(id, false)
}

func (s *lightService) SetOnOff(id string, on bool) (*models.HueActionResponse, error) {
	update := models.LightUpdate{
		On: &models.On{
			On: &on,
		},
	}
	return s.SetLightState(id, &update)
}

func (s *lightService) Rename(id string, name string) (*models.HueActionResponse, error) {
	update := models.LightUpdate{
		Metadata: &models.Metadata{
			Name: &name,
		},
	}
	return s.SetLightState(id, &update)
}

func (s *lightService) SetBrightness(id string, brightness float64) (*models.HueActionResponse, error) {
	update := models.LightUpdate{
		Dimming: &models.Dimming{
			Brightness: &brightness,
		},
	}
	return s.SetLightState(id, &update)
}

func (s *lightService) SetColor(id string, r, g, b int) (*models.HueActionResponse, error) {
	x, y := util.RGBToXY(r, g, b)
	return s.SetColorXY(id, x, y)
}

func (s *lightService) SetColorXY(id string, x, y float64) (*models.HueActionResponse, error) {
	update := models.LightUpdate{
		Color: &models.Color{
			XY: &models.XY{
				X: x,
				Y: y,
			},
		},
	}
	return s.SetLightState(id, &update)
}

func (s *lightService) SetTemperature(id string, mirek int) (*models.HueActionResponse, error) {
	update := models.LightUpdate{
		ColorTemperature: &models.ColorTemperature{
			Mirek: &mirek,
		},
	}
	return s.SetLightState(id, &update)
}

func (s *lightService) Identify(id string, durationMs int64) (*models.HueActionResponse, error) {
	action := "identify"
	update := models.LightUpdate{
		Identify: &models.Identify{
			Action:   &action,
			Duration: &durationMs,
		},
	}
	return s.SetLightState(id, &update)
}
