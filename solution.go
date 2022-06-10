package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func main() {
	//fmt.Println("Hello :world_map:!")

	//emoji.Println(":beer: Beer!!!")

	pizzaMessage := emoji.Sprint("Hello :world_map:!")
	fmt.Println(pizzaMessage)
}
