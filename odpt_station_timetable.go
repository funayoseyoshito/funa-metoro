package metro

import (
	"fmt"
	"strings"
	"time"
)

type StationTimeTable struct {
	Context           string                 `json:"@context"`
	ID                string                 `json:"@id"`
	Type              string                 `json:"@type"`
	OwlSameAs         string                 `json:"owl:sameAs"`
	DcDate            time.Time              `json:"dc:date"`
	OdptStation       string                 `json:"odpt:station"`
	OdptRailway       string                 `json:"odpt:railway"`
	OdptOperator      string                 `json:"odpt:operator"`
	OdptRailDirection string                 `json:"odpt:railDirection"`
	OdptWeekdays      StationTimeTableObject `json:"odpt:weekdays"`
	OdptSaturdays     StationTimeTableObject `json:"odpt:saturdays"`
	OdptHolidays      StationTimeTableObject `json:"odpt:holidays"`
}

type StationTimeTableObject []struct {
	OdptDepartureTime      string `json:"odpt:departureTime"`
	OdptDestinationStation string `json:"odpt:destinationStation"`
	OdptTrainType          string `json:"odpt:trainType"`
	OdptIsLast             bool   `json:"odpt:isLast,omitempty"`
	OdptIsOrigin           bool   `json:"odpt:isOrigin"`
	OdptCarComposition     int    `json:"odpt:carComposition"`
	OdptNotes              string `json:"odpt:notes"`
}

type StationTimeTables []StationTimeTable

func (t StationTimeTable) UCODE() string {
	return strings.Replace(t.ID, "urn:ucode:_", "", -1)
}

func (t StationTimeTable) StationName() string {
	return strings.Replace(t.OdptStation, "odpt.Station:", "", -1)
}

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

func (t *StationTimeTables) Dump() {
	for _, v := range *t {
		fmt.Printf("\n%s", v)
	}
}

func (m *Metro) StationTimeTable() StationTimeTables {
	m.apiPath = "datapoints"
	m.SetParam("rdf:type", "odpt:StationTimetable")
	r := m.requet(&StationTimeTables{})
	t, _ := r.(*StationTimeTables)

	return *t
}
