package bls

type FootNotes []struct {
	Code string `json:"code"`
	Text string `json:"text"`
}

type DataPoints []struct {
	Year       string    `json:"year"`
	Period     string    `json:"period"`
	PeriodName string    `json:"periodName"`
	Latest     string    `json:"latest"`
	Value      string    `json:"value"`
	Footnotes  FootNotes `json:"footnotes"`
}

type SeriesData struct {
	SeriesID string     `json:"seriesID"`
	Data     DataPoints `json:"data"`
}

type CpiResponse struct {
	Status       string     `json:"status"`
	ResponseTime int        `json:"responseTime"`
	Message      []string   `json:"message"`
	Results      SeriesData `json:"Results"`
}
