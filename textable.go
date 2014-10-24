package textable

type Textable struct {
	Headers []string
	Align   []Align
	Columns []*column
}

func New(headers ...string) *Textable {
	return &Textable{
		Headers: headers,
		Align:   make([]Align, len(headers)),
		Columns: make([]*column, len(headers)),
	}
}

func (tt *Textable) AddRow(columns ...interface{}) error {
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
