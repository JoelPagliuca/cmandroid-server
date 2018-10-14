package cmandroid

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// AdbBinary path to the adb binary
var AdbBinary string

// AdbRun run an adb command
func AdbRun(device string, args []string) (string, error) {
	var cmd *exec.Cmd
	if device != "" {
		argsWithDevice := append([]string{"-d", device}, args...)
		cmd = exec.Command(AdbBinary, argsWithDevice...)
	} else {
		cmd = exec.Command(AdbBinary, args...)
	}
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	commandOutput := stdout.String()
	return commandOutput, err
}

// AdbStartServer used to make sure the adb server is ready
func AdbStartServer() {
	cmd := exec.Command(AdbBinary, "start-server")
	cmd.Run()
}

// AdbGetDevices `adb devices`, return status
func AdbGetDevices() (ListOfDevices, error) {
	commandOutput, err := AdbRun("", []string{"devices"})
	splitCommandOutput := strings.Split(commandOutput, "\n")
	if splitCommandOutput[0] == "List of devices attached" {
		var deviceIds []string
		for _, deviceLine := range splitCommandOutput[1:] {
			deviceId := strings.Split(deviceLine, "")
			deviceIds = append(deviceId, deviceIds...)
		}
		return ListOfDevices{deviceIds}, nil
	} else {
		return ListOfDevices{}, err
	}
}

// AdbTap send a tap event to the device
func AdbTap(deviceId string, x int, y int) error {
	args := []string{"input", "tap", fmt.Sprint(x), fmt.Sprint(y)}
	_, err := AdbRun(deviceId, args)
	return err
}
