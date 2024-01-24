package multiplier

import "fmt"

type MultiplierGenerator struct {
	count int
}

func NewMultiplierGenerator() *MultiplierGenerator {
	return &MultiplierGenerator{
		count: 2,
	}
}

func (g *MultiplierGenerator) Name() string {
	return "multiplier"
}

func (g *MultiplierGenerator) WithCount(i int) {
	g.count = i
}

func (g *MultiplierGenerator) Generate(input string) ([]string, error) {
	results := []string{}

	result := input
	for i := 2; i <= g.count; i++ {
		result = fmt.Sprintf("%s%s", result, input)
		results = append(results, result)
	}

	return results, nil
}
