package models

type Device struct {
	ID          string               `json:"id"` // E.g. from the Light.Owner.ID
	IDV1        string               `json:"id_v1,omitempty"`
	Type        string               `json:"type"` // "device"
	ProductData DeviceProductData    `json:"product_data"`
	Metadata    DeviceMetadata       `json:"metadata"`
	Identify    *Identify            `json:"identify,omitempty"`
	UserTest    *UserTest            `json:"usertest,omitempty"`
	DeviceMode  *DeviceMode          `json:"device_mode,omitempty"`
	Services    []ResourceIdentifier `json:"services"`
}

type DevicePut struct {
	Type       string          `json:"type,omitempty"` // "device"
	Metadata   *DeviceMetadata `json:"metadata,omitempty"`
	Identify   *Identify       `json:"identify,omitempty"`
	UserTest   *UserTest       `json:"usertest,omitempty"`
	DeviceMode *DeviceMode     `json:"device_mode,omitempty"`
}

type DeviceProductData struct {
	ModelID              string `json:"model_id"`
	ManufacturerName     string `json:"manufacturer_name"`
	ProductName          string `json:"product_name"`
	ProductArchetype     string `json:"product_archetype"`
	Certified            bool   `json:"certified"`
	SoftwareVersion      string `json:"software_version"`
	HardwarePlatformType string `json:"hardware_platform_type,omitempty"`
}

type DeviceMetadata struct {
	Name      *string `json:"name,omitempty"`
	Archetype *string `json:"archetype,omitempty"`
}

type UserTest struct {
	Status   string `json:"status,omitempty"` // set, changing
	UserTest bool   `json:"usertest"`
}

type DeviceMode struct {
	Status     string   `json:"status,omitempty"`
	Mode       string   `json:"mode"`
	ModeValues []string `json:"mode_values,omitempty"`
}
