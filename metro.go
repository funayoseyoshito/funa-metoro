package metro

import (
	"fmt"
	//"io/ioutil"
	"log"
	//"net/http"
	"net/url"
	"reflect"
	"encoding/json"
)

const api_base_url = "https://api.tokyometroapp.jp/api/v2"

type Metro struct {
	apiPath string
	params  params
	response ApiResponse
}

type params struct {
	rdfType     string `param:"rdf:type"`
	consumerKey string `param:"acl:consumerKey"`
}

type ApiResponse interface {
	Dump()
}

var apiResults interface{}

//NewMetro return New Metro struct
func NewMetro(token string) *Metro {
	m := &Metro{}
	m.params.consumerKey = token
	return m
}

func (m *Metro) getRequestUrl() string {

	u, err := url.Parse(api_base_url)
	if err != nil {
		log.Fatal(err)
	}

	u.Path += "/" + m.apiPath
	q := u.Query()

	rt, rv := reflect.TypeOf(m.params), reflect.ValueOf(m.params)
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		k := field.Tag.Get("param")
		v := rv.Field(i)
		q.Add(k, v.String())
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func (m *params) set() {
	fmt.Println("hello")
}

func (m *Metro) Execute() ApiResponse {

	//url := m.getRequestUrl()
	//res, err := http.Get(url)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//body, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}

	body := []byte(`[{"@context":"http://vocab.tokyometroapp.jp/context_odpt_TrainInformation.json","@id":"urn:ucode:_00001C000000000000010000030C3BE3","dc:date":"2017-02-09T22:55:03+09:00","dct:valid":"2017-02-09T23:00:03+09:00","odpt:operator":"odpt.Operator:TokyoMetro","odpt:railway":"odpt.Railway:TokyoMetro.Fukutoshin","odpt:timeOfOrigin":"2016-02-09T11:04:00+09:00","odpt:trainInformationText":"現在、平常どおり運転しています。","@type":"odpt:TrainInformation"},{"@context":"http://vocab.tokyometroapp.jp/context_odpt_TrainInformation.json","@id":"urn:ucode:_00001C000000000000010000030C3BE7","dc:date":"2017-02-09T22:55:03+09:00","dct:valid":"2017-02-09T23:00:03+09:00","odpt:operator":"odpt.Operator:TokyoMetro","odpt:railway":"odpt.Railway:TokyoMetro.Marunouchi","odpt:timeOfOrigin":"2016-01-26T11:31:00+09:00","odpt:trainInformationText":"現在、平常どおり運転しています。","@type":"odpt:TrainInformation"},{"@context":"http://vocab.tokyometroapp.jp/context_odpt_TrainInformation.json","@id":"urn:ucode:_00001C000000000000010000030C3BE6","dc:date":"2017-02-09T22:55:03+09:00","dct:valid":"2017-02-09T23:00:03+09:00","odpt:operator":"odpt.Operator:TokyoMetro","odpt:railway":"odpt.Railway:TokyoMetro.Hibiya","odpt:timeOfOrigin":"2016-02-09T11:01:00+09:00","odpt:trainInformationText":"現在、平常どおり運転しています。","@type":"odpt:TrainInformation"},{"@context":"http://vocab.tokyometroapp.jp/context_odpt_TrainInformation.json","@id":"urn:ucode:_00001C000000000000010000030C3BE2","dc:date":"2017-02-09T22:55:03+09:00","dct:valid":"2017-02-09T23:00:03+09:00","odpt:operator":"odpt.Operator:TokyoMetro","odpt:railway":"odpt.Railway:TokyoMetro.Chiyoda","odpt:timeOfOrigin":"2016-02-09T19:00:00+09:00","odpt:trainInformationText":"現在、平常どおり運転しています。","@type":"odpt:TrainInformation"},{"@context":"http://vocab.tokyometroapp.jp/context_odpt_TrainInformation.json","@id":"urn:ucode:_00001C000000000000010000030C3BE4","dc:date":"2017-02-09T22:55:03+09:00","dct:valid":"2017-02-09T23:00:03+09:00","odpt:operator":"odpt.Operator:TokyoMetro","odpt:railway":"odpt.Railway:TokyoMetro.Ginza","odpt:timeOfOrigin":"2016-01-18T21:36:00+09:00","odpt:trainInformationText":"現在、平常どおり運転しています。","@type":"odpt:TrainInformation"},{"@context":"http://vocab.tokyometroapp.jp/context_odpt_TrainInformation.json","@id":"urn:ucode:_00001C000000000000010000030C3BE5","dc:date":"2017-02-09T22:55:03+09:00","dct:valid":"2017-02-09T23:00:03+09:00","odpt:operator":"odpt.Operator:TokyoMetro","odpt:railway":"odpt.Railway:TokyoMetro.Hanzomon","odpt:timeOfOrigin":"2016-02-08T10:03:00+09:00","odpt:trainInformationText":"現在、平常どおり運転しています。","@type":"odpt:TrainInformation"},{"@context":"http://vocab.tokyometroapp.jp/context_odpt_TrainInformation.json","@id":"urn:ucode:_00001C000000000000010000030C3BE8","dc:date":"2017-02-09T22:55:03+09:00","dct:valid":"2017-02-09T23:00:03+09:00","odpt:operator":"odpt.Operator:TokyoMetro","odpt:railway":"odpt.Railway:TokyoMetro.Namboku","odpt:timeOfOrigin":"2016-02-09T10:50:00+09:00","odpt:trainInformationText":"現在、平常どおり運転しています。","@type":"odpt:TrainInformation"},{"@context":"http://vocab.tokyometroapp.jp/context_odpt_TrainInformation.json","@id":"urn:ucode:_00001C000000000000010000030C3BEA","dc:date":"2017-02-09T22:55:03+09:00","dct:valid":"2017-02-09T23:00:03+09:00","odpt:operator":"odpt.Operator:TokyoMetro","odpt:railway":"odpt.Railway:TokyoMetro.Yurakucho","odpt:timeOfOrigin":"2016-02-09T10:54:00+09:00","odpt:trainInformationText":"現在、平常どおり運転しています。","@type":"odpt:TrainInformation"},{"@context":"http://vocab.tokyometroapp.jp/context_odpt_TrainInformation.json","@id":"urn:ucode:_00001C000000000000010000030C3BE9","dc:date":"2017-02-09T22:55:03+09:00","dct:valid":"2017-02-09T23:00:03+09:00","odpt:operator":"odpt.Operator:TokyoMetro","odpt:railway":"odpt.Railway:TokyoMetro.Tozai","odpt:timeOfOrigin":"2016-01-30T14:50:00+09:00","odpt:trainInformationText":"現在、平常どおり運転しています。","@type":"odpt:TrainInformation"}]`)

	//fmt.Println("==========")
	//fmt.Println(string(body))
	//fmt.Println("==========")

	err := json.Unmarshal(body, &m.response)
	if err != nil {
		fmt.Println("ここ！")
		log.Fatal(err)
	}

	return m.response
}
