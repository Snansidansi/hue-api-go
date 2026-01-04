package hueapi

import (
	"fmt"
	"net/http"

	"github.com/snansidansi/hueapi/models"
	"github.com/snansidansi/hueapi/util"
)

type LightService struct {
	client *Client
}

func (l *LightService) GetAllLights() (*models.HueResponse[models.Light], error) {
	urlSuffix := "resource/light"
	return doGetRequest[models.Light](l.client, urlSuffix)
}

func (l *LightService) GetLightByID(id string) (*models.HueResponse[models.Light], error) {
	urlSuffix := fmt.Sprintf("resource/light/%s", id)
	return doGetRequest[models.Light](l.client, urlSuffix)
}

func (s *LightService) SetLightState(id string, update *models.LightPut) (*models.HueActionResponse, error) {
	urlSuffix := fmt.Sprintf("resource/light/%s", id)
	return doActionRequest(s.client, http.MethodPut, urlSuffix, update)
}

func (s *LightService) On(id string) (*models.HueActionResponse, error) {
	return s.SetOnOff(id, true)
}

func (s *LightService) Off(id string) (*models.HueActionResponse, error) {
	return s.SetOnOff(id, false)
}

func (s *LightService) SetOnOff(id string, on bool) (*models.HueActionResponse, error) {
	update := models.LightPut{
		On: &models.OnPut{
			On: &on,
		},
	}
	return s.SetLightState(id, &update)
}

func (s *LightService) Rename(id string, name string) (*models.HueActionResponse, error) {
	update := models.LightPut{
		Metadata: &models.MetadataPut{
			Name: &name,
		},
	}
	return s.SetLightState(id, &update)
}

func (s *LightService) SetBrightness(id string, brightness float64) (*models.HueActionResponse, error) {
	update := models.LightPut{
		Dimming: &models.DimmingPut{
			Brightness: &brightness,
		},
	}
	return s.SetLightState(id, &update)
}

func (s *LightService) SetColor(id string, r, g, b int) (*models.HueActionResponse, error) {
	x, y := util.RGBToXY(r, g, b)
	return s.SetColorXY(id, x, y)
}

func (s *LightService) SetColorXY(id string, x, y float64) (*models.HueActionResponse, error) {
	update := models.LightPut{
		Color: &models.ColorPut{
			XY: &models.XYPut{
				X: &x,
				Y: &y,
			},
		},
	}
	return s.SetLightState(id, &update)
}

func (s *LightService) SetTemperature(id string, mirek int) (*models.HueActionResponse, error) {
	update := models.LightPut{
		ColorTemperature: &models.ColorTemperaturePut{
			Mirek: &mirek,
		},
	}
	return s.SetLightState(id, &update)
}

func (s *LightService) Identify(id string, durationMs int64) (*models.HueActionResponse, error) {
	action := "identify"
	update := models.LightPut{
		Identify: &models.IdentifyPut{
			Action:   &action,
			Duration: &durationMs,
		},
	}
	return s.SetLightState(id, &update)
}
