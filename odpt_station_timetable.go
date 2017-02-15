package metro

import (
	"fmt"
	"time"
)

type StationTimetable []struct {
	ID                string    `json:"@id"`
	Type              string    `json:"@type"`
	Context           string    `json:"@context"`
	OwlSameAs         string    `json:"owl:sameAs"`
	OdptStation       string    `json:"odpt:station"`
	OdptRailway       string    `json:"odpt:railway"`
	OdptOperator      string    `json:"odpt:operator"`
	OdptRailDirection string    `json:"odpt:railDirection"`
	DcDate            time.Time `json:"dc:date"`
	OdptWeekdays      []struct {
		OdptDepartureTime      string `json:"odpt:departureTime"`
		OdptDestinationStation string `json:"odpt:destinationStation"`
		OdptTrainType          string `json:"odpt:trainType"`
		OdptIsLast             bool   `json:"odpt:isLast,omitempty"`
	} `json:"odpt:weekdays"`
	OdptSaturdays []struct {
		OdptDepartureTime      string `json:"odpt:departureTime"`
		OdptDestinationStation string `json:"odpt:destinationStation"`
		OdptTrainType          string `json:"odpt:trainType"`
		OdptIsLast             bool   `json:"odpt:isLast,omitempty"`
	} `json:"odpt:saturdays"`
	OdptHolidays []struct {
		OdptDepartureTime      string `json:"odpt:departureTime"`
		OdptDestinationStation string `json:"odpt:destinationStation"`
		OdptTrainType          string `json:"odpt:trainType"`
		OdptIsLast             bool   `json:"odpt:isLast,omitempty"`
	} `json:"odpt:holidays"`
}

//func (t TrainInfomation) UCODE() string {
//	return strings.Replace(t.ID, "urn:ucode:_", "", -1)
//}
//
//func (t TrainInfomation) OperatorName() string {
//	return getODPTOperatorName(t.Operator)
//}
//
//func (t TrainInfomation) RailsWayName() string {
//	return getODPTRailWayName(t.Railway)
//}
//
//

func (t *StationTimetable) Dump() {
	for _, v := range *t {
		fmt.Printf("\n%s", v)
	}
}

//
//func (m *Metro) GetStationTimeTable() *TrainInformations {
//	return m.GetODPTTrainInformationWithParam()
//}
//
////
func (m *Metro) GetStationTimeTableWithParam(p *Params) StationTimetable {

	m.apiPath = "datapoints"
	p.rdfType = "odpt:StationTimetable"
	r := m.requet(&StationTimetable{}, p)
	t, _ := r.(*StationTimetable)
	return *t
}
