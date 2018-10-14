package cmandroid

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/shlex"
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
	tap := TapData{}
	json.NewDecoder(r.Body).Decode(&tap)
	command := fmt.Sprintf("adb -d %s shell input tap %d %d", tap.DeviceID, tap.X, tap.Y)
	tokens, _ := shlex.Split(command)
	log.Printf("$ %v", tokens)
	w.WriteHeader(http.StatusOK)
}
