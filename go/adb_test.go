package cmandroid

import (
	"reflect"
	"testing"

	"github.com/JoelPagliuca/cmandroid-server/go/mocks"
)

// MockListOfDevices used for mocking purposes
var listOfDevices1 = ListOfDevices{[]string{"F39C7C991AD0", "5D05814EF3F1", "DF172F869D0C"}}

func TestAdbGetDevices(t *testing.T) {
	mockAdb := &mocks.AdbInterface{}
	mockAdb.On("Run", "", []string{"devices"}).Return(`List of devices attached

F39C7C991AD0	device
5D05814EF3F1	device
DF172F869D0C	device`, nil)
	type args struct {
		adb AdbInterface
	}
	tests := []struct {
		name    string
		args    args
		want    ListOfDevices
		wantErr bool
	}{
		{
			"Normal scenario",
			args{adb: mockAdb},
			listOfDevices1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AdbGetDevices(tt.args.adb)
			if (err != nil) != tt.wantErr {
				t.Errorf("AdbGetDevices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AdbGetDevices() = %v, want %v", got, tt.want)
			}
		})
	}
}
