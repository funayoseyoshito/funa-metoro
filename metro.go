package metro

import (
	"fmt"
	//"io/ioutil"
	"log"
	//"net/http"
	"encoding/json"
	"net/url"
	"reflect"
	//"net/http"
	//"io/ioutil"
)

const APIBaseUrl = "https://api.tokyometroapp.jp/api/v2"

type Metro struct {
	apiPath  string
	token    string
	response APIResponser
}

type Params struct {
	rdfType                    string `param:"rdf:type"`
	consumerKey                string `param:"acl:consumerKey"`
	ID                         string `param:"@id"`
	ODPTOperator               string `param:"odpt:operator"`
	ODPTRailway                string `param:"odpt:railway"`
	ODPTTrainInformationStatus string `param:"odpt:trainInformationStatus"`
	ODPTTrainInformationText   string `param:"odpt:trainInformationText"`
}

type APIResponser interface {
	Dump()
}

//NewMetro return New Metro struct
func NewMetro(token string) *Metro {
	m := &Metro{}
	m.token = token
	return m
}

func (m *Metro) getRequestURI(p Params) string {
	p.consumerKey = m.token
	u, err := url.Parse(APIBaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	u.Path += "/" + m.apiPath
	q := u.Query()

	rt, rv := reflect.TypeOf(p), reflect.ValueOf(p)
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		k := field.Tag.Get("param")
		if v := rv.Field(i).String(); v != "" {
			q.Add(k, v)
		}
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func (m *Params) set() {
	fmt.Println("hello")
}

func (m *Metro) requet(t APIResponser, p Params) APIResponser {

	url := m.getRequestURI(p)
	fmt.Println(url)
	//res, err := http.Get(url)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//body, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//body := []byte(`[{"@context":"http://vocab.tokyometroapp.jp/context_odpt_TrainInformation.json","@id":"urn:ucode:_00001C000000000000010000030C3BE3","dc:date":"2017-02-09T22:55:03+09:00","dct:valid":"2017-02-09T23:00:03+09:00","odpt:operator":"odpt.Operator:TokyoMetro","odpt:railway":"odpt.Railway:TokyoMetro.Fukutoshin","odpt:timeOfOrigin":"2016-02-09T11:04:00+09:00","odpt:trainInformationText":"現在、平常どおり運転しています。","@type":"odpt:TrainInformation"},{"@context":"http://vocab.tokyometroapp.jp/context_odpt_TrainInformation.json","@id":"urn:ucode:_00001C000000000000010000030C3BE7","dc:date":"2017-02-09T22:55:03+09:00","dct:valid":"2017-02-09T23:00:03+09:00","odpt:operator":"odpt.Operator:TokyoMetro","odpt:railway":"odpt.Railway:TokyoMetro.Marunouchi","odpt:timeOfOrigin":"2016-01-26T11:31:00+09:00","odpt:trainInformationText":"現在、平常どおり運転しています。","@type":"odpt:TrainInformation"},{"@context":"http://vocab.tokyometroapp.jp/context_odpt_TrainInformation.json","@id":"urn:ucode:_00001C000000000000010000030C3BE6","dc:date":"2017-02-09T22:55:03+09:00","dct:valid":"2017-02-09T23:00:03+09:00","odpt:operator":"odpt.Operator:TokyoMetro","odpt:railway":"odpt.Railway:TokyoMetro.Hibiya","odpt:timeOfOrigin":"2016-02-09T11:01:00+09:00","odpt:trainInformationText":"現在、平常どおり運転しています。","@type":"odpt:TrainInformation"},{"@context":"http://vocab.tokyometroapp.jp/context_odpt_TrainInformation.json","@id":"urn:ucode:_00001C000000000000010000030C3BE2","dc:date":"2017-02-09T22:55:03+09:00","dct:valid":"2017-02-09T23:00:03+09:00","odpt:operator":"odpt.Operator:TokyoMetro","odpt:railway":"odpt.Railway:TokyoMetro.Chiyoda","odpt:timeOfOrigin":"2016-02-09T19:00:00+09:00","odpt:trainInformationText":"現在、平常どおり運転しています。","@type":"odpt:TrainInformation"},{"@context":"http://vocab.tokyometroapp.jp/context_odpt_TrainInformation.json","@id":"urn:ucode:_00001C000000000000010000030C3BE4","dc:date":"2017-02-09T22:55:03+09:00","dct:valid":"2017-02-09T23:00:03+09:00","odpt:operator":"odpt.Operator:TokyoMetro","odpt:railway":"odpt.Railway:TokyoMetro.Ginza","odpt:timeOfOrigin":"2016-01-18T21:36:00+09:00","odpt:trainInformationText":"現在、平常どおり運転しています。","@type":"odpt:TrainInformation"},{"@context":"http://vocab.tokyometroapp.jp/context_odpt_TrainInformation.json","@id":"urn:ucode:_00001C000000000000010000030C3BE5","dc:date":"2017-02-09T22:55:03+09:00","dct:valid":"2017-02-09T23:00:03+09:00","odpt:operator":"odpt.Operator:TokyoMetro","odpt:railway":"odpt.Railway:TokyoMetro.Hanzomon","odpt:timeOfOrigin":"2016-02-08T10:03:00+09:00","odpt:trainInformationText":"現在、平常どおり運転しています。","@type":"odpt:TrainInformation"},{"@context":"http://vocab.tokyometroapp.jp/context_odpt_TrainInformation.json","@id":"urn:ucode:_00001C000000000000010000030C3BE8","dc:date":"2017-02-09T22:55:03+09:00","dct:valid":"2017-02-09T23:00:03+09:00","odpt:operator":"odpt.Operator:TokyoMetro","odpt:railway":"odpt.Railway:TokyoMetro.Namboku","odpt:timeOfOrigin":"2016-02-09T10:50:00+09:00","odpt:trainInformationText":"現在、平常どおり運転しています。","@type":"odpt:TrainInformation"},{"@context":"http://vocab.tokyometroapp.jp/context_odpt_TrainInformation.json","@id":"urn:ucode:_00001C000000000000010000030C3BEA","dc:date":"2017-02-09T22:55:03+09:00","dct:valid":"2017-02-09T23:00:03+09:00","odpt:operator":"odpt.Operator:TokyoMetro","odpt:railway":"odpt.Railway:TokyoMetro.Yurakucho","odpt:timeOfOrigin":"2016-02-09T10:54:00+09:00","odpt:trainInformationText":"現在、平常どおり運転しています。","@type":"odpt:TrainInformation"},{"@context":"http://vocab.tokyometroapp.jp/context_odpt_TrainInformation.json","@id":"urn:ucode:_00001C000000000000010000030C3BE9","dc:date":"2017-02-09T22:55:03+09:00","dct:valid":"2017-02-09T23:00:03+09:00","odpt:operator":"odpt.Operator:TokyoMetro","odpt:railway":"odpt.Railway:TokyoMetro.Tozai","odpt:timeOfOrigin":"2016-01-30T14:50:00+09:00","odpt:trainInformationText":"現在、平常どおり運転しています。","@type":"odpt:TrainInformation"}]`)
	body := []byte(`[{"@context":"http://vocab.tokyometroapp.jp/context_odpt_TrainInformation.json","@id":"urn:ucode:_00001C000000000000010000030C3BE3","dc:date":"2017-02-15T11:55:02+09:00","dct:valid":"2017-02-15T12:00:02+09:00","odpt:operator":"odpt.Operator:TokyoMetro","odpt:railway":"odpt.Railway:TokyoMetro.Fukutoshin","odpt:timeOfOrigin":"2016-02-09T11:04:00+09:00","odpt:trainInformationText":"現在、平常どおり運転しています。","@type":"odpt:TrainInformation"},{"@context":"http://vocab.tokyometroapp.jp/context_odpt_TrainInformation.json","@id":"urn:ucode:_00001C000000000000010000030C3BE7","dc:date":"2017-02-15T11:55:02+09:00","dct:valid":"2017-02-15T12:00:02+09:00","odpt:operator":"odpt.Operator:TokyoMetro","odpt:railway":"odpt.Railway:TokyoMetro.Marunouchi","odpt:timeOfOrigin":"2016-02-13T11:15:00+09:00","odpt:trainInformationText":"現在、平常どおり運転しています。","@type":"odpt:TrainInformation"},{"@context":"http://vocab.tokyometroapp.jp/context_odpt_TrainInformation.json","@id":"urn:ucode:_00001C000000000000010000030C3BE6","dc:date":"2017-02-15T11:55:02+09:00","dct:valid":"2017-02-15T12:00:02+09:00","odpt:operator":"odpt.Operator:TokyoMetro","odpt:railway":"odpt.Railway:TokyoMetro.Hibiya","odpt:timeOfOrigin":"2016-02-13T13:40:00+09:00","odpt:trainInformationText":"現在、平常どおり運転しています。","@type":"odpt:TrainInformation"},{"@context":"http://vocab.tokyometroapp.jp/context_odpt_TrainInformation.json","@id":"urn:ucode:_00001C000000000000010000030C3BE2","dc:date":"2017-02-15T11:55:02+09:00","dct:valid":"2017-02-15T12:00:02+09:00","odpt:operator":"odpt.Operator:TokyoMetro","odpt:railway":"odpt.Railway:TokyoMetro.Chiyoda","odpt:timeOfOrigin":"2016-02-09T19:00:00+09:00","odpt:trainInformationText":"現在、平常どおり運転しています。","@type":"odpt:TrainInformation"},{"@context":"http://vocab.tokyometroapp.jp/context_odpt_TrainInformation.json","@id":"urn:ucode:_00001C000000000000010000030C3BE4","dc:date":"2017-02-15T11:55:02+09:00","dct:valid":"2017-02-15T12:00:02+09:00","odpt:operator":"odpt.Operator:TokyoMetro","odpt:railway":"odpt.Railway:TokyoMetro.Ginza","odpt:timeOfOrigin":"2016-02-10T11:14:00+09:00","odpt:trainInformationText":"現在、平常どおり運転しています。","@type":"odpt:TrainInformation"},{"@context":"http://vocab.tokyometroapp.jp/context_odpt_TrainInformation.json","@id":"urn:ucode:_00001C000000000000010000030C3BE5","dc:date":"2017-02-15T11:55:02+09:00","dct:valid":"2017-02-15T12:00:02+09:00","odpt:operator":"odpt.Operator:TokyoMetro","odpt:railway":"odpt.Railway:TokyoMetro.Hanzomon","odpt:timeOfOrigin":"2016-02-11T14:40:00+09:00","odpt:trainInformationText":"現在、平常どおり運転しています。","@type":"odpt:TrainInformation"},{"@context":"http://vocab.tokyometroapp.jp/context_odpt_TrainInformation.json","@id":"urn:ucode:_00001C000000000000010000030C3BE8","dc:date":"2017-02-15T11:55:02+09:00","dct:valid":"2017-02-15T12:00:02+09:00","odpt:operator":"odpt.Operator:TokyoMetro","odpt:railway":"odpt.Railway:TokyoMetro.Namboku","odpt:timeOfOrigin":"2016-02-09T10:50:00+09:00","odpt:trainInformationText":"現在、平常どおり運転しています。","@type":"odpt:TrainInformation"},{"@context":"http://vocab.tokyometroapp.jp/context_odpt_TrainInformation.json","@id":"urn:ucode:_00001C000000000000010000030C3BEA","dc:date":"2017-02-15T11:55:02+09:00","dct:valid":"2017-02-15T12:00:02+09:00","odpt:operator":"odpt.Operator:TokyoMetro","odpt:railway":"odpt.Railway:TokyoMetro.Yurakucho","odpt:timeOfOrigin":"2016-02-09T10:54:00+09:00","odpt:trainInformationText":"現在、平常どおり運転しています。","@type":"odpt:TrainInformation"},{"@context":"http://vocab.tokyometroapp.jp/context_odpt_TrainInformation.json","@id":"urn:ucode:_00001C000000000000010000030C3BE9","dc:date":"2017-02-15T11:55:02+09:00","dct:valid":"2017-02-15T12:00:02+09:00","odpt:operator":"odpt.Operator:TokyoMetro","odpt:railway":"odpt.Railway:TokyoMetro.Tozai","odpt:timeOfOrigin":"2016-02-15T09:00:00+09:00","odpt:trainInformationStatus":"ダイヤ乱れ","odpt:trainInformationText":"7時28分頃、南行徳〜行徳駅間で車両点検のため、ダイヤが乱れています。この影響で快速運転、ＪＲ中央・総武各駅停車 三鷹方面との直通運転を中止しています。只今、東京メトロ線、都営地下鉄線、ＪＲ線、京成線、西武線、つくばエクスプレス線に振替輸送を実施しています。詳しくは、駅係員にお尋ねください。","@type":"odpt:TrainInformation"}]`)

	//fmt.Println("==========")
	//fmt.Println(string(body))
	//fmt.Println("==========")

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
