package worker

import "testing"

func TestHasStateChanged(t *testing.T) {

	tests := []struct {
		old string
		new string
		expected bool
	}{
		{"UP", "DOWN", true},
		{"DOWN", "UP", true},
		{"UP", "UP", false},
		{"DOWN", "DOWN", false},
	}

	for _, tt := range tests {

		result := HasStateChanged(tt.old, tt.new)

		if result != tt.expected {
			t.Errorf("Expected %v got %v for %s -> %s",
				tt.expected, result, tt.old, tt.new)
		}
	}
}
