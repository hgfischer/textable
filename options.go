package textable

type Opts struct {
	Start             uint8
	End               uint8
	Fields            []string
	Header            bool
	Border            bool
	HRules            Rule
	VRules            Rule
	IntFormat         string
	FloatFormat       string
	PaddingWidth      uint8
	LeftPaddingWidth  uint8
	RightPaddingWidth uint8
	VerticalChar      rune
	HorizontalChar    rune
	JunctionChar      rune
	SortBy            string
	ReverseSort       bool
	PrintEmpty        bool
	XHTML             bool
}
