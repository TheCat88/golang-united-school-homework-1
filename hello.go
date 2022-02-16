package main

import (
	"fmt"
	"strings"

	"github.com/kyokomi/emoji"
)

func main() {
	var a = emoji.Sprintf(":world_map:!")
	fmt.Printf(emoji.Sprintf("Hello " + strings.ReplaceAll(a, " ", "")))
}
