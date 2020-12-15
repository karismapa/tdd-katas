package calculator

import "testing"

func TestSum(t *testing.T) {
	total := Sum("")
	if total != 0 {
		t.Errorf("Sum was incorrect, got %d, want %d", total, 0)
	}

	total2 := Sum("5")
	if total2 != 5 {
		t.Errorf("Sum was incorrect, got %d, want %d", total, 5)
	}

	total3 := Sum("5,7")
	if total3 != 12 {
		t.Errorf("Sum was incorrect, got %d, want %d", total, 12)
	}

	total4 := Sum("5,7,9,11")
	if total4 != 32 {
		t.Errorf("Sum was incorrect, got %d, want %d", total, 32)
	}

	total5 := Sum("5,7,9\n11")
	if total5 != 32 {
		t.Errorf("Sum was incorrect, got %d, want %d", total, 32)
	}

	total6 := Sum("//;\n5;7;9")
	if total6 != 21 {
		t.Errorf("Sum was incorrect, got %d, want %d", total, 21)
	}

	total7 := Sum("//;\n5;7;9;-5;8")
	if total7 != 29 {
		t.Errorf("Sum was incorrect, got %d, want %d", total, 29)
	}
}
