package recipeconverter

import (
	"fmt"

	"github.com/kkyr/go-recipe/pkg/recipe"
)

// ConvertRecipe gets the information for a recipe and returns a list of lines, or an error
func ConvertRecipe(url string) ([]string, error) {
	recipe, err := recipe.ScrapeURL(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching recipe: %w", err)
	}

	var lines []string

	if name, ok := recipe.Name(); ok {
		lines = append(lines, name)
	}
	lines = append(lines, "===\n")

	if description, ok := recipe.Description(); ok {
		lines = append(lines, description)
	}

	if image, ok := recipe.ImageURL(); ok {
		lines = append(lines, fmt.Sprintf("![](%s)\n", image))
	}

	lines = append(lines, fmt.Sprintf("[Recipe](%s)", url))

	if yields, ok := recipe.Yields(); ok {
		lines = append(lines, "Serves: "+yields)
	}

	if prepTime, ok := recipe.PrepTime(); ok {
		prepLine := fmt.Sprintf("Prep time: %d minutes", int(prepTime.Minutes()))
		lines = append(lines, prepLine)
	}

	if cookTime, ok := recipe.CookTime(); ok {
		cookLine := fmt.Sprintf("Cook time: %d minutes", int(cookTime.Minutes()))
		lines = append(lines, cookLine)
	}

	if totalTime, ok := recipe.TotalTime(); ok {
		totalLine := fmt.Sprintf("Total time: %d minutes", int(totalTime.Minutes()))
		lines = append(lines, totalLine)
	}

	lines = append(lines, "\n## Ingredients")
	if ingredients, ok := recipe.Ingredients(); ok {
		for _, ingredient := range ingredients {
			lines = append(lines, "- "+ingredient)
		}
	}

	lines = append(lines, "\n## Method")
	if instructions, ok := recipe.Instructions(); ok {
		for i, instruction := range instructions {
			lines = append(lines, fmt.Sprintf("%d. %s", i+1, instruction))
		}
	}

	lines = append(lines, "")

	return lines, nil
}
