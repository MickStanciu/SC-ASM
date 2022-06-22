package struct1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type grower interface {
	grow()
}

type Lettuce struct {
	Size      int `json:"size"`
	Condition int `json:"-"`
}

type Cucumber struct {
	Size int `json:"size"`
}

func NewLettuceV() Lettuce {
	return Lettuce{Size: 0, Condition: 0}
}

func NewCucumberV() Cucumber {
	return Cucumber{Size: 0}
}

func NewLettuceP() *Lettuce {
	return &Lettuce{Size: 0, Condition: 0}
}

func NewCucumberP() *Cucumber {
	return &Cucumber{Size: 0}
}

func (l Lettuce) grow() {
	l.Size++
}

func (c *Cucumber) grow() {
	c.Size++
}

func grow(g grower) {
	g.grow()
}

func TestShouldGrow_ButIsItReally(t *testing.T) {
	lv := NewLettuceV()
	grow(&lv)
	grow(lv)
	assert.EqualValues(t, 2, lv.Size)

	cv := NewCucumberV()
	grow(&cv)
	grow(&cv)
	//grow(cv)
	assert.EqualValues(t, 2, cv.Size)
}

func TestShouldGrow_But(t *testing.T) {
	lv := NewLettuceV()
	lv.grow()
	lv.grow()
	assert.EqualValues(t, 2, lv.Size)
}

func TestShouldGrow_AAAAHHHHH(t *testing.T) {
	lv := NewLettuceP()
	lv.grow()
	lv.grow()
	assert.EqualValues(t, 2, lv.Size)
}

func TestShouldGrow2(t *testing.T) {
	lv := NewLettuceP()
	grow(lv)
	grow(*lv)
	assert.EqualValues(t, 2, lv.Size)

	cv := NewCucumberP()
	grow(cv)
	grow(cv)
	//grow(&cv)
	assert.EqualValues(t, 2, cv.Size)
}

func TestShouldGrow_CoolBananas(t *testing.T) {
	c := NewCucumberP()
	c.grow()
	c.grow()
	grow(c)
	assert.EqualValues(t, 3, c.Size)
}
