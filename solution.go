package solution

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

func main() {

	fmt.Printf(GetMessage())
}

func GetMessage() string {
	var a = emoji.Sprintf(":world_map:!")
	return ("Hello " + a)
}
