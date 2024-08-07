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
	allPaths    [][]string
	// path        []string
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("go run . example.txt")
		return
	}
	ParseFile(os.Args[1])
	Duffs(farm.Start.Name, farm.End.Name, []string{}, &allPaths)
	result := Conflicts()
	PrintAllPaths()
	fmt.Println(result)
}

func ParseFile(file string) {
	open, or := os.ReadFile(file)
	Err(or)
	split := strings.Split(string(open), "\n")
	farm.AntNum, or = strconv.Atoi(strings.TrimSpace(split[0]))
	Err(or)

	farm.Rooms = make(map[string]*Room)

	for _, line := range split[1:] {
		line = strings.TrimSpace(line)
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

func Duffs(start, end string, path []string, allPaths *[][]string) {
	path = append(path, start)
	if start == end {
		pathCopy := make([]string, len(path))
		copy(pathCopy, path)
		*allPaths = append(*allPaths, pathCopy)
		return
	}

	StartRoom := farm.Rooms[start]
	for _, neighbor := range StartRoom.Jeran {
		visited := false
		for _, node := range path {
			if neighbor.Name == node {
				visited = true
				break
			}
		}
		if !visited {
			Duffs(neighbor.Name, end, path, allPaths)
		}
	}
}

func PrintAllPaths() {
	for i, path := range allPaths {
		fmt.Printf("Path %d: %v\n", i+1, path)
	}
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

func Conflicts() [][]string {
	var result [][]string
	mapindexs := make(map[int]bool)
	for index1, path1 := range allPaths {
		var indexs []int
		var found bool
		for index2 := index1 + 1; index2 < len(allPaths); index2++ {
			path2 := allPaths[index2]
			if Resolve(path1, path2) && !mapindexs[index2] {
				indexs = append(indexs, index2)
				mapindexs[index2] = true
				found = true
			}
		}

		if !found && !mapindexs[index1] {
			result = append(result, path1)
		} else if !mapindexs[index1] {
			indexs = append(indexs, index1)
			result = append(result, MinLenght(indexs))
		}

	}
	return result
}

func Resolve(path1, path2 []string) bool {
	for _, room1 := range path1[1 : len(path1)-1] {
		for _, room2 := range path2[1 : len(path2)-1] {
			if room1 == room2 {
				return true
			}
		}
	}
	return false
}

func MinLenght(indexs []int) []string {
	min := allPaths[indexs[0]]
	for _, index := range indexs[1:] {
		if len(min) > len(allPaths[index]) {
			min = allPaths[index]
		}
	}
	return min
}
