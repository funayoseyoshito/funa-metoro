package metro

import (
	"fmt"
	"time"
	"strings"
)

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

func (t TrainInfomation) getID () string {
	return strings.Replace(t.ID, "urn:ucode:_", "", -1)
}



type TrainInformations []TrainInfomation

func (t *TrainInformations) Dump() {

	for _, v := range *t {
		fmt.Printf("\n%s\n",v)
		//fmt.Println(v.getID())
		fmt.Println(v.Date)
	}
}


func (m *Metro) GetODPTTrainInformation() *TrainInformations {
	return m.GetODPTTrainInformationWithParam(params{})
}

func (m *Metro) GetODPTTrainInformationWithParam(p params) *TrainInformations {
	m.apiPath = "datapoints"
	p.rdfType = "odpt:TrainInformation"
	r := m.requet(&TrainInformations{}, p)
	t ,_ := r.(*TrainInformations)
	return t
}
