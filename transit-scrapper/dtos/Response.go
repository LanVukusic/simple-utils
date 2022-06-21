package dtos

// Response is the response from Nomago API
type Response struct {
	Departures []struct {
		Planer       int64   `json:"planer"`
		FromCode     int     `json:"from_code"`
		ToCode       int     `json:"to_code"`
		TimeFrom     string  `json:"time_from"`
		TimeTo       string  `json:"time_to"`
		Departure    string  `json:"departure"`
		RideTime     string  `json:"ride_time"`
		ArrivedAt    string  `json:"arrived_at"`
		Price        float64 `json:"price"`
		PriceAvarage float64 `json:"priceAvarage"`
		AdultPrice   float64 `json:"adultPrice"`
		ConvertPrice float64 `json:"convert_price"`
		Km           int     `json:"km"`
		Stops        int     `json:"stops"`
		Products     []struct {
			ProdID   int     `json:"prod_id"`
			Naziv    string  `json:"naziv"`
			Price    float64 `json:"price"`
			Notax    int     `json:"notax"`
			Tax      float64 `json:"tax"`
			FromCode int     `json:"from_code"`
			ToCode   int     `json:"to_code"`
			DestDate string  `json:"dest_date"`
			Quantity int     `json:"quantity"`
		} `json:"products"`
		Stations []struct {
			Zap            int64  `json:"zap"`
			ZacetnaPostaja string `json:"zacetnaPostaja"`
			KoncnaPostaja  string `json:"koncnaPostaja"`
			Odhod          string `json:"odhod"`
			Prihod         string `json:"prihod"`
			CasVoznje      int    `json:"casVoznje"`
			Cakanje        any    `json:"cakanje,omitempty"`
			InStations     []struct {
				Zap     int    `json:"zap"`
				Postaja string `json:"postaja"`
				Prihod  string `json:"prihod"`
				Odhod   string `json:"odhod"`
			} `json:"in_stations"`
		} `json:"stations"`
	} `json:"departures"`
}
