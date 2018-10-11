package swagger

type ListOfDevices struct {
	Devices []string `json:"devices,omitempty"`
}

// MockListOfDevices used for mocking purposes
var MockListOfDevices = ListOfDevices{[]string{"F39C7C991AD0", "5D05814EF3F1", "DF172F869D0C"}}
