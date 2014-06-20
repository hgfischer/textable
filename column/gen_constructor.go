package column

func NewForTypeOf(value interface{}) Column {
	switch value.(type) {

	case bool:
		return new(BoolCol)

	case uintptr:
		return new(UintptrCol)

	case complex128:
		return new(Complex128Col)

	case complex64:
		return new(Complex64Col)

	case float32:
		return new(Float32Col)

	case float64:
		return new(Float64Col)

	case int:
		return new(IntCol)

	case int8:
		return new(Int8Col)

	case int16:
		return new(Int16Col)

	case int32:
		return new(Int32Col)

	case int64:
		return new(Int64Col)

	case uint:
		return new(UintCol)

	case uint8:
		return new(Uint8Col)

	case uint16:
		return new(Uint16Col)

	case uint32:
		return new(Uint32Col)

	case uint64:
		return new(Uint64Col)

	case string:
		return new(StringCol)

	default:
		return new(StringCol)
	}
	return nil
}
