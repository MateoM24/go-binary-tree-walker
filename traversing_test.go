package traversing

import (
	"golang.org/x/tour/tree"
	"testing"
)

func TestWalk(t *testing.T) {

	scenarios := []struct {
		Name           string
		T1             *tree.Tree
		T2             *tree.Tree
		ExpectedResult bool
	}{
		{"Same trees", tree.New(3), tree.New(3), true},
		{"Different tree", tree.New(4), tree.New(3), false},
		{"Empty trees", tree.New(0), tree.New(0), true},
	}

	for _, scenario := range scenarios {
		if scenario.ExpectedResult != Same(scenario.T1, scenario.T2) {
			t.Fatalf("Scenario '%s' has failed", scenario.Name)
		}
	}

}
