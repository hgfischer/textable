package textable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	TEST_HEADERS = []string{"Country", "Population", "Educ. Exp. (% GDP)", "HDI", "HDI change (*1000)"}
	TEST_DATA    = [][]interface{}{
		{"Algeria", uint(38700000), 4.3, 0.713, nil},
		{"Argentina", uint(42669500), 6.5, 0.811, 6},
		{"Australia", uint(23536500), 4.5, 0.938, 3},
		{"Brazil", uint(202756000), 5.1, 0.730, 4},
		{"Canada", uint(35427524), 5.5, 0.911, 2},
		{"China", uint(1365220000), nil, 0.699, 10},
		{"Greece", uint(10816286), 4.0, 0.860, -6},
		{"India", uint(1245830000), 3.1, 0.554, 7},
		{"Kazakhstan", uint(17244000), 2.8, "?", 10},
		{"Portugal", uint(10477800), 5.2, 0.816, -1},
		{"Russia", uint(146000000), 3.9, 0.788, 6},
		{"United States", uint(318275000), 5.5, 0.937, 3},
	}
)

func TestNewTableShouldStoreHeadersAndInitializeColumns(t *testing.T) {
	table := New(TEST_HEADERS...)
	assert.Equal(t, len(TEST_HEADERS), len(table.Headers))
	assert.Equal(t, len(TEST_HEADERS), len(table.Columns))
}

func TestAddRowShouldSetProperTypeForEachColumnAndAppendValues(t *testing.T) {
	tt := New(TEST_HEADERS...)
	for nrow, data := range TEST_DATA {
		err := tt.AddRow(data...)
		assert.Nil(t, err)
		for ncol, _ := range data {
			assert.Equal(t, nrow+1, tt.Columns[ncol].Len())
		}
	}
}

func TestAddRowWithWrongNumberOfColumnsShouldReturnError(t *testing.T) {
	tt := New(TEST_HEADERS...)
	err := tt.AddRow()
	assert.NotNil(t, err)
}
