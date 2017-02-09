package metro

import (
	"fmt"
	"time"
)

type TrainInformations2 []TrainInfomation2

type TrainInfomation2 struct {
	Context                string    `json:"@context"`
	Id                     string    `json:"@id"`
	Type                   string    `json:"@type"`
	Date                   time.Time `json:"dc:date"`
	Valid                  time.Time `json:"dct:valid"`
	Operator               string    `json:"odpt:operator"`
	TimeOfOrigin           time.Time `json:"odpt:timeOfOrigin"`
	Railway                string    `json:"odpt:railway"`
	TrainInformationStatus string    `json:"odpt:trainInformationStatus"`
	TrainInformationText   string    `json:"odpt:trainInformationText"`
}

func (t *TrainInformations2) Dump() {
	fmt.Println("2222222222222222222")
	for k, v := range *t {
		fmt.Println(k, v)
	}
}

func (m *Metro) GetODPTTrainInformation2() *Metro {
	m.apiPath = "datapoints"
	m.params.rdfType = "odpt:TrainInformation"
	m.response = &TrainInformations2{}
	return m
}
