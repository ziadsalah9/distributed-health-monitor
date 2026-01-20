package worker

import "testing"


func TestEvaluateStatus(t *testing.T) {

	tests := []struct {
		code int
		expected string
	}{
		{200, "UP"},
		{201, "UP"},
		{404, "DOWN"},
		{500, "DOWN"},
	}

	for _, tt := range tests {

		result := EvaluateStatus(tt.code)

		if result != tt.expected {
			t.Errorf("Expected %s got %s", tt.expected, result)
		}
	}
}
