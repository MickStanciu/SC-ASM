package struct1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type speaker interface {
	speak(string)
}

type killer interface {
	kill(*animal)
}

type animal struct {
	species string
	alive   bool
}

type mammal struct {
	*animal
	hasTail bool
}

type reptile struct {
	*animal
	skinColour string
}

func (a *animal) kill(b *animal) {
	b.alive = false
}

func (a *animal) speak(val string) {
	fmt.Println(val)
}

func execute(k killer, v *animal) {
	k.kill(v)
}

func TestKill(t *testing.T) {
	rat := &mammal{
		animal: &animal{
			species: "rodent",
			alive:   true,
		},
		hasTail: true,
	}

	lizard := &reptile{
		animal: &animal{
			species: "who knows",
			alive:   true,
		},
		skinColour: "green",
	}

	mouse := &mammal{
		animal: &animal{
			species: "rodent",
			alive:   true,
		},
		hasTail: true,
	}

	rat.kill(lizard.animal)
	assert.False(t, lizard.alive)

	rat.speak("squeak")

	execute(rat, mouse.animal)
	assert.False(t, mouse.alive)
}
