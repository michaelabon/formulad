package formulad

import "testing"

func TestCarPositionComparison(t *testing.T) {
	a := CarPosition{1, true, 4, 6}
	b := CarPosition{1, true, 3, 6}
	if a.compare(&b) <= 0 {
		t.Errorf("Car a is in a corner ahead of b.")
	}
	if b.compare(&a) >= 0 {
		t.Errorf("Car b is in a corner behind a.")
	}
	a = CarPosition{1, false, 3, 3}
	b = CarPosition{1, true,  3, 3}
	if a.compare(&b) <= 0 {
		t.Errorf("Car a is in the straight ahead of b.")
	}
	if b.compare(&a) >= 0 {
		t.Errorf("Car b in the corner behind a.")
	}
	a = CarPosition{1, false, 3, 5}
	b = CarPosition{1, false, 3, 4}
	if a.compare(&b) <= 0 {
		t.Errorf("Car a is ahead of b in this straight.")
	}
	if b.compare(&a) >= 0 {
		t.Errorf("Car b is behind a in this straight.")
	}
	a = CarPosition{1, true, 3, 5}
	b = CarPosition{1, true, 3, 4}
	if a.compare(&b) <= 0 {
		t.Errorf("Car a is ahead of b in this corner.")
	}
	if b.compare(&a) >= 0 {
		t.Errorf("Car b is behind a in this corner.")
	}
	a = CarPosition{1, false, 3, 4}
	b = CarPosition{1, false, 3, 4}
	if a.compare(&b) != 0 {
		t.Errorf("Car a is beside b in this corner.")
	}
	if b.compare(&a) != 0 {
		t.Errorf("Car b is beside a in this corner.")
	}
	a = CarPosition{1, true, 3, 4}
	b = CarPosition{1, true, 3, 4}
	if a.compare(&b) != 0 {
		t.Errorf("Car a is beside b in this corner.")
	}
	if b.compare(&a) != 0 {
		t.Errorf("Car b is beside a in this corner.")
	}
}

func TestCarPositionSorting(t *testing.T) {
	a = CarPosition{1, true,  3, 6}
	b = CarPosition{1, true,  3, 4}
	c = CarPosition{1, true,  3, 5}

	d = CarPosition{1, false, 3, 5}
	e = CarPosition{1, false, 3, 4}
	f = CarPosition{1, false, 3, 3}

	g = CarPosition{1, true,  2, 8}
	h = CarPosition{1, true,  2, 8}
	i = CarPosition{1, true,  2, 1}

	j = CarPosition{1, false, 2, 2}
	k = CarPosition{1, false, 2, 5}
	m = CarPosition{1, false, 2, 9}

    var positions = []CarPosition{a, b, c, d, e, f, g, h, i, j, k, m}
	sort(&positions)

	for i = 0; i < (len(positions) - 1); i++ {
		if positions[i].compare(&positions[i+1]) > 0 {
			t.Errorf("failed")  // Convert them to text somehow
		}
	}
}
