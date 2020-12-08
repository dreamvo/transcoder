package utils

import "testing"

func TestUtils(t *testing.T) {
	t.Run("DurationToSeconds", func(t *testing.T) {
		suite := map[string]float64{
			"00:00:05.31": 5.310000,
			"00:02:14.31": 134.310000,
			"00:00:27.25": 27.25,
			"invalid":     0,
		}

		for duration, expected := range suite {
			if e := DurationToSeconds(duration); e != expected {
				t.Errorf("expected: %f, got: %f", expected, e)
			}
		}
	})
}
