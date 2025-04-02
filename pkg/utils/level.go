package utils

import (
	"marcel-games-backend/internal/constants"
	"math/rand"
	"slices"
)

func GetNextLevelCountryCodes(level int) []string {
	sorted := sortCountriesByArea()

	countryCount := getNumberOfCountries(level)
	countrySelectWindow := getCountrySelectWindow(level)
	availableCountries := sorted[:countryCount+countrySelectWindow]

	result := make([]string, 0, countryCount)
	indices := rand.Perm(len(availableCountries))[:countryCount]
	for _, idx := range indices {
		result = append(result, availableCountries[idx].Code)
	}

	return result
}

func sortCountriesByArea() []constants.Country {
	sorted := make([]constants.Country, len(constants.Countries))
	copy(sorted, constants.Countries)
	slices.SortFunc(sorted, func(i, j constants.Country) int {
		if i.Area > j.Area {
			return -1
		}
		if i.Area < j.Area {
			return 1
		}
		return 0
	})
	return sorted
}

func getCountrySelectWindow(level int) int {
	switch {
	case level <= 15:
		return 20
	case level <= 30:
		return 30
	case level <= 50:
		return 50
	case level <= 100:
		return 75
	case level <= 250:
		return 100
	case level <= 500:
		return 150
	case level <= 1000:
		return 200
	default:
		return 200
	}
}

func getNumberOfCountries(level int) int {
	switch {
	case level <= 15:
		return 1
	case level <= 30:
		return rand.Intn(3) + 1
	case level <= 50:
		return rand.Intn(3) + 2
	case level <= 100:
		return rand.Intn(4) + 2
	case level <= 250:
		return rand.Intn(6) + 5
	case level <= 500:
		return rand.Intn(8) + 8
	case level <= 1000:
		return rand.Intn(4) + 12
	default:
		return rand.Intn(6) + 15
	}
}
