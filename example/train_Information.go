package main

import (
	"os"

	"fmt"

	"github.com/funayoseyoshito/metro"
)

func main() {
	m := metro.NewMetro(os.Getenv("metoro_token")).SetParam("odpt:trainInformationStatus", "ダイヤ乱れ")
	t := m.ODPTTrainInformation()

	for _, v := range t {
		fmt.Println("=================")
		fmt.Println(v)
		fmt.Println("-----")
		fmt.Println(v.UCODE())
		fmt.Println(v.OperatorName())
		fmt.Println(v.RailsWayName())
		fmt.Println("=================")
	}
}
