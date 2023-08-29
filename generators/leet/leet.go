package leet

var (
	leetMap = map[string]string{
		"a": "@",
		"b": "8",
		"e": "3",
		"g": "6",
		"i": "1",
		"o": "0",
		"s": "5",
		"t": "7",
	}
)

type LeetGenerator struct{}

func NewLeetGenerator() *LeetGenerator {
	return &LeetGenerator{}
}

func (g *LeetGenerator) Name() string {
	return "leet"
}

func (g *LeetGenerator) Generate(input string) ([]string, error) {
	results := []string{}

	generatePermutations(input, "", 0, &results)

	return results, nil
}

func generatePermutations(input string, currentPerm string, index int, results *[]string) {
	if index == len(input) {
		*results = append(*results, currentPerm)
		return
	}

	generatePermutations(input, currentPerm+checkLeetMap(string(input[index])), index+1, results)
}

func checkLeetMap(input string) string {
	if c, ok := leetMap[input]; ok {
		return c
	}

	return input
}
