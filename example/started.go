package main

import (
	"os"

	"github.com/funayoseyoshito/metro"
)

func main() {
	m := metro.NewMetro(os.Getenv("metoro_token"))
	t := m.GetODPTTrainInformation()
	t.Dump()
	//fmt.Println(t)
}
