package punctuation

import "fmt"

type PunctuationGenerator struct {
	charset []string
}

func NewPunctuationGenerator() *PunctuationGenerator {
	return &PunctuationGenerator{
		charset: []string{"!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "_", "=", "+", "[", "]", "{", "}", "|", "\\", ";", ":", "'", "\"", ",", "<", ".", ">", "/", "?", "`", "~"},
	}
}

func (g *PunctuationGenerator) Name() string {
	return "punctuation"
}

func (g *PunctuationGenerator) WithCharset(charset []string) {
	g.charset = charset
}

func (g *PunctuationGenerator) Generate(input string) ([]string, error) {
	results := []string{}

	for _, i := range g.charset {
		results = append(results, fmt.Sprintf("%s%s", input, i))
		results = append(results, fmt.Sprintf("%s%s", i, input))
	}

	return results, nil
}
