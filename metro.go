package metro

import (
	"fmt"
	"strconv"
	//"io/ioutil"
	"log"
	//"net/http"
	"encoding/json"
	"net/url"
	//"net/http"
	//"io/ioutil"
	"io/ioutil"
	"path/filepath"
)

const APIBaseUrl = "https://api.tokyometroapp.jp/api/v2"

type Metro struct {
	apiPath  string
	token    string
	response APIResponser
	param    Param
}

type APIResponser interface {
	Dump()
}

type Params struct {
	rdfType                    string `param:"rdf:type"`
	consumerKey                string `param:"acl:consumerKey"`
	ID                         string `param:"@id"`
	ODPTOperator               string `param:"odpt:operator"`
	ODPTRailway                string `param:"odpt:railway"`
	ODPTTrainInformationStatus string `param:"odpt:trainInformationStatus"`
	ODPTTrainInformationText   string `param:"odpt:trainInformationText"`
	ODPStation                 string `param:"odpt:station"`
}

type Param map[string]interface{}

//NewMetro return New Metro struct
func NewMetro(token string) *Metro {
	m := &Metro{}
	m.param = make(Param, 5)
	m.SetParam("acl:consumerKey", token)
	return m
}

func (m *Metro) SetParam(k string, v interface{}) *Metro {
	m.param[k] = v
	return m
}

func (m *Metro) getRequestURI(p Param) string {
	u, err := url.Parse(APIBaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	u.Path += "/" + m.apiPath
	q := u.Query()

	for k, v := range p {
		var getStringValue string
		switch x := v.(type) {
		case int:
			getStringValue = strconv.Itoa(x)
		case string:
			getStringValue = x
		case bool:
			getStringValue = strconv.FormatBool(x)
		default:
			panic("param must be string or int")
		}
		q.Add(k, getStringValue)
	}

	u.RawQuery = q.Encode()
	return u.String()
}

// request
func (m *Metro) requet(t APIResponser) APIResponser {

	url := m.getRequestURI(m.param)
	fmt.Println(url)

	//TODO json from file
	fileName := "station_timetable.json"
	//fileName := "train_information.json"
	body, _ := ioutil.ReadFile(filepath.Join("example", fileName))

	//res, err := http.Get(url)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//body, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}

	err := json.Unmarshal(body, &t)
	if err != nil {

		log.Fatal(err)
	}

	return t
}

func getODPTOperatorName(id string) string {
	Operator := map[string]string{
		"odpt.Operator:TokyoMetro": "東京メトロ",
	}
	return Operator[id]
}

func getODPTRailWayName(id string) string {
	RailWay := map[string]string{
		"odpt.Railway:TokyoMetro.Ginza":                  "銀座線",
		"odpt.Railway:TokyoMetro.Marunouchi":             "丸ノ内線",
		"odpt.Railway:TokyoMetro.Hibiya":                 "日比谷線",
		"odpt.Railway:TokyoMetro.Tozai":                  "東西線",
		"odpt.Railway:TokyoMetro.Chiyoda":                "千代田線",
		"odpt.Railway:TokyoMetro.Yurakucho":              "有楽町線",
		"odpt.Railway:TokyoMetro.Hanzomon":               "半蔵門線",
		"odpt.Railway:TokyoMetro.Namboku":                "南北線",
		"odpt.Railway:TokyoMetro.Fukutoshin":             "副都心線",
		"odpt.Railway:JR-East":                           "JR線",
		"odpt.Railway:JR-East.Chuo":                      "中央線",
		"odpt.Railway:JR-East.ChuoKaisoku":               "中央線快速　",
		"odpt.Railway:JR-East.ChuoSobu":                  "中央・総武線各駅停車",
		"odpt.Railway:JR-East.Joban":                     "常磐線",
		"odpt.Railway:JR-East.KeihinTohoku":              "京浜東北線",
		"odpt.Railway:JR-East.Keiyo":                     "京葉線",
		"odpt.Railway:JR-East.Musashino":                 "武蔵野線",
		"odpt.Railway:JR-East.NaritaExpress":             "成田エクスプレス",
		"odpt.Railway:JR-East.Saikyo":                    "埼京線",
		"odpt.Railway:JR-East.ShonanShinjuku":            "湘南新宿ライン",
		"odpt.Railway:JR-East.Sobu":                      "総武線",
		"odpt.Railway:JR-East.Takasaki":                  "高崎線",
		"odpt.Railway:JR-East.Tokaido":                   "東海道線",
		"odpt.Railway:JR-East.Utsunomiya":                "宇都宮線",
		"odpt.Railway:JR-East.Yamanote":                  "山手線",
		"odpt.Railway:JR-East.Yokosuka":                  "横須賀線",
		"odpt.Railway:Keio.Inokashira":                   "井の頭線",
		"odpt.Railway:Keio.Keio":                         "京王線",
		"odpt.Railway:Keio.New":                          "京王新線",
		"odpt.Railway:Keisei.KeiseiMain":                 "京成本線",
		"odpt.Railway:Keisei.KeiseiOshiage":              "押上線",
		"odpt.Railway:MIR.TX":                            "つくばエクスプレス線",
		"odpt.Railway:Odakyu.Odawara":                    "小田原線",
		"odpt.Railway:SaitamaRailway.SaitamaRailway":     "埼玉高速鉄道線",
		"odpt.Railway:Seibu.Ikebukuro":                   "池袋線",
		"odpt.Railway:Seibu.SeibuYurakucho":              "西武有楽町線",
		"odpt.Railway:Seibu.Shinjuku":                    "新宿線",
		"odpt.Railway:TWR.Rinkai":                        "りんかい線",
		"odpt.Railway:Tobu.Isesaki":                      "伊勢崎線",
		"odpt.Railway:Tobu.Tojo":                         "東上線",
		"odpt.Railway:Toei.Asakusa":                      "浅草線",
		"odpt.Railway:Toei.Mita":                         "三田線",
		"odpt.Railway:Toei.NipporiToneri":                "日暮里・舎人ライナー",
		"odpt.Railway:Toei.Oedo":                         "大江戸線",
		"odpt.Railway:Toei.Shinjuku":                     "新宿線",
		"odpt.Railway:Toei.TodenArakawa":                 "都電荒川線",
		"odpt.Railway:Tokyu.DenEnToshi":                  "田園都市線",
		"odpt.Railway:Tokyu.Meguro":                      "目黒線",
		"odpt.Railway:Tokyu.Toyoko":                      "東横線",
		"odpt.Railway:ToyoRapidRailway.ToyoRapidRailway": "東葉高速線",
		"odpt.Railway:Yurikamome.Yurikamome":             "ゆりかもめ",
	}
	return RailWay[id]
}
