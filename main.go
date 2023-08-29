package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/4armed/corppass/generators"
	"github.com/4armed/corppass/generators/leet"
	"github.com/4armed/corppass/generators/numerics"
	"github.com/4armed/corppass/generators/punctuation"
	"github.com/4armed/corppass/generators/upperlower"
)

var (
	executable    string = "corpass"
	generatorList string
	verbose       bool
	logLevel      *slog.LevelVar = new(slog.LevelVar)
)

func main() {
	flag.StringVar(&generatorList, "generators", "upperlower,leet,numerics,punctuation", "comma separated list of generators to use")
	flag.BoolVar(&verbose, "verbose", false, "enable verbose logging")
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Printf("Usage: %s [name]\n", executable)
		return
	}

	input := flag.Args()[0]

	if verbose {
		logLevel.Set(slog.LevelDebug)
	}

	loggerOptions := &slog.HandlerOptions{
		Level: logLevel,
	}

	logger := slog.New(slog.NewTextHandler(os.Stderr, loggerOptions))

	g := generators.NewGeneratorEngine()
	g.MustRegisterGenerator(upperlower.NewUpperLowerGenerator())
	g.MustRegisterGenerator(leet.NewLeetGenerator())
	g.MustRegisterGenerator(numerics.NewNumericsGenerator())
	g.MustRegisterGenerator(punctuation.NewPunctuationGenerator())

	// Process the generator list
	generators := strings.Split(generatorList, ",")

	logger.Debug("using generators", "names", generators)
	err := g.Validate(generators)
	if err != nil {
		fmt.Println(err)
		return
	}

	results, err := g.Generate(context.Background(), generators, input)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
