package upperlower

import (
	"strings"
)

type UpperLowerGenerator struct{}

func NewUpperLowerGenerator() *UpperLowerGenerator {
	return &UpperLowerGenerator{}
}

func (g *UpperLowerGenerator) Name() string {
	return "upperlower"
}

func (g *UpperLowerGenerator) Generate(input string) ([]string, error) {
	// strLen := len(input)

	results := []string{}

	generatePermutations(input, "", 0, &results)

	return results, nil
}

func generatePermutations(input string, currentPerm string, index int, results *[]string) {
	if index == len(input) {
		*results = append(*results, currentPerm)
		return
	}

	generatePermutations(input, currentPerm+strings.ToLower(string(input[index])), index+1, results)
	generatePermutations(input, currentPerm+strings.ToUpper(string(input[index])), index+1, results)
}
