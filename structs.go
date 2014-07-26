package goefa

type DateTimeType struct {
	Date struct {
		Day   int `xml:"day,attr"`
		Month int `xml:"month,attr"`
		Year  int `xml:"year,attr"`
	} `xml:"itdDate"`

	Time struct {
		Hour   int `xml:"hour,attr"`
		Minute int `xml:"minute,attr"`
	} `xml:"itdTime"`
}

type Line struct {
	Number    string `xml:"number,attr"`
	Direction string `xml:"direction,attr"`
}

type Departure struct {
	Countdown int    `xml:"countdown,attr"`
	Platform  string `xml:"platform,attr"`

	DateTime    DateTimeType `xml:"itdDateTime"`
	ServingLine Line         `xml:"itdServingLine"`
}

type StopInfo struct {
	State string `xml:"state,attr"`

	IdfdStop struct {
		StopName  string `xml:",chardata"`
		MatchQlty int    `xml:"matchQuality,attr"`
		StopID    int    `xml:"stopID,attr"`
	} `xml:"odvNameElem"`
}

type StopResult struct {
	Stop       StopInfo    `xml:"itdDepartureMonitorRequest>itdOdv>itdOdvName"`
	Departures []Departure `xml:"itdDepartureMonitorRequest>itdDepartureList>itdDeparture"`
}
