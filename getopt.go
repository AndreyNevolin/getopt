package getopt

type Option struct {
	// Long name of the option
	LongName string
	// Short (or single-character) name of the option
	ShortName rune
}

type Parser struct {
	args []string
}

func NewParser(args []string) *Parser {
	parser_p := &Parser{
		args: args,
	}

	return parser_p
}
