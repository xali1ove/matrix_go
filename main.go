package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Block struct {
	Color string
}

type Point struct {
	Row, Col int
}

func main() {
	var height, width, numColors int

	fmt.Println("Enter height, width, and number of colors:")
	fmt.Scanln(&height, &width, &numColors)

	matrix := generateMatrix(height, width, numColors)
	fmt.Println("Generated Matrix:")
	printMatrix(matrix)

	group := findLargestGroup(matrix)
	fmt.Println("Largest Group:")
	printGroup(group)

	fmt.Println("Matrix with Largest Group:")
	printMatrixWithGroup(matrix, group)
}

func generateMatrix(height, width, numColors int) [][]Block {
	matrix := make([][]Block, height)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < height; i++ {
		matrix[i] = make([]Block, width)
		for j := 0; j < width; j++ {
			colorIndex := rand.Intn(numColors)
			matrix[i][j] = Block{Color: fmt.Sprintf("Color%d", colorIndex+1)}
		}
	}

	return matrix
}

func findLargestGroup(matrix [][]Block) []Point {
	group := make([]Point, 0)
	visited := make(map[Point]bool)

	for i, row := range matrix {
		for j := range row {
			point := Point{i, j}
			if !visited[point] {
				currGroup := make([]Point, 0)
				dfs(matrix, point, matrix[i][j], visited, &currGroup)

				if len(currGroup) > len(group) {
					group = currGroup
				}
			}
		}
	}

	return group
}

func dfs(matrix [][]Block, point Point, target Block, visited map[Point]bool, group *[]Point) {
	if !isValidPoint(matrix, point, target, visited) {
		return
	}

	*group = append(*group, point)
	visited[point] = true

	directions := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, dir := range directions {
		newPoint := Point{point.Row + dir.Row, point.Col + dir.Col}
		dfs(matrix, newPoint, target, visited, group)
	}
}

func isValidPoint(matrix [][]Block, point Point, target Block, visited map[Point]bool) bool {
	if point.Row < 0 || point.Row >= len(matrix) || point.Col < 0 || point.Col >= len(matrix[0]) {
		return false
	}

	if visited[point] || matrix[point.Row][point.Col] != target {
		return false
	}

	return true
}

func printMatrix(matrix [][]Block) {
	for _, row := range matrix {
		for _, block := range row {
			fmt.Printf("[%s] ", block.Color)
		}
		fmt.Println()
	}
}

func printGroup(group []Point) {
	fmt.Println("Largest Group:")
	for _, point := range group {
		fmt.Printf("(%d, %d) ", point.Row+1, point.Col+1)
	}
	fmt.Println()
}

func printMatrixWithGroup(matrix [][]Block, group []Point) {
	for i, row := range matrix {
		for j, block := range row {
			if containsPoint(group, i, j) {
				fmt.Printf("[%-*s] ", getMaxColorLength(matrix), block.Color)
			} else {
				fmt.Printf("%-*s ", getMaxColorLength(matrix)+2, " ")
			}
		}
		fmt.Println()
	}
}

func getMaxColorLength(matrix [][]Block) int {
	maxLength := 0
	for _, row := range matrix {
		for _, block := range row {
			length := len(block.Color)
			if length > maxLength {
				maxLength = length
			}
		}
	}
	return maxLength
}

func containsPoint(group []Point, row, col int) bool {
	for _, point := range group {
		if point.Row == row && point.Col == col {
			return true
		}
	}
	return false
}
