package column

type Column interface {
	Len() int
	Append(row interface{}) error
}
