package column

type Column interface {
	Len() uint
	Append(row interface{}) error
	At(index uint) (value interface{}, exists bool)
}
