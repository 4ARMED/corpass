package generators

import (
	"context"
	"fmt"
	"strings"

	"slices"
)

type Generator interface {
	Generate(string) ([]string, error)
	Name() string
}

type GeneratorEngine struct {
	results    []string
	generators []Generator
}

type GeneratorEngineOption func(*GeneratorEngine)

func NewGeneratorEngine(options ...GeneratorEngineOption) *GeneratorEngine {
	g := &GeneratorEngine{
		generators: []Generator{},
	}

	for _, option := range options {
		if option == nil {
			continue
		}

		option(g)
	}

	return g
}

func (ge *GeneratorEngine) RegisterGenerator(g Generator) error {
	for _, existing := range ge.generators {
		if existing.Name() == g.Name() {
			return fmt.Errorf("generator with name %s already exists", g.Name())
		}
	}

	ge.generators = append(ge.generators, g)

	return nil
}

func (ge *GeneratorEngine) MustRegisterGenerator(g Generator) {
	err := ge.RegisterGenerator(g)
	if err != nil {
		panic(err)
	}
}

func (ge *GeneratorEngine) Generate(ctx context.Context, generatorsToUse []string, input string) ([]string, error) {
	// seed the results with our initial input
	if len(ge.results) == 0 {
		ge.results = append(ge.results, input)
	}

	for _, generatorName := range generatorsToUse {
		for _, g := range ge.generators {
			if g.Name() != strings.TrimSpace(generatorName) {
				continue
			}

			for _, i := range ge.results {
				result, err := g.Generate(i)
				if err != nil {
					return nil, err
				}

				ge.results = append(ge.results, result...)
			}
		}

	}

	// De-dupe
	slices.Sort(ge.results)
	return slices.Compact(ge.results), nil
}

func (ge *GeneratorEngine) Validate(generatorNames []string) error {
	for _, name := range generatorNames {
		found := false
		for _, g := range ge.generators {
			if g.Name() == name {
				found = true
			}
		}
		if !found {
			return fmt.Errorf("invalid generator name: %s", name)
		}
	}

	return nil
}
