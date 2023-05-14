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

func EnergyFromKcal(kcal float64) EnergyKJ {
	return EnergyKJ(kcal * 4.184)
}

func SodiumFromSalt(saltMg float64) SodiumMilligram {
	return SodiumMilligram(saltMg / 2.5)
}

func (eKJ EnergyKJ) GetPoints(scoreType ScoreType) int {

}

func (sGram SugarGram) GetPoints(scoreType ScoreType) int {

}

func (sFAcids SaturatedFattyAcids) GetPoints(scoreType ScoreType) int {

}

func (pGram ProteinGram) GetPoints(scoreType ScoreType) int {

}

func (sMilligram SodiumMilligram) GetPoints(scoreType ScoreType) int {

}

func (fPercent FruitsPercent) GetPoints(scoreType ScoreType) int {

}

func (fGram FibreGram) GetPoints(scoreType ScoreType) int {

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
