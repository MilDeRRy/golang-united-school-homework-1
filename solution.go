package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	pizzaMessage := emoji.Sprint("Hello :world_map:!")
	return pizzaMessage
}
