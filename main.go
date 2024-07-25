package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	Name  string
	Jeran []*Room
}
type Farm struct {
	Rooms  map[string]*Room
	AntNum int
	Start  Room
	End    Room
}
type Ant struct{}

var (
	farm        Farm
	mwjoodStart bool
	mwjoodEnd   bool

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
	farm.AntNum, or = strconv.Atoi(split[0])
	Err(or)

	for _, line := range split[1:] {
		if line == "" {
			continue
		}
		f := strings.Fields(line)

		if line == "##start" {
			mwjoodStart = true
			continue
		}
		if line == "##end" {
			mwjoodEnd = true
			continue
		}
		if mwjoodStart {
			farm.Start.Name = f[0]
			fmt.Println(farm.Start.Name)
			mwjoodStart = false
		}
		if mwjoodEnd {
			farm.End.Name = f[0]
			fmt.Println(farm.End.Name)
			mwjoodEnd = false
		}

	}
}

func Err(Error error) {
	if Error != nil {
		fmt.Println(Error)
		os.Exit(0)
	}
}
