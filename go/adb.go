package cmandroid

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// AdbInterface has run function for running adb commands
type AdbInterface interface {
	Run(device string, args []string) (string, error)
}

// Adb implements AdbInterface + path to binary
type Adb struct {
	AdbBinary string // path to binary
}

// Run run an adb command
func (adb Adb) Run(device string, args []string) (string, error) {
	var cmd *exec.Cmd
	if device != "" {
		argsWithDevice := append([]string{"-d", device}, args...)
		cmd = exec.Command(adb.AdbBinary, argsWithDevice...)
	} else {
		cmd = exec.Command(adb.AdbBinary, args...)
	}
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	commandOutput := stdout.String()
	return commandOutput, err
}

// AdbStartServer used to make sure the adb server is ready
func AdbStartServer(adb AdbInterface) {
	adb.Run("", []string{"start-server"})
}

// AdbGetDevices `adb devices`, return status
func AdbGetDevices(adb AdbInterface) (ListOfDevices, error) {
	commandOutput, err := adb.Run("", []string{"devices"})
	splitCommandOutput := strings.Split(commandOutput, "\n")
	if splitCommandOutput[0] == "List of devices attached" {
		var deviceIds []string
		for _, deviceLine := range splitCommandOutput[1:] {
			deviceID := strings.Split(deviceLine, "")
			deviceIds = append(deviceID, deviceIds...)
		}
		return ListOfDevices{deviceIds}, nil
	}
	return ListOfDevices{}, err
}

// AdbTap send a tap event to the device
func AdbTap(adb AdbInterface, deviceID string, x int, y int) error {
	args := []string{"input", "tap", fmt.Sprint(x), fmt.Sprint(y)}
	_, err := adb.Run(deviceID, args)
	return err
}
