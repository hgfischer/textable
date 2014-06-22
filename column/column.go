package column

type Column interface {
	Len() uint
	Append(row interface{}) error
	At(index uint) (value interface{}, exists bool)
	Last() (value interface{})
	String() string
}

func NewForTypeOf(value interface{}) Column {
	switch value.(type) {
	case string:
		return new(StringCol)
	case float64, float32:
		return new(FloatCol)
	case uint, uint8, uint16, uint32, uint64:
		return new(UIntCol)
	case int, int8, int16, int32, int64:
		return new(IntCol)
	default:
		return new(StringCol)
	}
}
