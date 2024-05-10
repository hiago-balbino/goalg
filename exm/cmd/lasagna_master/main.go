package main

func PreparationTime(layers []string, avgTimePerLayer int) int {
	if avgTimePerLayer == 0 {
		avgTimePerLayer = 2
	}

	return avgTimePerLayer * len(layers)
}

func Quantities(layers []string) (int, float64) {
	noodlesQtd := 0
	sauceQtd := 0.0

	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodlesQtd += 50
		case "sauce":
			sauceQtd += 0.2
		}
	}

	return noodlesQtd, sauceQtd
}

func AddSecretIngredient(friendIngredients, ownIngredients []string) {
	ownIngredients[len(ownIngredients)-1] = friendIngredients[len(friendIngredients)-1]
}

func ScaleRecipe(currentAmounts []float64, expectedNumberOfPortions int) []float64 {
	amountsNeeded := []float64{}

	for _, currentAmount := range currentAmounts {
		amountsNeeded = append(amountsNeeded, (currentAmount * float64(expectedNumberOfPortions) / 2))
	}

	return amountsNeeded
}
