package textable

type Align uint8

const (
	ALIGN_LEFT Align = iota
	ALIGN_CENTER
	ALIGN_RIGHT
	ALIGN_DECIMAL
)

type Rule uint8

const (
	RULE_FRAME Rule = iota
	RULE_ALL
	RULE_NONE
	RULE_HEADER
)

type TextTransform uint8

const (
	TT_NONE TextTransform = iota
	TT_CAPITALIZE
	TT_UPPERCASE
	TT_LOWERCASE
)
