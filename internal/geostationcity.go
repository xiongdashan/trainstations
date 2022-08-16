package internal

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type GeoStationCity struct {
	geoKey string
}

func NewGeoStationCity(geoKey string) *GeoStationCity {
	return &GeoStationCity{
		geoKey: geoKey,
	}
}

func (g *GeoStationCity) GetCity(name string) *GeoReponseData {
	return g.req(name)
}

func (g *GeoStationCity) req(name string) *GeoReponseData {
	if name == "" {
		return nil
	}
	reqUrl := fmt.Sprintf("https://restapi.amap.com/v3/geocode/geo?key=%s&address=%s", g.geoKey, name)
	client := resty.New()
	res, _ := client.R().Get(reqUrl)

	var result *GeoReponseData

	json.Unmarshal(res.Body(), &result)

	return result
}

type GeoReponseData struct {
	Status   string `json:"status"`
	Info     string `json:"info"`
	Infocode string `json:"infocode"`
	Count    string `json:"count"`
	Geocodes []struct {
		FormattedAddress string        `json:"formatted_address"`
		Country          string        `json:"country"`
		Province         string        `json:"province"`
		Citycode         string        `json:"citycode"`
		City             string        `json:"city"`
		District         string        `json:"district"`
		Township         []interface{} `json:"township"`
		Neighborhood     struct {
			Name []interface{} `json:"name"`
			Type []interface{} `json:"type"`
		} `json:"neighborhood"`
		Building struct {
			Name []interface{} `json:"name"`
			Type []interface{} `json:"type"`
		} `json:"building"`
		Adcode   string        `json:"adcode"`
		Street   []interface{} `json:"street"`
		Number   []interface{} `json:"number"`
		Location string        `json:"location"`
		Level    string        `json:"level"`
	} `json:"geocodes"`
}
