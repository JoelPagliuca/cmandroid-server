package cmandroid

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/shlex"
)

// App has everything we'll need for the api
type App struct {
	Adb AdbInterface
}

// GetDevicelist gives the list of attached devices
func (app *App) GetDevicelist(w http.ResponseWriter, r *http.Request) {
	deviceList, _ := AdbGetDevices(app.Adb)
	output, _ := json.Marshal(deviceList)
	log.Printf("$ adb devices")
	w.Write(output)
	w.WriteHeader(http.StatusOK)
}

// PostTap read in an x,y,deviceId and run that on the device
func (app *App) PostTap(w http.ResponseWriter, r *http.Request) {
	tap := TapData{}
	json.NewDecoder(r.Body).Decode(&tap)
	command := fmt.Sprintf("adb -d %s shell input tap %d %d", tap.DeviceID, tap.X, tap.Y)
	AdbTap(app.Adb, tap.DeviceID, tap.X, tap.Y)
	tokens, _ := shlex.Split(command)
	log.Printf("$ %v", tokens)
	w.WriteHeader(http.StatusOK)
}

// Index generic index page
func (app *App) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "cmandroid server")
}
