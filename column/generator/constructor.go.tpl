package column

func NewForTypeOf(value interface{}) Column {
	switch value.(type) {
	{{range .Types}}
	case {{.Type}}:
		return new({{.CapsName}})
	{{end}}
	default:
		return new({{.DefaultType.CapsName}})
	}
	return nil
}