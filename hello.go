package main

import (
	"strings"

	"github.com/kyokomi/emoji"
)

func GetMessage() string {
	var a = emoji.Sprintf(":world_map:!")
	return (emoji.Sprintf("Hello " + strings.ReplaceAll(a, " ", "")))
}
