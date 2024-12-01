package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Human struct {
	Firstname string
	Lastname  string
}

func RandomStuff() {
	emeka := Human{Firstname: "emeka", Lastname: "ony"}
	tt, err := json.Marshal(&emeka)
	if err != nil {
		log.Fatal("Error encoding the struct as json")
	}
	fmt.Printf("%v", tt)
}
