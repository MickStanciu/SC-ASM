package maps

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_DeclarationWhenEmpty(t *testing.T) {
	m := make(map[string]int)
	fmt.Println(m)

	assert.Empty(t, len(m))
}

func Test_Declaration_WithValues(t *testing.T) {
	m := map[string]int{"Audi": 10000, "BMW": 11000, "Lexus": 12000}
	fmt.Println(m)
	assert.EqualValues(t, 3, len(m))
}

func Test_AddingValues(t *testing.T) {
	m := make(map[string]int)
	m["Audi"] = 10000
	fmt.Println(m)
	assert.EqualValues(t, 1, len(m))
}

func Test_DeletingValues(t *testing.T) {
	m := map[string]int{"Audi": 10000, "BMW": 11000, "Lexus": 12000}
	delete(m, "BMW")
	delete(m, "BLAH")
	fmt.Println(m)
	assert.EqualValues(t, 2, len(m))
}

func Test_PassingToAnotherFunction(t *testing.T) {
	m := map[string]int{"Audi": 10000, "BMW": 11000, "Lexus": 12000}
	m["Peugeot"] = 20000
	fmt.Println(m)
	addCar(m, "Renault", 30000)
	ac := func(x map[string]int, make string, value int) {
		x[make] = value
	}
	ac(m, "Ford", 35000)

	fmt.Println(m)
	assert.EqualValues(t, 6, len(m))
}

func Test_PlaySeek(t *testing.T) {
	m := map[string]int{"Audi": 10000, "BMW": 11000, "Lexus": 12000}
	renault, e := m["Renault"]
	assert.False(t, e)
	assert.Empty(t, renault)
}

func Test_PlayScan(t *testing.T) {
	scanTheMap()
}

func scanTheMap() {
	m := map[string]int{"Audi": 10000, "BMW": 11000, "Lexus": 12000}
	for key, val := range m {
		fmt.Printf("#%s - %d\n", key, val)
	}
}

func addCar(m map[string]int, make string, value int) {
	m[make] = value
}
