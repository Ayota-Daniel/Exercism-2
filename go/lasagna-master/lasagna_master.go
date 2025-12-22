package lasagna

import (
	"slices"
)

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, averageLayerTime int) int {
	if averageLayerTime == 0 {
		averageLayerTime = 2
	}
	return len(layers) * averageLayerTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noodleQty int
	var sauceQty float64

	for _, layer := range layers {
		if layer == "noodles" {
			noodleQty += 50
			continue
		}
		if layer == "sauce" {
			sauceQty += 0.2
		}
	}

	return noodleQty, sauceQty
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendIngrdients []string, ownIngredients []string) {
	if ownIngredients[len(ownIngredients)-1] != "?" {
		ownIngredients = append(ownIngredients, "?")
	}
	ownSize := len(ownIngredients)
	ownIngredients = slices.Replace(ownIngredients, ownSize-1, ownSize, friendIngrdients[len(friendIngrdients)-1])
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portion int) []float64 {
	var newAmounts []float64
	for i := 0; i < len(amounts); i++ {
		newAmounts = append(newAmounts, (amounts[i]/2)*float64(portion))
	}
	return newAmounts
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
