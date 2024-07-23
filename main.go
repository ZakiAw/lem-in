package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	room  Room
	start Start
	end   End
	ant   Ant
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("go run . example.txt")
		return
	}
	ParseFile(os.Args[1])
}

func ParseFile(file string) {
	open, or := os.ReadFile(file)
	Err(or)
	split := strings.Split(string(open), "\n")
	ant.AntNum, or = strconv.Atoi(split[0])
	Err(or)
}

func Err(Error error) string {
	if Error != nil {
		return "Error"
	}
	return ""
}
