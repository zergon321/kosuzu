package main

import (
	"fmt"

	"github.com/zergon321/kosuzu"
)

type Person struct {
	Name    string
	Age     int32
	Numbers []byte
}

func main() {
	person := &Person{
		Name:    "Vasya",
		Age:     16,
		Numbers: []byte{32, 25, 78},
	}
	packet, err := kosuzu.Serialize(32, person)
	handleError(err)

	data, err := packet.Bytes()
	handleError(err)

	fmt.Println(data)

	restored := new(Person)
	err = kosuzu.Deserialize(packet, &restored)
	handleError(err)

	fmt.Println(restored)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
