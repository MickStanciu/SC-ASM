package slices

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CreateWithMake(t *testing.T) {
	s1 := make([]int, 5)
	fmt.Println(s1)
	assert.EqualValues(t, 5, len(s1))
	assert.EqualValues(t, 5, cap(s1))
}

func Test_CreateWithMakeAndCapacity(t *testing.T) {
	s1 := make([]int, 5, 10)
	fmt.Println(s1)
	assert.EqualValues(t, 5, len(s1))
	assert.EqualValues(t, 10, cap(s1))
}

func Test_CreateWithSliceLiteral(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(s1)
	assert.EqualValues(t, 5, len(s1))
	assert.EqualValues(t, 5, cap(s1))
}

func Test_CreateWithIndexPositions(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5, 9: 99}
	fmt.Println(s1)
	assert.EqualValues(t, 10, len(s1))
	assert.EqualValues(t, 10, cap(s1))
}

func Test_CreateEmptySlice(t *testing.T) {
	s1 := make([]int, 0)
	var s2 []int // s2 := []int{}
	fmt.Println(s1)
	fmt.Println(s2)
}

func Test_SliceIt(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[1:3]

	fmt.Println(s1)
	fmt.Printf("Adr: %v\n", &s1[0])
	fmt.Printf("Adr: %v\n", &s1[1])
	fmt.Printf("Len: %d\n", len(s1))
	fmt.Printf("Cap: %d\n", cap(s1))
	fmt.Println("------")

	fmt.Println(s2)
	fmt.Printf("Adr: %v\n", &s2[0])
	fmt.Printf("Len: %d\n", len(s2))
	fmt.Printf("Cap: %d\n", cap(s2))
}

func Test_SliceIt_WhenChangingElements(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[1:3]

	fmt.Println(s1)
	fmt.Printf("Adr: %v\n", &s1[0])
	fmt.Printf("Adr: %v\n", &s1[1])
	fmt.Printf("Len: %d\n", len(s1))
	fmt.Printf("Cap: %d\n", cap(s1))
	fmt.Println("------")

	fmt.Println(s2)
	fmt.Printf("Adr: %v\n", &s2[0])
	fmt.Printf("Len: %d\n", len(s2))
	fmt.Printf("Cap: %d\n", cap(s2))
	fmt.Println("------")

	s2[0] = 99
	fmt.Println(s1)
	fmt.Println(s2)
}

func Test_GrowTheSlice(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(s1)
	fmt.Printf("Len: %d\n", len(s1))
	fmt.Printf("Cap: %d\n", cap(s1))
	fmt.Println("------")

	s1 = append(s1, []int{10, 11, 12}...)
	fmt.Println(s1)
	fmt.Printf("Len: %d\n", len(s1))
	fmt.Printf("Cap: %d\n", cap(s1))
}

// slice sharing the same underlying array
func slice1() {
	slice := []int{1, 2, 3, 4, 5, 6}
	newSlice := slice[1:3]
	fmt.Println(slice)
	fmt.Println(newSlice)

	newSlice[0] = 99
	fmt.Println(slice)
	fmt.Println(newSlice)
}

// if there is no room, will be a new array
func slice2() {
	slice := []int{1, 2, 3, 4, 5, 6}
	newSlice := append(slice, 10)
	fmt.Println(slice)
	fmt.Println(newSlice)

	newSlice[0] = 99
	fmt.Println(slice)
	fmt.Println(newSlice)
}

// if there is room, will not create a new array
func slice3() {
	slice := []int{1, 2, 3, 4, 5, 6}
	newSlice := slice[1:3]
	newSlice = append(newSlice, 10)
	fmt.Println(slice)
	fmt.Println(newSlice)

	newSlice[0] = 99
	fmt.Println(slice)
	fmt.Println(newSlice)
}

// if there is no room, will create a new array
func slice4() {
	slice := []int{10, 20, 30, 40, 50, 60}
	newSlice := slice[1:3:3]
	fmt.Println(slice)
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice = append(newSlice, 70) //capacity was maxed, new slice
	newSlice[0] = 99
	fmt.Println(slice)
	fmt.Println(newSlice)
}
