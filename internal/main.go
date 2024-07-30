package main

import (
	"encoding/json"
	"github.com/mailru/easyjson"
	"log"
)

// easyjson:json
type Peoples struct {
	p []People
}

// easyjson:json
type People struct {
	name string
	id   string
}

func main() {
	pe := People{name: "Jogn", id: "23123"}
	peS := Peoples{}
	for i := 0; i < 100; i++ {
		peS.p = append(peS.p, pe)

	}
	peBytes, _ := json.Marshal(peS)
	peBytesEasy, _ := easyjson.Marshal(peS)
	println("SIZE :::: ", len(peBytes))
	println("SIZE :::: ", len(peBytesEasy))
	log.Println(peBytesEasy)
	log.Println(peBytes)

}
