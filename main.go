package main

import (
	"fmt"

	"github.com/xiongdashan/trainstationsname/internal"
)

var geoStationCity *internal.GeoStationCity

const geoKey = "你的高德Key"

func main() {
	jsSoruce := new(internal.JsSource)
	jsSoruce.GetAllStations()
	geoStationCity = internal.NewGeoStationCity(geoKey)
	for _, station := range jsSoruce.Stations {
		fmt.Println(station.Name)
		setSationCity(station)
	}

	saver := internal.NewSaveJsFile("./data/train_stations.js")
	saver.Save(jsSoruce.Stations)

	fmt.Printf("共%d条\n", len(jsSoruce.Stations))
}

func setSationCity(station *internal.TrainStation) {
	geoCity := geoStationCity.GetCity(station.Name)
	if geoCity == nil || geoCity.Status == "0" {
		return
	}
	codes := geoCity.Geocodes[0]
	station.City = codes.City
	station.CityCode = codes.Citycode
	station.Province = codes.Province
}
