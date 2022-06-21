package nomagoParser

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
	"transit-trackers/dtos"

	"github.com/PuerkitoBio/goquery"
)

// get current time in epoch
func GetCurrentTime() int64 {
	return time.Now().Unix()
}

// get current date in format DD-MM-YYYY
func GetCurrentDate() string {
	var t = time.Now()
	return fmt.Sprintf("%d.%d.%d", t.Day()+1, t.Month(), t.Year())
}

// get html string from url
func GetHtml(url string) io.ReadCloser {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	return resp.Body
}

// parse the html body and return the CRSF token or "" if not found
func GetCRSF(html io.ReadCloser) string {
	doc, err := goquery.NewDocumentFromReader(html)
	if err != nil {
		log.Fatal(err)
	}

	var out string = ""
	docs := doc.Find("meta")
	docs.Each(func(i int, s *goquery.Selection) {
		if s.AttrOr("name", "") == "csrf-token" {
			out = s.AttrOr("content", "")
		}
	})
	return out
}

// request all bus stations search
func GetBusStations(stationName string) []dtos.Station {

	var url string = fmt.Sprintf("https://vozovnice.nomago.si/api/v1/destinations/search?q=%s&_=%d", url.QueryEscape(stationName), GetCurrentTime())

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	// read body
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	// convert body to json
	var res = []dtos.Station{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal(err)
	}

	return res
}

func GetDepartures(from string, to string, date string, token string) dtos.Response {
	if token == "" {
		log.Fatal("no valid token received")
	}

	var baseUrl string = "https://www.nomago.si/api/v1/departures"
	var url string = fmt.Sprintf("%s?dest_from=%s&dest_to=%s&bus_date_from=%s&_token=%s", baseUrl, from, to, date, token)

	resp, err := http.Post(url, "application/json", nil)

	// read body
	body, err := io.ReadAll(resp.Body)

	defer resp.Body.Close()
	// convert body to json
	var res = dtos.Response{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Cant unmarshal departures body ", err)
	}

	return res
}
