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
	var s2 []int
	//s3 := []int{}
	fmt.Println(s1)
	fmt.Println(s2)
}

func Test_MultiCreate(t *testing.T) {
	var s1 []string
	var s2 = []string{}
	var s3 = make([]string, 0)
	var s4 = make([]string, 17)

	fmt.Println(s1)
	fmt.Printf("Len: %d\n", len(s1))
	fmt.Printf("Cap: %d\n", cap(s1))

	fmt.Println(s2)
	fmt.Printf("Len: %d\n", len(s2))
	fmt.Printf("Cap: %d\n", cap(s2))

	fmt.Println(s3)
	fmt.Printf("Len: %d\n", len(s3))
	fmt.Printf("Cap: %d\n", cap(s3))

	fmt.Println(s4)
	fmt.Printf("Len: %d\n", len(s4))
	fmt.Printf("Cap: %d\n", cap(s4))

	// INIT ELEMENTS
	//s1[0] = "BMW" // this will blow!
	//s2[0] = "BMW" // this will blow!
	//s3[0] = "BMW" // this will blow!
	s4[0] = "BMW"
	//s4[1] = "AUDI" // this will blow!

	// APPEND STUFF
	s1 = append(s1, "PEUGEOT")
	fmt.Printf("address s1: %v\n", &s1[0])
	fmt.Println(s1)
	fmt.Printf("S1 Len: %d\n", len(s1))
	fmt.Printf("S1 Cap: %d\n", cap(s1))

	s2 = append(s2, "PEUGEOT")
	fmt.Println(s2)
	fmt.Printf("S2 Len: %d\n", len(s2))
	fmt.Printf("S2 Cap: %d\n", cap(s2))

	s3 = append(s3, "PEUGEOT")
	fmt.Println(s3)
	fmt.Printf("S3 Len: %d\n", len(s3))
	fmt.Printf("S3 Cap: %d\n", cap(s3))

	fmt.Printf("address s4: %v\n", &s4[0])
	s4 = append(s4, "PEUGEOT")
	fmt.Printf("address s4: %v\n", &s4[0])
	fmt.Println(s4)
	fmt.Printf("S4 Len: %d\n", len(s4))
	fmt.Printf("S4 Cap: %d\n", cap(s4))
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
	slice := make([]int, 5, 5)
	slice[0] = 10
	fmt.Println(&slice[0])
	slice = append(slice, 11)
	fmt.Println(&slice[0])
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

func Test_bla1(t *testing.T) {
	slice4()
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
