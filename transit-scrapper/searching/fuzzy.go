package searching

import (
	"transit-trackers/dtos"
)

// get best match from fuzzy search
func GetBestMatch(search string, options []dtos.Station) dtos.Station {
	// var BestStation dtos.Station
	// var rank int = 0

	// // loop over options
	// for _, station := range options {
	// 	// check if search is in station name
	// 	var rankFuzzy = fuzzy.RankMatch(search, station.Name)
	// 	if rankFuzzy > rank {
	// 		BestStation = station
	// 		rank = rankFuzzy
	// 	}
	// }

	return options[0]
}
