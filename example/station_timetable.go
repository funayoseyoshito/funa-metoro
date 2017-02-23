package main

import (
	"fmt"
	"os"

	"github.com/funayoseyoshito/metro"
)

func main() {
	m := metro.NewMetro(os.Getenv("metoro_token")).
		SetParam("odpt:station", "odpt.Station:TokyoMetro.Chiyoda.Ayase")
	s := m.ODPTStationTimeTable()

	for _, v := range s {

		fmt.Println("=================")
		fmt.Println(v.Context)
		fmt.Println(v.UCODE())
		fmt.Println(v.OwlSameAs)
		fmt.Println(v.DcDate)
		fmt.Println(v.OdptStation)
		fmt.Println(v.OdptRailway)
		fmt.Println(v.RailsWayName())
		fmt.Println(v.OdptOperator)
		fmt.Println(v.OperatorName())
		fmt.Println(v.OdptRailDirection)
		fmt.Println(v.RailDirection())

		fmt.Println("OdptWeekdays")
		for _, t := range v.OdptWeekdays {
			fmt.Println("--")
			fmt.Println(t.OdptDepartureTime)
			fmt.Println(t.OdptDestinationStation)
			fmt.Println(t.DestinationStation())
			fmt.Println(t.OdptTrainType)
			fmt.Println(t.TrainTypeName())
			fmt.Println(t.OdptIsLast)
			fmt.Println(t.OdptIsOrigin)
			fmt.Println(t.OdptCarComposition)
			fmt.Println(t.OdptNotes)
			fmt.Println("--")
		}

		//fmt.Println("OdptSaturdays")
		//for _, t := range v.OdptSaturdays {
		//fmt.Println("--")
		//	fmt.Println(t)
		//fmt.Println("--")
		//}

		//fmt.Println("OdptHolidays")
		//for _, t := range v.OdptHolidays {
		//fmt.Println("--")
		//	fmt.Println(t)
		//fmt.Println("--")
		//}

		fmt.Println("=================")
		break
	}
}
