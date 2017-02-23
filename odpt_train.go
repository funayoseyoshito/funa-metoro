package metro

import (
	"fmt"
	"strings"
	"time"
)

type Train struct {
	Context             string          `json:"@context"`
	Type                string          `json:"@type"`
	ID                  string          `json:"@id"`
	DcDate              time.Time       `json:"dc:date"`
	DctValid            time.Time       `json:"dct:valid"`
	OdptFrequency       int             `json:"odpt:frequency"`
	OdptRailway         RailWayID       `json:"odpt:railway"`
	OwlSameAs           string          `json:"owl:sameAs"`
	OdptTrainNumber     string          `json:"odpt:trainNumber"`
	OdptTrainType       TrainTypeID     `json:"odpt:trainType"`
	OdptDelay           int             `json:"odpt:delay"`
	OdptStartingStation string          `json:"odpt:startingStation"`
	OdptTerminalStation string          `json:"odpt:terminalStation"`
	OdptFromStation     string          `json:"odpt:fromStation"`
	OdptToStation       interface{}     `json:"odpt:toStation"`
	OdptRailDirection   RailDirectionID `json:"odpt:railDirection"`
	OdptTrainOwner      TrainOwnerID    `json:"odpt:trainOwner"`
}

type Trains []Train

func (t *Trains) Dump() {
	for _, v := range *t {
		fmt.Println(v)
	}
}

func (t Train) UCODE() string {
	return strings.Replace(t.ID, "urn:ucode:_", "", -1)
}

//func (t TrainInformations) Dump() {
//	for _, v := range t {
//		fmt.Printf("\n%s", v)
//	}
//}

func (m *Metro) ODPTTrain() Trains {
	m.apiPath = "datapoints"
	m.SetParam("rdf:type", "odpt:Train")

	if !m.isSetParam("odpt:railway") {
		panic("odpt:railway is required")
	}

	r := m.requet(&Trains{})
	t, _ := r.(*Trains)
	return *t
}
