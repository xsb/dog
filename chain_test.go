package dog

import (
	"testing"
	"time"
)

func TestFormatDurationReturnsProperlyFormattedString(t *testing.T) {
	cases := []struct {
		d    time.Duration
		want string
	}{
		{
			25*time.Hour + 59*time.Minute + 59*time.Second,
			"25h59m59s",
		},
		{
			1*time.Hour + 10*time.Minute,
			"1h10m",
		},
		{
			10 * time.Second,
			"10s",
		},
		{
			1*time.Second + 500*time.Millisecond,
			"2s",
		},
		{
			1*time.Second + 1*time.Millisecond,
			"1s",
		},
		{
			1 * time.Millisecond,
			"0.001s",
		},
		{
			200 * time.Nanosecond,
			"0.000s",
		},
	}

	for _, c := range cases {
		actual := formatDuration(c.d)
		if actual != c.want {
			t.Errorf("Got time string: %q, but expected: %q", actual, c.want)
		}
	}
}
