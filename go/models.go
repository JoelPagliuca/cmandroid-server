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

// MockListOfDevices used for mocking purposes
var MockListOfDevices = ListOfDevices{[]string{"F39C7C991AD0", "5D05814EF3F1", "DF172F869D0C"}}
