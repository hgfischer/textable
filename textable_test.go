package textable

import (
	"testing"
)

func TestTableSetup(t *testing.T) {
	tt := New("City name", "Area", "Population", "Annual Rainfall")
	tt.Align[0] = ALIGN_LEFT
	tt.Align[3] = ALIGN_DECIMAL
	tt.AddRow("Adelaide", 1295, 1158259, 600.5)
	tt.AddRow("Brisbane", 5905, 1857594, 1146.4)
	tt.AddRow("Darwin", 112, 120900, 1714.7)
	tt.AddRow("Hobart", 1357, 205556, 619.5)
	tt.AddRow("Sydney", 2058, 4336374, 1214.8)
	tt.AddRow("Melbourne", 1566, 3806092, 646.9)
	tt.AddRow("Perth", 5386, 1554769, 869.4)
	for _, c := range tt.Columns {
		t.Log(c.String())
	}
}
