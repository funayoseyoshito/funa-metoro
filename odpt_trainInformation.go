package metro

import (
	"fmt"
	"time"
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

//func (t *TrainInformations) Parse() ApiResponse {
//	return t
//}

//func (t *TrainInformations) SetTest() {
//	for k, v := range *t {
//		if k == 0 {
//			v.Railway = "OKOKOKOK"
//		}
//	}
//}

//func (t *TrainInformations) Yoshito() {
//	fmt.Println("funayose yoshito")
//}

func (t *TrainInformations) Dump() {

	fmt.Printf("%T\n", t)

	fmt.Println("11111111111111111111")
	//fmt.Println(t)

	for _, v := range *t {
		//v = TrainInfomation{}
		//(&v).Railway = "ああああああああああああああああああああああああ"
		fmt.Println("---------")
		fmt.Println(v.Railway)
		//t[i].Railway = "testtestetesttestestsetstest"
		fmt.Println("---------")
	}

	fmt.Println(t)
}

type TrainInformations []*TrainInfomation

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
