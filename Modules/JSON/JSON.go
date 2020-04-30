package JSON

import (
	"encoding/json"
)

// This module contains functions to parse and give demanded JSON files

// *************** Type Declarations *****************

type WorldCases struct {
	Results []struct {
		TotalCases             int `json:"total_cases"`
		TotalRecovered         int `json:"total_recovered"`
		TotalUnresolved        int `json:"total_unresolved"`
		TotalDeaths            int `json:"total_deaths"`
		TotalNewCasesToday     int `json:"total_new_cases_today"`
		TotalNewDeathsToday    int `json:"total_new_deaths_today"`
		TotalActiveCases       int `json:"total_active_cases"`
		TotalSeriousCases      int `json:"total_serious_cases"`
		TotalAffectedCountries int `json:"total_affected_countries"`
		Source                 struct {
			URL string `json:"url"`
		} `json:"source"`
	} `json:"results"`
	Stat string `json:"stat"`
}

type CountryCases struct {
	Countrydata []struct {
		Info struct {
			Ourid  int    `json:"ourid"`
			Title  string `json:"title"`
			Code   string `json:"code"`
			Source string `json:"source"`
		} `json:"info"`
		TotalCases          int `json:"total_cases"`
		TotalRecovered      int `json:"total_recovered"`
		TotalUnresolved     int `json:"total_unresolved"`
		TotalDeaths         int `json:"total_deaths"`
		TotalNewCasesToday  int `json:"total_new_cases_today"`
		TotalNewDeathsToday int `json:"total_new_deaths_today"`
		TotalActiveCases    int `json:"total_active_cases"`
		TotalSeriousCases   int `json:"total_serious_cases"`
		TotalDangerRank     int `json:"total_danger_rank"`
	} `json:"countrydata"`
	Stat string `json:"stat"`
}

type CountryCalculatedValues struct {

	// Data calculated from the grabbed data with CountryCases object
	totalClosed int // deaths + recovered // Number of closed cases in a country
	//active                    int // confirmed - closed // Number of active cases in a country
	totalDeathPercentage        int // deaths/closed      // Percentage of death in a country
	totalRecoveryPercentage     int // recovered/closed   // Percentage of recovery in a country
	totalActivePercentage       int // active/confirmed   // Percentage of active cases in a country
	totalClosedPercentage       int // closed/confirmed   // Percentage of closed cases in a country
	totalDeficit                int // totalNewCasesToday - totalNewDeathsToday // Value that show the change in trends
	totalSeriousCasesPercentage int // totalSeriousCases/totalActiveCases       // Percentage of Serious cases "
	totalActiveCasesPercentage  int // totalActiveCases/totalCases              // Percentage of active cases "
}

type WorldCalculatedValues struct {

	// Data calculated from the grabbed data with WorldCases object
	totalRecoveredPercentage    int // totalRecovered/totalCases                // Percentage of recovery in the World
	totalUnresolvedPercentage   int // totalUnresolved/totalCases               // Percentage of death "
	totalClosed                 int // totalCases - activeCases                 // Number of closed cases "
	totalDeathsPercentage       int // totalDeaths/totalClosed                  // Percentage of deaths over closed "
	totalActivePercentage       int // active/confirmed                         // Percentage of active cases "
	totalClosedPercentage       int // closed/confirmed                         // Percentage of closed cases "
	totalDeficit                int // totalNewCasesToday - totalNewDeathsToday // Value that show the change in trends
	totalSeriousCasesPercentage int // totalSeriousCases/totalActiveCases       // Percentage of Serious cases "

}

// *************** Function Declarations *******************

// ######## Endpoint getter function ########

