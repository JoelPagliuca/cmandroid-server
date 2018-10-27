package cmandroid

import (
	"reflect"
	"testing"

	"github.com/JoelPagliuca/cmandroid-server/go/mocks"
)

// MockListOfDevices used for mocking purposes
var listOfDevices1String = `List of devices attached

F39C7C991AD0	device
5D05814EF3F1	device
DF172F869D0C	device`
var listOfDevices1 = ListOfDevices{[]string{"F39C7C991AD0", "5D05814EF3F1", "DF172F869D0C"}}
var listOfDevices2String = `List of devices attached

`
var listOfDevices2 = ListOfDevices{}

func TestAdbGetDevices(t *testing.T) {
	t.Run("Normal scenario", func(t *testing.T) {
		mockAdb := &mocks.AdbInterface{}
		mockAdb.On("Run", "", []string{"devices"}).Return(listOfDevices1String, nil)
		want := listOfDevices1
		got, err := AdbGetDevices(mockAdb)
		if err != nil {
			t.Errorf("AdbGetDevices() error = %v", err)
			return
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("AdbGetDevices() = %v, want %v", got, want)
		}
	})
	t.Run("No devices", func(t *testing.T) {
		mockAdb := &mocks.AdbInterface{}
		mockAdb.On("Run", "", []string{"devices"}).Return(listOfDevices2String, nil)
		want := listOfDevices2
		got, err := AdbGetDevices(mockAdb)
		if err != nil {
			t.Errorf("AdbGetDevices() error = %v", err)
			return
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("AdbGetDevices() = %v, want %v", got, want)
		}
	})
}
