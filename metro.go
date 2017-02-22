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

type Param map[string]string

//NewMetro return New Metro struct
func NewMetro(token string) *Metro {
	m := &Metro{}
	m.param = make(Param, 5)
	m.SetParam("acl:consumerKey", token)
	return m
}

func (m *Metro) SetParam(k string, v interface{}) *Metro {

	var value string
	switch x := v.(type) {
	case int:
		value = strconv.Itoa(x)
	case string:
		value = x
	case bool:
		value = strconv.FormatBool(x)
	default:
		panic("param must be string or int or bool")
	}

	m.param[k] = value

	return m
}

func (m *Metro) isSetParam(k string) bool {
	_, ok := m.param[k]
	return ok
}

func (m *Metro) getRequestURI(p Param) string {
	u, err := url.Parse(APIBaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	u.Path += "/" + m.apiPath
	q := u.Query()

	for k, v := range p {
		q.Add(k, v)
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
	o := map[string]string{
		"odpt.Operator:TokyoMetro": "東京メトロ",
	}
	return o[id]
}

func getODPTTrainType(id string) string {
	t := map[string]string{
		"odpt.TrainType:TokyoMetro.Unknown":                "不明",
		"odpt.TrainType:TokyoMetro.Local":                  "各停",
		"odpt.TrainType:TokyoMetro.Express":                "急行",
		"odpt.TrainType:TokyoMetro.Rapid":                  "快速",
		"odpt.TrainType:TokyoMetro.SemiExpress":            "準急",
		"odpt.TrainType:TokyoMetro.TamaExpress":            "多摩急行",
		"odpt.TrainType:TokyoMetro.HolidayExpress":         "土休急行",
		"odpt.TrainType:TokyoMetro.CommuterSemiExpress":    "通勤準急",
		"odpt.TrainType:TokyoMetro.Extra":                  "臨時",
		"odpt.TrainType:TokyoMetro.RomanceCar":             "特急ロマンスカー",
		"odpt.TrainType:TokyoMetro.RapidExpress":           "快速急行",
		"odpt.TrainType:TokyoMetro.CommuterExpress":        "通勤急行",
		"odpt.TrainType:TokyoMetro.LimitedExpress":         "特急",
		"odpt.TrainType:TokyoMetro.CommuterLimitedExpress": "通勤特急",
		"odpt.TrainType:TokyoMetro.CommuterRapid":          "通勤快速",
		"odpt.TrainType:TokyoMetro.ToyoRapid":              "東葉快速",
		"odpt.TrainType:TokyoMetro.F-Liner":                "Fライナー",
	}
	return t[id]
}

func getODPTStationName(id string) string {
	s := map[string]string{
		"odpt.Station:JR-East.Joban.Abiko":                           "我孫子",
		"odpt.Station:JR-East.Joban.Toride":                          "取手",
		"odpt.Station:JR-East.Joban.Kashiwa":                         "柏",
		"odpt.Station:JR-East.Joban.Matsudo":                         "松戸",
		"odpt.Station:JR-East.Chuo.Mitaka":                           "三鷹",
		"odpt.Station:JR-East.ChuoChikatetsuTozai.Tsudanuma":         "津田沼",
		"odpt.Station:Toei.Mita.Mita":                                "三田",
		"odpt.Station:Toei.Mita.Shibakoen":                           "芝公園",
		"odpt.Station:Toei.Mita.Onarimon":                            "御成門",
		"odpt.Station:Toei.Mita.Uchisaiwaicho":                       "内幸町",
		"odpt.Station:Toei.Mita.Hibiya":                              "日比谷",
		"odpt.Station:Toei.Mita.Otemachi":                            "大手町",
		"odpt.Station:Toei.Mita.Jimbocho":                            "神保町",
		"odpt.Station:Toei.Mita.Suidobashi":                          "水道橋",
		"odpt.Station:Toei.Mita.Kasuga":                              "春日",
		"odpt.Station:Toei.Mita.Hakusan":                             "白山",
		"odpt.Station:Toei.Mita.Sengoku":                             "千石",
		"odpt.Station:Toei.Mita.Sugamo":                              "巣鴨",
		"odpt.Station:Toei.Mita.NishiSugamo":                         "西巣鴨",
		"odpt.Station:Toei.Mita.ShinItabashi":                        "新板橋",
		"odpt.Station:Toei.Mita.Itabashikuyakushomae":                "板橋区役所前",
		"odpt.Station:Toei.Mita.Itabashihoncho":                      "板橋本町",
		"odpt.Station:Toei.Mita.Motohasunuma":                        "本蓮沼",
		"odpt.Station:Toei.Mita.ShimuraSanchome":                     "志村坂上",
		"odpt.Station:Toei.Mita.Hasune":                              "蓮根",
		"odpt.Station:Toei.Mita.Nishidai":                            "西台",
		"odpt.Station:Toei.Mita.Takashimadaira":                      "高島平",
		"odpt.Station:Toei.Mita.ShinTakashimadaira":                  "新高島平",
		"odpt.Station:Toei.Mita.NishiTakashimadaira":                 "西高島平",
		"odpt.Station:SaitamaRailway.SaitamaRailway.UrawaMisono":     "浦和美園",
		"odpt.Station:SaitamaRailway.SaitamaRailway.Hatogaya":        "鳩ヶ谷",
		"odpt.Station:ToyoRapidRailway.ToyoRapid.ToyoKatsutadai":     "東葉勝田台",
		"odpt.Station:ToyoRapidRailway.ToyoRapid.YachiyoMidorigaoka": "八千代緑が丘",
		"odpt.Station:Odakyu.Tama.Karakida":                          "唐木田",
		"odpt.Station:Odakyu.Odawara.HonAtsugi":                      "本厚木",
		"odpt.Station:Odakyu.Odawara.HakoneYumoto":                   "箱根湯本",
		"odpt.Station:Odakyu.Odawara.Ebina":                          "海老名",
		"odpt.Station:Tobu.Nikko.MinamiKurihashi":                    "南栗橋",
		"odpt.Station:Tobu.Isesaki.Kuki":                             "久喜 　",
		"odpt.Station:Tobu.Isesaki.Takenotsuka":                      "竹ノ塚",
		"odpt.Station:Tobu.Isesaki.KitaKasukabe":                     "北春日部",
		"odpt.Station:Tobu.Isesaki.KitaKoshigaya":                    "北越谷",
		"odpt.Station:Tobu.Isesaki.TobuDoubutuKouen":                 "東武動物公園",
		"odpt.Station:Tobu.Tojo.Asaka":                               "朝霞",
		"odpt.Station:Tobu.Tojo.Asakadai":                            "朝霞台",
		"odpt.Station:Tobu.Tojo.Shiki":                               "志木",
		"odpt.Station:Tobu.Tojo.Yanasegawa":                          "柳瀬川",
		"odpt.Station:Tobu.Tojo.Mizuhodai":                           "みずほ台",
		"odpt.Station:Tobu.Tojo.Tsuruse":                             "鶴瀬",
		"odpt.Station:Tobu.Tojo.Fujimino":                            "ふじみ野",
		"odpt.Station:Tobu.Tojo.KamiFukuoka":                         "上福岡",
		"odpt.Station:Tobu.Tojo.Shingashi":                           "新河岸",
		"odpt.Station:Tobu.Tojo.Kawagoe":                             "川越",
		"odpt.Station:Tobu.Tojo.Kawagoeshi":                          "川越市",
		"odpt.Station:Tobu.Tojo.Kasumigaseki":                        "霞ヶ関",
		"odpt.Station:Tobu.Tojo.Tsurugashima":                        "鶴ヶ島",
		"odpt.Station:Tobu.Tojo.Wakaba":                              "若葉",
		"odpt.Station:Tobu.Tojo.Sakado":                              "坂戸",
		"odpt.Station:Tobu.Tojo.KitaSakado":                          "北坂戸",
		"odpt.Station:Tobu.Tojo.Takasaka":                            "高坂",
		"odpt.Station:Tobu.Tojo.HigashiMatsuyama":                    "東松山",
		"odpt.Station:Tobu.Tojo.ShinrinKoen":                         "森林公園",
		"odpt.Station:Tokyu.Toyoko.Hiyoshi":                          "日吉",
		"odpt.Station:Tokyu.Toyoko.MusashiKosugi":                    "武蔵小杉",
		"odpt.Station:Tokyu.Toyoko.Yokohama":                         "横浜",
		"odpt.Station:Tokyu.Toyoko.Kikuna":                           "菊名",
		"odpt.Station:Tokyu.Toyoko.Motosumiyoshi":                    "元住吉",
		"odpt.Station:Tokyu.Toyoko.Okusawa":                          "奥沢",
		"odpt.Station:Tokyu.Meguro.Hiyoshi":                          "日吉",
		"odpt.Station:Tokyu.Meguro.Okusawa":                          "奥沢",
		"odpt.Station:Tokyu.Meguro.Motosumiyoshi":                    "元住吉",
		"odpt.Station:Tokyu.Meguro.MusashiKosugi":                    "武蔵小杉",
		"odpt.Station:Tokyu.DenEnToshi.FutakoTamagawa":               "二子玉川",
		"odpt.Station:Tokyu.DenEnToshi.Nagatsuta":                    "長津田",
		"odpt.Station:Tokyu.DenEnToshi.Saginuma":                     "鷺沼",
		"odpt.Station:Tokyu.DenEnToshi.ChuoRinkan":                   "中央林間",
		"odpt.Station:Minatomirai.Minatomirai.MotomachiChukagai":     "元町・中華街",
		"odpt.Station:Seibu.Ikebukuro.ShinSakuradai":                 "新桜台",
		"odpt.Station:Seibu.Ikebukuro.Nerima":                        "練馬",
		"odpt.Station:Seibu.Ikebukuro.Nakamurabashi":                 "中村橋",
		"odpt.Station:Seibu.Ikebukuro.Fujimidai":                     "富士見台",
		"odpt.Station:Seibu.Ikebukuro.NerimaTakanodai":               "練馬高野台",
		"odpt.Station:Seibu.Ikebukuro.ShakujiiKoen":                  "石神井公園",
		"odpt.Station:Seibu.Ikebukuro.OizumiGakuen":                  "大泉学園",
		"odpt.Station:Seibu.Ikebukuro.Hoya":                          "保谷",
		"odpt.Station:Seibu.Ikebukuro.Hibarigaoka":                   "ひばりヶ丘",
		"odpt.Station:Seibu.Ikebukuro.HigashiKurume":                 "東久留米",
		"odpt.Station:Seibu.Ikebukuro.Kiyose":                        "清瀬",
		"odpt.Station:Seibu.Ikebukuro.Akitsu":                        "秋津",
		"odpt.Station:Seibu.Ikebukuro.Tokorozawa":                    "所沢",
		"odpt.Station:Seibu.Ikebukuro.NishiTokorozawa":               "西所沢",
		"odpt.Station:Seibu.Ikebukuro.Kotesashi":                     "小手指",
		"odpt.Station:Seibu.Ikebukuro.Sayamagaoka":                   "狭山ヶ丘",
		"odpt.Station:Seibu.Ikebukuro.MusashiFujisawa":               "武蔵藤沢",
		"odpt.Station:Seibu.Ikebukuro.InariyamaKoen":                 "稲荷山公園",
		"odpt.Station:Seibu.Ikebukuro.Irumashi":                      "入間市",
		"odpt.Station:Seibu.Ikebukuro.Bushi":                         "仏子",
		"odpt.Station:Seibu.Ikebukuro.Motokaji":                      "元加治",
		"odpt.Station:Seibu.Ikebukuro.Hanno":                         "飯能",
	}
	return s[id]
}

func getODPTRailDirectionName(id string) string {
	d := map[string]string{
		"odpt.RailDirection:TokyoMetro.Asakusa":           "浅草方面",
		"odpt.RailDirection:TokyoMetro.Ogikubo":           "荻窪方面",
		"odpt.RailDirection:TokyoMetro.Ikebukuro":         "池袋方面",
		"odpt.RailDirection:TokyoMetro.Honancho":          "方南町方面",
		"odpt.RailDirection:TokyoMetro.NakanoSakaue":      "中野坂上方面",
		"odpt.RailDirection:TokyoMetro.NakaMeguro":        "中目黒方面",
		"odpt.RailDirection:TokyoMetro.KitaSenju":         "北千住方面",
		"odpt.RailDirection:TokyoMetro.NishiFunabashi":    "西船橋方面",
		"odpt.RailDirection:TokyoMetro.Nakano":            "中野方面",
		"odpt.RailDirection:TokyoMetro.YoyogiUehara":      "代々木上原方面",
		"odpt.RailDirection:TokyoMetro.Ayase":             "綾瀬方面",
		"odpt.RailDirection:TokyoMetro.KitaAyase":         "北綾瀬方面",
		"odpt.RailDirection:TokyoMetro.ShinKiba":          "新木場方面",
		"odpt.RailDirection:TokyoMetro.Oshiage":           "押上方面",
		"odpt.RailDirection:TokyoMetro.Shibuya":           "渋谷方面",
		"odpt.RailDirection:TokyoMetro.AkabaneIwabuchi":   "赤羽岩淵方面",
		"odpt.RailDirection:TokyoMetro.Meguro":            "目黒方面",
		"odpt.RailDirection:TokyoMetro.ShirokaneTakanawa": "白金高輪方面",
		"odpt.RailDirection:TokyoMetro.Wakoshi":           "和光市方面",
		"odpt.RailDirection:TokyoMetro.KotakeMukaihara":   "小竹向原方面",
	}
	return d[id]
}

func getODPTRailWayName(id string) string {
	r := map[string]string{
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
	return r[id]
}
