package cmandroid

import (
	"fmt"
	"os/exec"
)

var AdbBinary string

// AdbGetDevices `adb devices`, return status
func AdbGetDevices() (ListOfDevices, int) {
	return MockListOfDevices, 0
}

// AdbTap send a tap event to the device
func AdbTap(deviceId string, x int, y int) error {
	cmd := exec.Command(AdbBinary, "-d", deviceId, "input", "tap", fmt.Sprint(x), fmt.Sprint(y))
	err := cmd.Run()
	return err
}
