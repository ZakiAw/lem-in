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

var (
	farm        Farm
	mwjoodStart bool
	mwjoodEnd   bool
	path		[]string
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

	farm.Rooms = make(map[string]*Room)

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
			continue
		}
		if mwjoodEnd {
			farm.End.Name = f[0]
			fmt.Println(farm.End.Name)
			mwjoodEnd = false
			continue
		}
		if strings.Contains(line, "-") {
			split := strings.Split(line, "-")
			from := split[0]
			to := split[1]
			mapmaker(from, to)
		}
	}
	fmt.Print(FindPath(farm.Start.Name, farm.End.Name))
}

func Err(Error error) {
	if Error != nil {
		fmt.Println(Error)
		os.Exit(0)
	}
}

func mapmaker(from, to string) {
	fromNode, fromExists := farm.Rooms[from]
	if !fromExists {
		fromNode = &Room{Name: from}
		farm.Rooms[from] = fromNode
	}
	toNode, toExists := farm.Rooms[to]
	if !toExists {
		toNode = &Room{Name: to}
		farm.Rooms[to] = toNode
	}
	fromNode.Jeran = append(fromNode.Jeran, toNode)
	toNode.Jeran = append(toNode.Jeran, fromNode)
}

func FindPath(start, end string) []string {
	
	path = append(path, start)
	Start := farm.Rooms[start]
	for _, i := range Start.Jeran {
		visited := false
		for _,o := range path {
			if i.Name == o {
		visited = true
			}
		}
		if visited {
			continue
		}
		if i.Name != end || i.Name != start {
			FindPath(i.Name, end)
		}
	}
	return path
}

func PrintFarm() {
	for _, room := range farm.Rooms {
		fmt.Printf("Room %s has neighbors: ", room.Name)
		for _, neighbor := range room.Jeran {
			fmt.Printf("%s ", neighbor.Name)
		}
		fmt.Println()
	}
}
