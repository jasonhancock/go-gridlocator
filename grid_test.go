package gridlocator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		lat      float64
		long     float64
		expected string
		err      error
	}{
		{0, 0, "JJ00aa", nil},
		{48.146666666666667, 11.608333333333333, "JN58td", nil}, // Munich
		{-34.91, -56.211666666666667, "GF15vc", nil},            // Montevideo
		{38.92, -77.065, "FM18lw", nil},                         // Washington, D.C.
		{-41.283333333333333, 174.745, "RE78ir", nil},           // Wellington
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%.5f %.5f", tt.lat, tt.long), func(t *testing.T) {
			result, err := Convert(tt.lat, tt.long)
			if tt.err != nil {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.err.Error())
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestConvertGridLocation(t *testing.T) {
	tests := []struct {
		location string
		lat      float64
		long     float64
		err      error
	}{
		{"JJ00aa", 0, 0, nil},
		{"JN58td", 48.146666666666667, 11.608333333333333, nil}, // Munich
		{"GF15vc", -34.91, -56.211666666666667, nil},            // Montevideo
		{"FM18lw", 38.92, -77.065, nil},                         // Washington, D.C.
		{"RE78ir", -41.283333333333333, 174.745, nil},           // Wellington
	}

	for _, tt := range tests {
		t.Run(tt.location, func(t *testing.T) {
			lat, long, err := ConvertGridLocation(tt.location)
			if tt.err != nil {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.err.Error())
				return
			}

			require.NoError(t, err)
			require.InDelta(t, tt.lat, lat, .1)
			require.InDelta(t, tt.long, long, .1)
		})
	}
}
