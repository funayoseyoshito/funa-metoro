package metro

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
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

func (t *TrainInfomation) dump() {
	fmt.Println(t)
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

func (m *Metro) Execute() {

	fmt.Printf("%T \n", apiResults)

	//panic("OKOKOKOK")

	url := m.getRequestUrl()
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("==========")
	fmt.Println(string(body))
	fmt.Println("==========")

	err = json.Unmarshal(body, &apiResults)
	if err != nil {
		fmt.Println("ここ！")
		log.Fatal(err)
	}

	//fmt.Println("=================11111")
	//fmt.Println(apiResults)
	//fmt.Println("=================11111")

	//fmt.Println("-----------------------------")
	//var r ApiResults = apiResults

	//for _, v := range apiResults {
	//	fmt.Println(v)
	//	fmt.Println(i, v)
	//fmt.Println("@@@@@@@")
	//(v).dump()
	//}

	s := reflect.ValueOf(apiResults)

	for i := 0; i < s.Len(); i++ {
		v := s.Index(i)
		fmt.Println(v)

		fmt.Printf("%T \n", v)
		//panic("OK")
	}

	fmt.Println("-----------------------------")

	//return &apiResults
}

func (m *Metro) GetODPTTrainInformation() *Metro {
	m.apiPath = "datapoints"
	m.params.rdfType = "odpt:TrainInformation"

	//fmt.Println(apiResults)
	//fmt.Println(apiResults.(string))
	i, _ := apiResults.(TrainInformations)
	apiResults = i
	//fmt.Println(i)
	//fmt.Printf("%T \n", i)
	//fmt.Println(v)
	//panic("OK")
	//panic("!!!!!!!!")
	//i, _ := apiResults.([](*TrainInfomation))
	//fmt.Println(i)

	//var apiResults []ApiResponse

	//fmt.Println([]TrainInfomation(&apiResults)

	//apiResults = []TrainInfomation{}
	//panic("OK")
	return m
}
