package textable

type TexTable struct {
	Headers []string
	Align   []Align
	Columns []*column
}

func New(headers ...string) *TexTable {
	return &TexTable{
		Headers: headers,
		Align:   make([]Align, len(headers)),
		Columns: make([]*column, len(headers)),
	}
}

func (tt *TexTable) AddRow(columns ...interface{}) error {
	if len(columns) != len(tt.Headers) {
		return getError(_ERR_INCOR_NUM_VALS, len(columns), len(tt.Headers))
	}

	for k, v := range columns {
		if tt.Columns[k] == nil {
			tt.Columns[k] = new(column)
		}
		tt.Columns[k].Append(v)
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
