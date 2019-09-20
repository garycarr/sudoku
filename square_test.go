package soduku

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRegion(t *testing.T) {
	tt := []struct {
		description    string
		pos            position
		expectedRegion region
	}{
		{
			description: "top left one",
			pos: position{
				rowNumber: 0,
				colNumber: 0,
			},
			expectedRegion: region{
				minRowNumber: 0,
				maxRowNumber: 2,
				minColNumber: 0,
				maxColNumber: 2,
			},
		},
		{
			description: "top left two",
			pos: position{
				rowNumber: 1,
				colNumber: 2,
			},
			expectedRegion: region{
				minRowNumber: 0,
				maxRowNumber: 2,
				minColNumber: 0,
				maxColNumber: 2,
			},
		},
		{
			description: "top right one",
			pos: position{
				rowNumber: 1,
				colNumber: 7,
			},
			expectedRegion: region{
				minRowNumber: 0,
				maxRowNumber: 2,
				minColNumber: 6,
				maxColNumber: 8,
			},
		},
		{
			description: "bottom right one",
			pos: position{
				rowNumber: 8,
				colNumber: 7,
			},
			expectedRegion: region{
				minRowNumber: 6,
				maxRowNumber: 8,
				minColNumber: 6,
				maxColNumber: 8,
			},
		},
		{
			description: "middle one",
			pos: position{
				rowNumber: 3,
				colNumber: 4,
			},
			expectedRegion: region{
				minRowNumber: 3,
				maxRowNumber: 5,
				minColNumber: 3,
				maxColNumber: 5,
			},
		},
	}
	for _, td := range tt {
		t.Run(td.description, func(t *testing.T) {
			s := square{}
			err := s.getRegion(td.pos)
			require.Nil(t, err)
			assert.Equal(t, td.expectedRegion, s.reg)
		})
	}
}
