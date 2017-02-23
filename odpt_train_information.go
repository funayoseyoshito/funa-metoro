package metro

import (
	"fmt"
	"strings"
	"time"
)

type TrainInfomation struct {
	Context                string    `json:"@context"`
	ID                     string    `json:"@id"`
	Type                   string    `json:"@type"`
	Date                   time.Time `json:"dc:date"`
	Valid                  time.Time `json:"dct:valid"`
	Operator               string    `json:"odpt:operator"`
	TimeOfOrigin           time.Time `json:"odpt:timeOfOrigin"`
	Railway                string    `json:"odpt:railway"`
	TrainInformationStatus string    `json:"odpt:trainInformationStatus"`
	TrainInformationText   string    `json:"odpt:trainInformationText"`
}

func (t TrainInfomation) UCODE() string {
	return strings.Replace(t.ID, "urn:ucode:_", "", -1)
}

func (t TrainInfomation) OperatorName() string {
	return getODPTOperatorName(t.Operator)
}

func (t TrainInfomation) RailsWayName() string {
	return getODPTRailWayName(t.Railway)
}

type TrainInformations []TrainInfomation

func (t TrainInformations) Dump() {
	for _, v := range t {
		fmt.Printf("\n%s", v)
	}
}

func (m *Metro) ODPTTrainInformation() TrainInformations {
	m.apiPath = "datapoints"
	m.SetParam("rdf:type", "odpt:TrainInformation")
	r := m.requet(&TrainInformations{})
	t, _ := r.(*TrainInformations)
	return *t
}