func EndPointGet(Name string, body []byte) interface{} { // Gets the end point values and

	if Name == "World" || Name == "WorldCases" {
		object := WorldCases{}
		json.Unmarshal(body, &object.Results[0].TotalCases)
		json.Unmarshal(body, &object.Results[0].TotalRecovered)
		json.Unmarshal(body, &object.Results[0].TotalUnresolved)
		json.Unmarshal(body, &object.Results[0].TotalDeaths)
		json.Unmarshal(body, &object.Results[0].TotalNewCasesToday)
		json.Unmarshal(body, &object.Results[0].TotalNewDeathsToday)
		json.Unmarshal(body, &object.Results[0].TotalActiveCases)
		json.Unmarshal(body, &object.Results[0].TotalSeriousCases)
		json.Unmarshal(body, &object.Results[0].TotalAffectedCountries)
		return object
	} else if Name == "Country" || Name == "CountryCases" {
		object := CountryCases{}
		json.Unmarshal(body, &object.Countrydata[0].Info.Title)
		json.Unmarshal(body, &object.Countrydata[0].TotalCases)
		json.Unmarshal(body, &object.Countrydata[0].TotalRecovered)
		json.Unmarshal(body, &object.Countrydata[0].TotalUnresolved)
		json.Unmarshal(body, &object.Countrydata[0].TotalDeaths)
		json.Unmarshal(body, &object.Countrydata[0].TotalNewCasesToday)
		json.Unmarshal(body, &object.Countrydata[0].TotalNewDeathsToday)
		json.Unmarshal(body, &object.Countrydata[0].TotalActiveCases)
		json.Unmarshal(body, &object.Countrydata[0].TotalSeriousCases)
		json.Unmarshal(body, &object.Countrydata[0].TotalDangerRank)

		return object
	} else {
		return nil
	}
}

// ######## Information handling function ########

func CountryCalculator(country CountryCases) CountryCalculatedValues { // Calculates the given values by the API for
	// each country
	totalClosed := country.Countrydata[0].TotalRecovered + country.Countrydata[0].TotalCases
	totalDeathPercentage := (country.Countrydata[0].TotalDeaths / totalClosed) * 100
	totalRecoveryPercentage := (country.Countrydata[0].TotalRecovered / totalClosed) * 100
	totalActivePercentage := (country.Countrydata[0].TotalActiveCases / country.Countrydata[0].TotalCases) * 100
	totalClosedPercentage := (totalClosed / country.Countrydata[0].TotalCases) * 100
	totalDeficit := country.Countrydata[0].TotalNewCasesToday - country.Countrydata[0].TotalNewDeathsToday
	totalSeriousCasesPercentage := country.Countrydata[0].TotalSeriousCases / country.Countrydata[0].TotalActiveCases
	totalActiveCasesPercentage := country.Countrydata[0].TotalActiveCases / country.Countrydata[0].TotalCases

	return CountryCalculatedValues{totalClosed, totalDeathPercentage,
		totalRecoveryPercentage, totalActivePercentage,
		totalClosedPercentage, totalDeficit,
		totalSeriousCasesPercentage, totalActiveCasesPercentage}
}

func WorldCalculator(world WorldCases) WorldCalculatedValues { //Calculates the given values by the API for the world

	totalRecoveredPercentage := (world.Results[0].TotalRecovered / world.Results[0].TotalCases) * 100
	totalUnresolvedPercentage := (world.Results[0].TotalUnresolved / world.Results[0].TotalCases) * 100
	totalClosed := world.Results[0].TotalCases - world.Results[0].TotalUnresolved
	totalDeathsPercentage := (world.Results[0].TotalDeaths / totalClosed) * 100
	totalActivePercentage := (world.Results[0].TotalActiveCases / world.Results[0].TotalCases) * 100
	totalClosedPercentage := (totalClosed / world.Results[0].TotalCases) * 100
	totalDeficit := world.Results[1].TotalNewCasesToday - world.Results[1].TotalNewDeathsToday
	totalSeriousCasesPercentage := (world.Results[1].TotalSeriousCases / world.Results[1].TotalActiveCases) * 100

	return WorldCalculatedValues{totalRecoveredPercentage,
		totalUnresolvedPercentage, totalClosed,
		totalDeathsPercentage, totalActivePercentage,
		totalClosedPercentage, totalDeficit,
		totalSeriousCasesPercentage}
}
