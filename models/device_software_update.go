package models

type DeviceSoftwareUpdate struct {
	ID          string `json:"id,omitempty"`
	Owner       Owner  `json:"owner"`
	State       string `json:"state,omitempty"`
	LastInstall string `json:"last_install,omitempty"`
}
