package numerics

import "fmt"

type NumericsGenerator struct {
	start int
	end   int
}

func NewNumericsGenerator() *NumericsGenerator {
	return &NumericsGenerator{
		start: 0,
		end:   9,
	}
}

func (g *NumericsGenerator) Name() string {
	return "numerics"
}

func (g *NumericsGenerator) WithStart(i int) {
	g.start = i
}

func (g *NumericsGenerator) WithEnd(i int) {
	g.end = i
}

func (g *NumericsGenerator) Generate(input string) ([]string, error) {
	results := []string{}

	for i := g.start; i <= g.end; i++ {
		results = append(results, fmt.Sprintf("%s%d", input, i))
		results = append(results, fmt.Sprintf("%d%s", i, input))
	}

	return results, nil
}
