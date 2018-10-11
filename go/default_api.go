package swagger

import (
	"encoding/json"
	"log"
	"net/http"
)

// GetDevicelist gives the list of attached devices
func GetDevicelist(w http.ResponseWriter, r *http.Request) {
	output, _ := json.Marshal(MockListOfDevices)
	log.Printf("$ adb devices")
	w.Write(output)
	w.WriteHeader(http.StatusOK)
}

// PostTap read in an x,y,deviceId and run that on the device
func PostTap(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
