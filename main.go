package main

import "fmt"

func main() {

	nutriScore := GetNuTritionalScore(NutritionalData{
		Energy:              EnergyFromKcal(),
		Sugar:               SugarGram(0),
		SaturatedFattyAcids: SaturatedFattyAcids(),
		Sodium:              SodiumMilligram(),
		Fruits:              FruitsPercent(),
		Fibre:               FibreGram(),
		Protein:             ProteinGram(),
	}, Food)

	fmt.Printf("Nutritional Score: %d\n", nutriScore.Value)
	fmt.Printf("NutriScore: %s\n", nutriScore.GetNutriScore())
}
