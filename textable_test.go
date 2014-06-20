package textable

import (
	"testing"
)

func TestTableSetup(t *testing.T) {
	table := New("City name", "Area", "Population", "Annual Rainfall")
	table.Align[0] = ALIGN_LEFT
	table.Align[3] = ALIGN_DECIMAL
	table.AddRow("Adelaide", 1295, 1158259, 600.5)
	table.AddRow("Brisbane", 5905, 1857594, 1146.4)
	table.AddRow("Darwin", 112, 120900, 1714.7)
	table.AddRow("Hobart", 1357, 205556, 619.5)
	table.AddRow("Sydney", 2058, 4336374, 1214.8)
	table.AddRow("Melbourne", 1566, 3806092, 646.9)
	table.AddRow("Perth", 5386, 1554769, 869.4)
}
