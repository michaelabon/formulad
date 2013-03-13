package formulad

import "testing"

func TestCarPositionComparison(t *testing.T) {
	a := CarPosition{1, true, 4, 6}
	b := CarPosition{1, true, 3, 6}
	if a.compare(&b) <= 0 {
		t.Errorf("Car a is in a further corner.")
	}
}
