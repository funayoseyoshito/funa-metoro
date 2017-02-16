package main

import (
	"os"

	"fmt"

	"github.com/funayoseyoshito/metro"
)

func main() {

	m := metro.NewMetro(os.Getenv("metoro_token"))
	s := m.GetStationTimeTableWithParam(&metro.Params{ODPStation: "odpt.Station:TokyoMetro.Chiyoda.Ayase"})
	//s := m.GetStationTimeTableWithParam(&metro.Params{ODPTRailway: "odpt.Railway:TokyoMetro.Chiyoda"})
	//fmt.Println(s)

	for _, v := range s {
		fmt.Println("=================")

		fmt.Println(v.UCODE())
		fmt.Println(v.OwlSameAs)
		fmt.Println(v.DcDate)
		fmt.Println(v.OdptStation)
		fmt.Println(v.StationName())
		//fmt.Println(v.OdptStation)
		//fmt.Println(v.OdptStation, v.OdptRailDirection)

		//fmt.Println(v)
		//fmt.Println(v.OdptStation)
		//fmt.Println(v)
		//pp.Println(v.OdptWeekdays)

		//for _, t := range v.OdptWeekdays {
		//	fmt.Println(t)
		//}

		//for _, t := range v.OdptSaturdays {
		//	fmt.Println(t)
		//}

		//for _, t := range v.OdptHolidays {
		//	fmt.Println(t)
		//}

		fmt.Println("=================")
	}
}
