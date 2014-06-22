package textable

import "github.com/hgfischer/textable/column"

type TexTable struct {
	Header  []string
	Align   []Align
	Columns []column.Column
}

func New(header ...string) *TexTable {
	return &TexTable{
		Header:  header,
		Align:   make([]Align, len(header)),
		Columns: make([]column.Column, len(header)),
	}
}

func (tt *TexTable) AddRow(columns ...interface{}) error {
	if len(columns) != len(tt.Header) {
		return getError(_ERR_INCOR_NUM_VALS, len(columns), len(tt.Header))
	}

	for k, v := range columns {
		if tt.Columns[k] == nil {
			tt.Columns[k] = column.NewForTypeOf(v)
		}
		if err := tt.Columns[k].Append(v); err != nil {
			return err
		}
	}

	return nil
}

func (tt *TexTable) String() string {
	return `+-----------+------+------------+-----------------+
| City name | Area | Population | Annual Rainfall |
+-----------+------+------------+-----------------+
| Adelaide  | 1295 |  1158259   |      600.5      |
| Brisbane  | 5905 |  1857594   |      1146.4     |
| Darwin    | 112  |   120900   |      1714.7     |
| Hobart    | 1357 |   205556   |      619.5      |
| Sydney    | 2058 |  4336374   |      1214.8     |
| Melbourne | 1566 |  3806092   |      646.9      |
| Perth     | 5386 |  1554769   |      869.4      |
+-----------+------+------------+-----------------+`
}
