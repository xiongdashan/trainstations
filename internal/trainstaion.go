package internal

// @bjb|北京北|VAP|beijingbei|bjb|0
type TrainStation struct {
	PinYinShortName string
	Name            string
	Code            string
	PinYinFull      string
	City            string
	CityCode        string
	Province        string
}

func (t *TrainStation) FromStringArray(source []string) {
	if len(source) == 0 {
		return
	}
	t.PinYinShortName = t.getPropertyVal(source, 0)
	t.Name = t.getPropertyVal(source, 1)
	t.Code = t.getPropertyVal(source, 2)
	t.PinYinFull = t.getPropertyVal(source, 3)
}

func (t *TrainStation) getPropertyVal(source []string, postion int) string {
	if len(source) < postion {
		return ""
	}
	return source[postion]
}
