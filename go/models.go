package cmandroid

// ListOfDevices {"devices":["F39C7C991AD0","5D05814EF3F1","DF172F869D0C"]}
type ListOfDevices struct {
	Devices []string `json:"devices,omitempty"`
}

// TapData a device and coordinate to perform a tap
type TapData struct {
	DeviceID string `json:"deviceId"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
}

// KeyboardData a device and keyboard data (text)
type KeyboardData struct {
	DeviceID string `json:"deviceId"`
	Text     string `json:"text"`
}

// PackageData package to be run
type PackageData struct {
	DeviceID    string `json:"deviceId"`
	PackageName string `json:"packageName"`
}
