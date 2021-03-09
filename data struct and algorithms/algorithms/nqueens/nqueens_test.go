package nqueens

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNQueens1(t *testing.T) {
	result := NQueens(2)
	assert.Equal(t, len(result), 0)
	//
	power := 1 << 2
	for _, item := range result {
		for _, line := range item {
			check := 1
			for check < power {
				if int(line) & check == 0 {
					fmt.Printf("0  ")
				} else {
					fmt.Printf("1  ")
				}
				check = check << 1
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}
}

func TestNQueens2(t *testing.T) {
	result := NQueens(3)
	assert.Equal(t, len(result), 0)
	//
	power := 1 << 3
	for _, item := range result {
		for _, line := range item {
			check := 1
			for check < power {
				if int(line) & check == 0 {
					fmt.Printf("0  ")
				} else {
					fmt.Printf("1  ")
				}
				check = check << 1
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}
}

func TestNQueens3(t *testing.T) {
	result := NQueens(4)
	assert.Equal(t, len(result), 2)
	//
	power := 1 << 4
	for _, item := range result {
		for _, line := range item {
			check := 1
			for check < power {
				if int(line) & check == 0 {
					fmt.Printf("0  ")
				} else {
					fmt.Printf("1  ")
				}
				check = check << 1
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}
}

func TestNQueens4(t *testing.T) {
	result := NQueens(8)
	assert.Equal(t, len(result), 92)
	//
	power := 1 << 8
	for _, item := range result {
		for _, line := range item {
			check := 1
			for check < power {
				if int(line) & check == 0 {
					fmt.Printf("0  ")
				} else {
					fmt.Printf("1  ")
				}
				check = check << 1
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}
}
