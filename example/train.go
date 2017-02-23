package main

import (
	"fmt"
	"os"

	"github.com/funayoseyoshito/metro"
)

func main() {
	m := metro.NewMetro(os.Getenv("metoro_token"))
	m.SetParam("odpt:railway", "odpt:railway=odpt.Railway:TokyoMetro.Chiyoda")
	t := m.ODPTTrain()

	for _, v := range t {
		fmt.Println("-----")
		fmt.Println(v.OdptTrainOwner.TrainOwnerName())
		fmt.Println(v.UCODE())
		fmt.Println(v.OwlSameAs)
		fmt.Println(v.OdptTrainNumber)
		fmt.Println(v.OdptTrainType)
		fmt.Println(v.OdptTrainType.TrainType())
		fmt.Println(v.DcDate)
		fmt.Println(v.DctValid)
		fmt.Println(v.OdptFrequency)
		fmt.Println(v.OdptRailway)
		fmt.Println(v.OdptRailway.RailWayName())
		fmt.Println(v.OdptTrainOwner)
		fmt.Println("-----")
	}
}
