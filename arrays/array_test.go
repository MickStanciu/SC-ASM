package arrays

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Declaration_WhenMaxed(t *testing.T) {
	a1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a1)

	assert.EqualValues(t, 5, len(a1))
	assert.EqualValues(t, 5, cap(a1))
}

func Test_Declaration_WhenEmpty(t *testing.T) {
	a1 := [5]int{}
	fmt.Println(a1)

	assert.EqualValues(t, 5, len(a1))
	assert.EqualValues(t, 5, cap(a1))
}

func Test_Declaration_WhenNotMaxed(t *testing.T) {
	a1 := [5]int{1, 2, 3}
	fmt.Println(a1)

	assert.EqualValues(t, 5, len(a1))
	assert.EqualValues(t, 5, cap(a1))

}

func Test_Declaration_WhenPartiallyFilled(t *testing.T) {
	a1 := [5]int{1, 2, 4: 10}
	fmt.Println(a1)

	assert.EqualValues(t, 5, len(a1))
	assert.EqualValues(t, 5, cap(a1))
}

func Test_Declaration_WhenPartialFilling(t *testing.T) {
	a1 := [5]int{2: 99}
	fmt.Println(a1)

	assert.EqualValues(t, 5, len(a1))
	assert.EqualValues(t, 5, cap(a1))
	assert.EqualValues(t, 99, a1[2])
}

func Test_DeclarationAlternative(t *testing.T) {
	a1 := [...]string{"bmw", "audi", "mercedes"}
	fmt.Println(a1)

	assert.EqualValues(t, 3, len(a1))
	assert.EqualValues(t, 3, cap(a1))
}

func Test_Accessing(t *testing.T) {
	a1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a1)

	assert.EqualValues(t, 1, a1[0])
	assert.EqualValues(t, 3, a1[2])
	assert.EqualValues(t, 5, a1[len(a1)-1])
}

func Test_ChangingValues(t *testing.T) {
	a1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a1)
	a1[0] = 99
	fmt.Println(a1)

	assert.EqualValues(t, 99, a1[0])
	assert.EqualValues(t, 3, a1[2])
	assert.EqualValues(t, 5, a1[len(a1)-1])
}

func Test_CopyArrays(t *testing.T) {
	a1 := [5]int{1, 2, 3, 4, 5}
	a2 := a1
	fmt.Println(a1)
	fmt.Printf("Adr: %v\n", &a1[0])

	fmt.Println(a2)
	fmt.Printf("Adr: %v\n", &a2[0])

	assert.EqualValues(t, a1, a2)

	a1[0] = 99
	fmt.Println(a1)
	fmt.Println(a2)
}

func Test_PassByValue(t *testing.T) {
	a1 := [5]int{10, 11, 12, 13, 14}
	fmt.Println(a1)
	fmt.Printf("Adr: %v\n", &a1[0])
	printArrayDetailsByVal(a1)
	fmt.Println(a1)
}

func Test_PassByPointer(t *testing.T) {
	a1 := [5]int{10, 11, 12, 13, 14}
	fmt.Println(a1)
	fmt.Printf("Adr: %v\n", &a1[0])
	printArrayDetailsByPointer(&a1)
	fmt.Println(a1)
}

func printArrayDetailsByVal(val [5]int) {
	fmt.Println(val)
	fmt.Printf("Len: %d\n", len(val))
	fmt.Printf("Cap: %d\n", cap(val))
	fmt.Printf("Adr: %v\n", &val[0])
	fmt.Println("------")
	val[1] = 99
}

func printArrayDetailsByPointer(val *[5]int) {
	fmt.Println(*val)
	fmt.Printf("Len: %d\n", len(val))
	fmt.Printf("Cap: %d\n", cap(val))
	fmt.Printf("Adr: %v\n", &val[0])
	fmt.Println("------")
	val[1] = 99
}

func Test_PlayScan(t *testing.T) {
	scanTheArrayV1()

	scanTheArrayV2()
}

func scanTheArrayV1() {
	a1 := [5]int{10, 11, 12, 13, 14}
	for i := 0; i < len(a1); i++ {
		fmt.Printf("#%d - %d\n", i, a1[i])
	}
}

func scanTheArrayV2() {
	a1 := [5]int{10, 11, 12, 13, 14}
	for idx, val := range a1 {
		fmt.Printf("#%d - %d\n", idx, val)
	}
}
