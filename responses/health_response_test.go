package responses

import (
	"reflect"
	"testing"
)

func TestHealth(t *testing.T) {
	tests := []struct {
		name string
		want *HealthCheck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Health(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Health() = %v, want %v", got, tt.want)
			}
		})
	}
}
