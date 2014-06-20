package textable

type TextTransform uint8

const (
	TT_NONE TextTransform = iota
	TT_CAPITALIZE
	TT_UPPERCASE
	TT_LOWERCASE
)
