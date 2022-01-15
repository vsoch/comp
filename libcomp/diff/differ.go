package diff

import (
	"fmt"
	"time"

	"github.com/vsoch/comp/libcomp/comp"
	"github.com/vsoch/comp/libcomp/env"
)

type Differ struct {
	SrcA      string
	SrcB      string
	EnvA      *env.Environment
	EnvB      *env.Environment
	CreatedAt time.Time
}

func GetEnv(src string) *env.Environment {

	var environ *env.Environment

	// A local environment vs a container environment
	if src == "." || src == "." {
		environ = env.New()
	} else {
		container := comp.New(src)
		environ = container.Env()
	}
	return environ

}

// Get a diff from the current environments
func (d *Differ) GetDiff() *Diff {

	diffs := NewDiff()

	// We want change of state from A to B
	for keyB, valB := range d.EnvB.Envars {

		// If we have the value in A it is either unchanged or changed
		if valA, ok := d.EnvA.Envars[keyB]; ok {

			// Unchanged if they are the same
			if valA == valB {
				diffs.Unchanged[keyB] = valB
			} else {
				diffs.Changed[keyB] = Change{Name: keyB, Original: valA, New: valB}
			}

			// We don't have the value in A, so it was added
		} else {
			diffs.Added[keyB] = valB
		}
	}

	// One more loop to find what was removed in B
	for keyA, valA := range d.EnvA.Envars {
		if _, ok := d.EnvB.Envars[keyA]; !ok {
			diffs.Removed[keyA] = valA
		}
	}
	return diffs
}

func (d *Differ) PrintDiff() {

	diffs := d.GetDiff()
	diffs.PrintRemoved()
	diffs.PrintAdded()
	diffs.PrintChanged()

	//	fmt.Println(diffs.Added)
	//	fmt.Println(diffs.Removed)
	//	fmt.Println(diffs.Unchanged)
	//	fmt.Println(diffs.Changed)
}

// Create a new Differ between two sources
func NewDiffer(srcA string, srcB string) *Differ {
	envA := GetEnv(srcA)
	envB := GetEnv(srcB)
	fmt.Println(envA)
	fmt.Println(envB)
	return &Differ{SrcA: srcA, SrcB: srcB, EnvA: envA, EnvB: envB, CreatedAt: time.Now()}
}
