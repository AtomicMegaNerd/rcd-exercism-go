package lasagna

const Noodles = "noodles"
const Sauce = "sauce"

func PreparationTime(layers []string, avgPrepTime int) int {
	if avgPrepTime == 0 {
		avgPrepTime = 2
	}

	return len(layers) * avgPrepTime
}

func Quantities(layers []string) (int, float64) {
	noodleAmount := 0
	sauceAmount := 0.0

	for _, layer := range layers {
		if layer == Noodles {
			noodleAmount += 50
		}
		if layer == Sauce {
			sauceAmount += 0.2
		}
	}
	return noodleAmount, sauceAmount
}

func AddSecretIngredient(friendsList []string, myList []string) []string {
	secretIngredient := friendsList[len(friendsList)-1]
	for ix, ingredient := range myList {
		if ingredient == "?" {
			myList[ix] = secretIngredient
		}
	}
	return myList
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	mul := float64(portions) / 2.0
	scaled := make([]float64, len(quantities))
	copy(scaled, quantities)
	for ix := range scaled {
		scaled[ix] *= mul
	}
	return scaled
}
