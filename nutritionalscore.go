package main

type ScoreType int

const (
	Food ScoreType = iota
	Beverage
	Water
	Cheese
)

type NutritionalScore struct {
	Value     int
	Positive  int
	Negative  int
	ScoreType ScoreType
}

type EnergyKJ float64

type SugarGram float64

type SaturatedFattyAcids float64

type SodiumMilligram float64

type FruitsPercent float64

type FibreGram float64

type ProteinGram float64

type NutritionalData struct {
	Energy              EnergyKJ
	Sugar               SugarGram
	SaturatedFattyAcids SaturatedFattyAcids
	Sodium              SodiumMilligram
	Fruits              FruitsPercent
	Fibre               FibreGram
	Protein             ProteinGram
	IsWater             bool
}

func GetNutritionalScore(data NutritionalData, scoreType ScoreType) NutritionalScore {

	value := 0
	positive := 0
	negative := 0

	if scoreType != Water {
		fruitPoints := data.Fruits.GetPoints(scoreType)
		fibrePoints := data.Fibre.GetPoints(scoreType)

		negative = data.Energy.GetPoints(scoreType) + data.Sugar.GetPoints(scoreType) + data.SaturatedFattyAcids.GetPoints(scoreType) + data.Sodium.GetPoints(scoreType)
		positive = fruitPoints + fibrePoints + data.Protein.GetPoints(scoreType)
	}

	return NutritionalScore{
		Value:     value,
		Positive:  positive,
		Negative:  negative,
		ScoreType: scoreType,
	}
}
