package camera

import (
	"reflect"
	"testing"
)

func TestNewCamera(t *testing.T) {
	tests := []struct {
		name string
		want *Camera
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCamera(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCamera() = %v, want %v", got, tt.want)
			}
		})
	}
}
