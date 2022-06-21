package dtos

// Request is the request to Nomago API
type Request struct {
	From        string `json:"dest_from"`
	To          string `json:"dest_to"`
	BusDateFrom string `json:"bus_date_from"`
	Token       string `json:"_token"`
}
