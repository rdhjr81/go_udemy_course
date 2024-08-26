package exercises

import "testing"

func TestAdd(t *testing.T) {
	total := add(2, 3)
	if total != 5 {
		t.Errorf("sum incorrect, actual: %d, expected: %d", total, 5)
	}
}

func TestSubtract(t *testing.T) {
	total := subtract(5, 3)
	if total != 2 {
		t.Errorf("sum incorrect, actual: %d, expected: %d", total, 2)
	}
}

func TestDoMath(t *testing.T) {
	initial := 5
	addResult := doMath(initial, 2, add)
	if addResult != 7 {
		t.Errorf("add test result incorrect, actual: %d, expected: %d", 7, addResult)
	}

	subtractResult := doMath(initial, 2, subtract)
	if subtractResult != 3 {
		t.Errorf("add test result incorrect, actual: %d, expected: %d", 3, subtractResult)
	}

}

func TestArea(t *testing.T) {
	sq1 := square{
		length: 5,
		width:  4,
	}
	result := sq1.area()

	if result != 20 {
		t.Errorf("expected area: %d, actual area: %f", 20, result)
	}
}
