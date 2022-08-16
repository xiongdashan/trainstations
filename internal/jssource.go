package internal

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type JsSource struct {
	Stations []*TrainStation
}

func (j *JsSource) GetAllStations() {
	nameUrl := "https://kyfw.12306.cn/otn/resources/js/framework/station_name.js"
	client := resty.New()
	res, err := client.R().Get(nameUrl)
	checkErr(err)
	stationsName := string(res.Body())
	stationsName = j.removeJsChar(stationsName)
	fmt.Println(stationsName)
	j.splitStationsJson(stationsName)
}

func (j *JsSource) splitStationsJson(stationsJson string) {
	stationItems := strings.Split(stationsJson, "@")
	for _, station := range stationItems {
		if station == "" {
			continue
		}
		j.setItemToResult(station)
	}
}

func (JsSource) removeJsChar(stationJson string) string {
	result := strings.ReplaceAll(stationJson, "var station_names ='", "")
	result = strings.ReplaceAll(result, "';", "")
	return result
}

// @bjb|北京北|VAP|beijingbei|bjb|0
func (j *JsSource) setItemToResult(stationJson string) {
	stationPro := strings.Split(stationJson, "|")
	result := new(TrainStation)
	result.FromStringArray(stationPro)
	j.Stations = append(j.Stations, result)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
