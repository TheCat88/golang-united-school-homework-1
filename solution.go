package solution

import (
	"fmt"
	"strings"

	"github.com/kyokomi/emoji"
)

func main() {

	fmt.Printf(GetMessage())
}

func GetMessage() string {
	var a = emoji.Sprintf(":world_map:!")
	return (emoji.Sprintf("Hello " + strings.ReplaceAll(a, " ", "")))
}
