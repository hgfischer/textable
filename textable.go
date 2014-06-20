package textable

type TexTable struct {
	Header []string
	Align  []Align
	Data   [][]interface{}
}

func New(header ...string) *TexTable {
	return &TexTable{
		Header: header,
		Align:  make([]Align, len(header)),
	}
}

func (t *TexTable) AddRow(columns ...interface{}) error {
	for k, v := range columns {
		println(k, v)
	}

	return nil
}

func (t *TexTable) String() string {
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
