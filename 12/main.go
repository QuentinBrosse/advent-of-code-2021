package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Path []*Room

func (p Path) AlreadyVisited(target *Room) bool {
	for _, room := range p {
		if room.Name == target.Name {
			return true
		}
	}
	return false
}

func (p Path) LastRoom() *Room {
	return p[len(p)-1]
}

type Room struct {
	Name        string
	Small       bool
	Connections map[string]*Room
}

type Cave struct {
	connections map[string]*Room
}

func NewCave() *Cave {
	return &Cave{
		connections: make(map[string]*Room),
	}
}

func (c *Cave) addRoom(key string) *Room {
	vertex := &Room{
		Name:        key,
		Connections: make(map[string]*Room),
	}
	if strings.ToLower(key) == key {
		vertex.Small = true
	}
	c.connections[key] = vertex
	return vertex
}

func (c *Cave) AddConnection(key1, key2 string) {
	vertex1 := c.connections[key1]
	if vertex1 == nil {
		vertex1 = c.addRoom(key1)
	}

	vertex2 := c.connections[key2]
	if vertex2 == nil {
		vertex2 = c.addRoom(key2)
	}

	vertex1.Connections[key2] = vertex2
	vertex2.Connections[key1] = vertex1
}

func (c *Cave) GetRoom(key string) *Room {
	return c.connections[key]
}

func (c *Cave) CountAllPaths(startKey, endKey string, canVisitASmallRoomTwice bool) int {
	start, end := c.GetRoom(startKey), c.GetRoom(endKey)
	return run(start, end, Path{start}, !canVisitASmallRoomTwice)
}

func run(start, end *Room, path Path, noMoreDouble bool) (count int) {
	for _, room := range path.LastRoom().Connections {
		if room == end {
			count += 1
			continue
		}
		if room == start {
			continue
		}

		nmd := noMoreDouble
		if room.Small && path.AlreadyVisited(room) {
			if nmd {
				continue
			}
			nmd = true
		}

		count += run(start, end, append(path, room), nmd)
	}
	return count
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(file)
	}

	cave := NewCave()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "-")
		cave.AddConnection(parts[0], parts[1])
	}

	startKey, endKey := "start", "end"

	fmt.Println("Part 1:", cave.CountAllPaths(startKey, endKey, false))
	fmt.Println("Part 2:", cave.CountAllPaths(startKey, endKey, true))
}
