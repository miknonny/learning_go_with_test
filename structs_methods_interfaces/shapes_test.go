package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}

	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// Using table driven test.
// The idea here is not to test similar code we are trying to test similar interfaces.
// Ablity to test if/else blocks of code.
// We name the test so we can tell exactly which values passed into our tests is failing.
func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: &Rectangle{10, 10}, want: 100.0},
		{name: "Circle", shape: &Circle{10}, want: 314.1592653589793},
		{name: "Rectangle", shape: &Triangle{12, 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			CheckShape(t, tt.shape, tt.want)
		})
	}
}

// Helper is used so the error occurs on the line of code being checked and not on the helper function itself.
func CheckShape(t testing.TB, shape Shape, want float64) {
	t.Helper()
	got := shape.Area()

	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
