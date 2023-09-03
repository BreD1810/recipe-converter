package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/BreD1810/recipe-converter/internal/recipeconverter"
)

func main() {
	flag.Parse()
	recipeURL := flag.Arg(0)
	if recipeURL == "" {
		fmt.Fprintln(os.Stderr, "No recipe URL provided")
		os.Exit(1)
	}

	res, err := recipeconverter.ConvertRecipe(recipeURL)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error converting recipe")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, line := range res {
		fmt.Println(line)
	}
}
