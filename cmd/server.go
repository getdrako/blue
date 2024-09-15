package main

import (
	"drako/server/pkg/build"
	"fmt"
)

func main() {
	fmt.Println("version:", build.Tag)
	fmt.Println("commit:", build.Commit)
	fmt.Println("datetime:", build.Datetime)
	fmt.Println("debug:", build.DRAKO_DEBUG)
}
