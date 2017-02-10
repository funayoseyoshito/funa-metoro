package main

import (
	"fmt"
	"os"

	"github.com/funayoseyoshito/metro"
)

type user struct {
	name string
}

type Human interface {
	hello()
}

func (h *user) hello() {
	fmt.Println("hello " + h.name)
}

var tmp []Human

func main() {
	//i, _ := tmp.([]user)
	//fmt.Println(i)
	//i, _ := tmp.(*user)
	//fmt.Printf("%T \n", i)
	//fmt.Printf("%T \n", tmp)
	//tmp = i
	//fmt.Printf("%T \n", tmp)
	//panic("================>ok")

	m := metro.NewMetro(os.Getenv("metoro_token"))
	r := m.GetODPTTrainInformation().Execute()
	r.Parse()
	//fmt.Println()
	//r.Parse()
	//r.Dump()
	//r.SetTest()
	//r.Dump()
	//vv, ok := r.(*metro.TrainInformations)

	//fmt.Println(ok)
	//fmt.Println(vv)

	//for i, v := range *vv {
	//	fmt.Println(i, v)
	//}
	//
	//vv.Yoshito()
	//fmt.Println("````````````````````")
	//r.Dump()

	//m = metro.NewMetro(os.Getenv("metoro_token"))
	//r = m.GetODPTTrainInformation2().Execute()
	//r.Dump()
}
