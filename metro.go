package metro

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
)

const api_base_url = "https://api.tokyometroapp.jp/api/v2"

type Metro struct {
	apiPath string
	params  params
}

type params struct {
	rdfType     string `param:"rdf:type"`
	consumerKey string `param:"acl:consumerKey"`
}

type ApiResponse interface {
	dump()
}

type ApiResults []ApiResponse

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

func (m *params) request(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}
