package swagger

import (
	"encoding/json"
	"log"
	"net/http"
)

// GetDevicelist gives the list of attached devices
func GetDevicelist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	output, _ := json.Marshal(MockListOfDevices)
	log.Printf("$ adb devices")
	w.Write(output)
	w.WriteHeader(http.StatusOK)
}

func PostTap(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
