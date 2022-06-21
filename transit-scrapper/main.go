package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"transit-trackers/dtos"
	"transit-trackers/nomagoParser"
	"transit-trackers/searching"
	"transit-trackers/views"

	tea "github.com/charmbracelet/bubbletea"
)

var from = flag.String("from", "", "Name of the station you depart from")
var to = flag.String("to", "", "Name of the station you arrive at")
var baseUrl string = "https://www.nomago.si/avtobusne-vozovnice/vozni-red"

func main() {
	flag.Parse()
	// fmt.Printf("from: %s - %s \n", *from, *to)

	model := views.InitialModel()
	p := tea.NewProgram(model)
	go func() {
		if err := p.Start(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	var html io.ReadCloser = nomagoParser.GetHtml(baseUrl)
	var crsf string = nomagoParser.GetCRSF(html)

	var froms = nomagoParser.GetBusStations(*from)
	var tos = nomagoParser.GetBusStations(*to)

	var fromStation = searching.GetBestMatch(*from, froms)
	var toStation = searching.GetBestMatch(*to, tos)

	alert := views.LoadingItem{}
	alert.Msg = fmt.Sprintf("%s - %s", fromStation.Name, toStation.Name)
	p.Send(alert)

	var departures dtos.Response = nomagoParser.GetDepartures(fromStation.ID, toStation.ID, nomagoParser.GetCurrentDate(), crsf)
	items := []views.Item{}

	if len(departures.Departures) == 0 {
		exitAlert := views.ExitAlert{}
		exitAlert.Msg = "No departures found"
		p.Send(exitAlert)
	}

	for _, departure := range departures.Departures {
		items = append(items, views.Item(departure.TimeFrom))
	}

	p.Send(items)

	for !model.Quitting {
	}
	fmt.Println("Quitting")
}
