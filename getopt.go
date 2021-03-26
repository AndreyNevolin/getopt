package getopt

type Option struct {
	// Long name of the option
	LongName string
	// Short (or single-character) name of the option
	ShortName rune
	// Specifies how results are returned for an option. If "Flag" is nil, then
	// Parser.getopt() sets one of its return values to "Val". Otherwise, Parser.getopt()
	// sets the return value to 0, and memory location referenced by "Flag" is set to
	// "Val" if the option is found, but left unchanged if the option is not found
	Flag *int
	// The value to be returned by Parser.getopt(), or to be loaded into the memory
	// location pointed to by "Flag"
	Val int
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
