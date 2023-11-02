package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Поиск в глубину и поменячаем где уже были в слайсе visit

func findCircleNum(isConnected [][]int) int {
	var provinceCount int

	cityCount := len(isConnected)
	visited := make([]bool, cityCount)

	for city, _ := range visited {
		if visited[city] {
			continue
		}
		provinceCount++

		visit(city, isConnected, visited)
	}

	return provinceCount
}

func visit(city int, connections [][]int, visited []bool) {
	visited[city] = true
	for next, isVisited := range visited {
		if connections[city][next] == 1 && !isVisited {
			visit(next, connections, visited)
		}
	}
}

func Test_Simple(t *testing.T) {
	for cities := 1; cities < 10; cities++ {
		connections := make([][]int, cities)
		for i := 0; i < cities; i++ {
			connections[i] = make([]int, cities)
		}

		assert.Equal(t, cities, findCircleNum(connections))
	}
}

func Test_WithCOnn(t *testing.T) {
	tests := []struct {
		name   string
		cities int
		conns  [][2]int
		expect int
	}{
		{
			name: "3 cities one province",
			conns: [][2]int{
				{0, 1}, {0, 2},
			},
			cities: 3,
			expect: 1,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			connections := make([][]int, tc.cities)
			for i := 0; i < tc.cities; i++ {
				connections[i] = make([]int, tc.cities)
			}

			for j := 0; j < len(tc.conns); j++ {
				city1, city2 := tc.conns[j][0], tc.conns[j][1]
				connections[city1][city2] = 1
				connections[city2][city1] = 1
			}

			assert.Equal(t, tc.expect, findCircleNum(connections))

		})
	}

}
