package main

import "testing"

func TestAbs(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  float64
	}{
		{
			name:  "positive value int",
			value: 1,
			want:  1,
		},
		{
			name:  "positive value float",
			value: 1.2,
			want:  1.2,
		},
		{
			name:  "negative value",
			value: -221,
			want:  221,
		},
		{
			name:  "negative value float",
			value: -22.1222,
			want:  22.1222,
		},
		{
			name:  "zero",
			value: 0,
			want:  0,
		},

		{
			name:  "negative zero",
			value: -0,
			want:  0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(y *testing.T) {
			if absVal := Abs(test.value); absVal != test.want {
				t.Errorf("Abs() = %f, want %f", absVal, test.want)
			}

		})
	}
}
