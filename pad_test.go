package pad

import "testing"

// TestPad will test the Pad function.
func TestPad(t *testing.T) {
	if actual := Pad("foo", LeftDir, 0, '-'); actual != "foo" {
		t.Errorf("Expected 'foo', got '%s'", actual)
	}

	if actual := Pad("foo", LeftDir, 3, '-'); actual != "foo" {
		t.Errorf("Expected 'foo', got '%s'", actual)
	}

	if actual := Pad("foo", LeftDir, 6, '-'); actual != "---foo" {
		t.Errorf("Expected '---foo', got '%s'", actual)
	}

	if actual := Pad("foo", LeftDir, -1, '-'); actual != "foo" {
		t.Errorf("Expected 'foo', got '%s'", actual)
	}

	if actual := Pad("", LeftDir, 5, '-'); actual != "-----" {
		t.Errorf("Expected '-----', got '%s'", actual)
	}

	if actual := Pad("foo", RightDir, 0, '-'); actual != "foo" {
		t.Errorf("Expected 'foo', got '%s'", actual)
	}

	if actual := Pad("foo", RightDir, 3, '-'); actual != "foo" {
		t.Errorf("Expected 'foo', got '%s'", actual)
	}

	if actual := Pad("foo", RightDir, 6, '-'); actual != "foo---" {
		t.Errorf("Expected 'foo---', got '%s'", actual)
	}

	if actual := Pad("foo", RightDir, -1, '-'); actual != "foo" {
		t.Errorf("Expected 'foo', got '%s'", actual)
	}

	if actual := Pad("", RightDir, 5, '-'); actual != "-----" {
		t.Errorf("Expected '-----', got '%s'", actual)
	}
}

// TestLeftChar tests the LeftChar function to ensure it will properly pad the
// left side of strings using a specified character.
func TestLeftChar(t *testing.T) {
	if actual := LeftChar("foo", 0, '-'); actual != "foo" {
		t.Errorf("Expected 'foo', got '%s'", actual)
	}

	if actual := LeftChar("foo", 3, '-'); actual != "foo" {
		t.Errorf("Expected 'foo', got '%s'", actual)
	}

	if actual := LeftChar("foo", 6, '-'); actual != "---foo" {
		t.Errorf("Expected '---foo', got '%s'", actual)
	}

	if actual := LeftChar("foo", -1, '-'); actual != "foo" {
		t.Errorf("Expected 'foo', got '%s'", actual)
	}

	if actual := LeftChar("", 5, '-'); actual != "-----" {
		t.Errorf("Expected '-----', got '%s'", actual)
	}
}

// TestRightChar tests the RightChar function to ensure it will properly pad the
// right side of strings using a specified character.
func TestRightChar(t *testing.T) {
	if actual := RightChar("foo", 0, '-'); actual != "foo" {
		t.Errorf("Expected 'foo', got '%s'", actual)
	}

	if actual := RightChar("foo", 3, '-'); actual != "foo" {
		t.Errorf("Expected 'foo', got '%s'", actual)
	}

	if actual := RightChar("foo", 6, '-'); actual != "foo---" {
		t.Errorf("Expected 'foo---', got '%s'", actual)
	}

	if actual := RightChar("foo", -1, '-'); actual != "foo" {
		t.Errorf("Expected 'foo', got '%s'", actual)
	}

	if actual := RightChar("", 5, '-'); actual != "-----" {
		t.Errorf("Expected '-----', got '%s'", actual)
	}
}

// TestLeft tests the Left function to ensure it will properly pad the left side
// of strings.
func TestLeft(t *testing.T) {
	if actual := Left("foo", 0); actual != "foo" {
		t.Errorf("Expected 'foo', got '%s'", actual)
	}

	if actual := Left("foo", 3); actual != "foo" {
		t.Errorf("Expected 'foo', got '%s'", actual)
	}

	if actual := Left("foo", 6); actual != "   foo" {
		t.Errorf("Expected '   foo', got '%s'", actual)
	}

	if actual := Left("foo", -1); actual != "foo" {
		t.Errorf("Expected 'foo', got '%s'", actual)
	}

	if actual := Left("", 5); actual != "     " {
		t.Errorf("Expected '     ', got '%s'", actual)
	}
}

// TestRight tests the Right function to ensure it will properly pad the right side
// of strings.
func TestRight(t *testing.T) {
	if actual := Right("foo", 0); actual != "foo" {
		t.Errorf("Expected 'foo', got '%s'", actual)
	}

	if actual := Right("foo", 3); actual != "foo" {
		t.Errorf("Expected 'foo', got '%s'", actual)
	}

	if actual := Right("foo", 6); actual != "foo   " {
		t.Errorf("Expected 'foo   ', got '%s'", actual)
	}

	if actual := Right("foo", -1); actual != "foo" {
		t.Errorf("Expected 'foo', got '%s'", actual)
	}

	if actual := Right("", 5); actual != "     " {
		t.Errorf("Expected '     ', got '%s'", actual)
	}
}
