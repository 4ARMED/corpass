package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/4armed/corpass/generators"
	"github.com/4armed/corpass/generators/leet"
	"github.com/4armed/corpass/generators/numerics"
	"github.com/4armed/corpass/generators/punctuation"
	"github.com/4armed/corpass/generators/upperlower"
)

var (
	Executable    string = "corpass"
	Version       string = ""
	generatorList string
	verbose       bool
	version       bool
	logLevel      *slog.LevelVar = new(slog.LevelVar)
)

func main() {
	flag.StringVar(&generatorList, "generators", "upperlower,leet,numerics,punctuation", "comma separated list of generators to use")
	flag.BoolVar(&verbose, "verbose", false, "enable verbose logging")
	flag.BoolVar(&version, "version", false, "print version and exit")
	flag.Usage = func() {
		fmt.Printf("Usage: %s [string]\n\n", Executable)
		flag.PrintDefaults()
	}
	flag.Parse()

	if version {
		fmt.Printf("%s %s", Executable, Version)
		return
	}

	if len(flag.Args()) == 0 {
		flag.Usage()
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
