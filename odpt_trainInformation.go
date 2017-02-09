package metro

import (
	"fmt"
	"time"
)

type TrainInfomation struct {
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

func (t *TrainInformations) Dump() {
	fmt.Println("11111111111111111111")
	//fmt.Println(t)
	for k, v := range *t {
		fmt.Println(k, v)
	}
}

type TrainInformations []TrainInfomation

func fetchTrainName(railway string) string {
	name := map[string]string{
		"Ginza":      "銀座線",
		"Marunouchi": "丸の内線",
		"Chiyoda":    "千代田線",
		"Hibiya":     "日比谷線",
		"Namboku":    "南北線",
		"Yurakucho":  "有楽町線",
		"Fukutoshin": "副都心線",
		"Hanzomon":   "半蔵門線",
		"Tozai":      "東西線",
	}
	return name[railway]
}

func (m *Metro) GetODPTTrainInformation() *Metro {
	m.apiPath = "datapoints"
	m.params.rdfType = "odpt:TrainInformation"
	m.response = &TrainInformations{}
	return m
}