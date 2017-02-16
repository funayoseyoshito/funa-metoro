package main

import (
	"os"

	"fmt"

	"github.com/funayoseyoshito/metro"
	"github.com/k0kubun/pp"
)

func main() {

	m := metro.NewMetro(os.Getenv("metoro_token"))
	s := m.GetStationTimeTableWithParam(&metro.Params{ODPStation: "odpt.Station:TokyoMetro.Chiyoda.Ayase"})
	//s := m.GetStationTimeTableWithParam(&metro.Params{ODPTRailway: "odpt.Railway:TokyoMetro.Chiyoda"})
	//fmt.Println(s)

	for _, v := range s {
		fmt.Println("=================")
		//fmt.Println(v)
		//fmt.Println(v.OdptStation)
		//fmt.Println(v)
		//pp.Println(v.OdptWeekdays)
		for i, t := range v.OdptWeekdays {
			if i > 0 && i < 2 {

				pp.Println(t)
			}
		}
		fmt.Println("=================")
	}

}
