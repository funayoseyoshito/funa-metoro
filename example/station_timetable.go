package main

import (
	"os"

	"fmt"

	"github.com/funayoseyoshito/metro"
)

func main() {

	m := metro.NewMetro(os.Getenv("metoro_token"))
	s := m.GetStationTimeTableWithParam(&metro.Params{ODPStation: "odpt.Station:TokyoMetro.Tozai.Otemachi"})
	fmt.Println(s)
}
