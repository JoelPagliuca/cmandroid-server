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

func AdbTap(deviceId string, x int, y int) int {
	cmd := exec.Command(AdbBinary, "-d", deviceId, "input", "tap", fmt.Sprint(x), fmt.Sprint(y))
	cmd.Run()
	return 0
}
