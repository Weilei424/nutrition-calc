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

var enengyLevels = []float64{3350, 3015, 2680, 2345, 2010, 1675, 1340, 1005, 670, 335}
var sugarLevels = []float64{45, 40, 36, 31, 27, 22.5, 18, 13.5, 9, 4.5}
var saturatedFattyAcidsLevels = []float64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
var sodiumLevels = []float64{900, 710, 720, 630, 540, 450, 360, 270, 180, 90}
var fibreLevels = []float64{4.7, 3.7, 2.8, 1.9, 0.9}
var proteinLevels = []float64{8, 6.4, 4.8, 3.2, 1.6}

var energyLevelsBeverage = []float64{270, 240, 210, 180, 150, 120, 90, 60, 30, 0}
var sugarLevelsBeverage = []float64{13.5, 12, 10.5, 9, 7.5, 6, 4.5, 3, 1.5, 0}

func EnergyFromKcal(kcal float64) EnergyKJ {
	return EnergyKJ(kcal * 4.184)
}

func SodiumFromSalt(saltMg float64) SodiumMilligram {
	return SodiumMilligram(saltMg / 2.5)
}

func (eKJ EnergyKJ) GetPoints(scoreType ScoreType) int {
	if scoreType == Beverage {
		return getPointsFromRange(float64(eKJ), energyLevelsBeverage)
	}
	return getPointsFromRange(float64(eKJ), enengyLevels)
}

func (sGram SugarGram) GetPoints(scoreType ScoreType) int {
	if scoreType == Beverage {
		return getPointsFromRange(float64(sGram), sugarLevels)
	}
	return getPointsFromRange(float64(sGram), sugarLevels)
}

func (sFAcids SaturatedFattyAcids) GetPoints(scoreType ScoreType) int {
	if scoreType == Beverage {
		return getPointsFromRange(float64(sFAcids), saturatedFattyAcidsLevels)
	}
	return getPointsFromRange(float64(sFAcids), saturatedFattyAcidsLevels)
}

func (pGram ProteinGram) GetPoints(scoreType ScoreType) int {
	if scoreType == Beverage {
		return getPointsFromRange(float64(pGram), proteinLevels)
	}
	return getPointsFromRange(float64(pGram), proteinLevels)
}

func (sMilligram SodiumMilligram) GetPoints(scoreType ScoreType) int {
	if scoreType == Beverage {
		return getPointsFromRange(float64(sMilligram), sodiumLevels)
	}
	return getPointsFromRange(float64(sMilligram), sodiumLevels)
}

func (fPercent FruitsPercent) GetPoints(scoreType ScoreType) int {
	if scoreType == Beverage {
		if fPercent > 80 {
			return 10
		} else if fPercent > 60 {
			return 4
		} else if fPercent > 40 {
			return 2
		}

		return 0
	}

	if fPercent > 80 {
		return 5
	} else if fPercent > 60 {
		return 2
	} else if fPercent > 40 {
		return 1
	}

	return 0
}

func (fGram FibreGram) GetPoints(scoreType ScoreType) int {
	if scoreType == Beverage {
		return getPointsFromRange(float64(fGram), fibreLevels)
	}
	return getPointsFromRange(float64(fGram), fibreLevels)
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

func getPointsFromRange(v float64, steps []float64) int {
	length := len(steps)
	for i, l := range steps {
		if v > l {
			return length - i
		}
	}

	return 0
}

func (ns NutritionalScore) GetNutriScore() string {

}
