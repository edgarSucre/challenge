package misc_test

import (
	"testing"

	"github.com/edgarSucre/challenge/misc"
	"github.com/stretchr/testify/assert"
)

func TestSecondsToDate(t *testing.T) {
	tests := []struct {
		name     string
		seconds  int
		expected string
	}{
		{
			"just seconds",
			12,
			"12 seconds",
		},
		{
			"with minutes",
			190,
			"3 minutes, 10 seconds",
		},
		{
			"with hours",
			10845,
			"3 hours, 0 minutes, 45 seconds",
		},
		{
			"with days",
			5198580,
			"60 days, 4 hours, 3 minutes, 0 seconds",
		},

		{
			"with years",
			36734580,
			"1 years, 60 days, 4 hours, 3 minutes, 0 seconds",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := misc.SecondsToDate(tt.seconds)

			assert.Equal(t, tt.expected, actual)
		})
	}
}
